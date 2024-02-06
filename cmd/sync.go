package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/urfave/cli/v2"
)

func sync(cCtx *cli.Context) error {
	fmt.Println("sync")

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}
	config, err := loadConfig(configPath)
	if err != nil {
		return err
	}
	fmt.Println(config)

	for _, extension := range config.Extensions {
		normalizedRepoName := normalizeRepoName(&extension)
		fmt.Println(normalizedRepoName)
		extensionPath, err := getExtensionPath()
		if err != nil {
			return err
		}

		cloneTo := filepath.Join(extensionPath, getCloneTo(&extension))
		fmt.Println(cloneTo)
		stat, err := os.Stat(cloneTo)
		if err != nil {
			if err = clone(normalizedRepoName, cloneTo); err != nil {
				return err
			}
		} else if stat.IsDir() {
			// NOTE: exist already
			fmt.Println("exist already")
		} else {
			return fmt.Errorf("same file exists?")
		}

		repo, err := git.PlainOpen(cloneTo)
		if err != nil {
			return err
		}
		worktree, err := repo.Worktree()
		if err != nil {
			return err
		}
		var hash plumbing.Hash
		tagRef, err := repo.Tag(extension.Version)
		if err != nil {
			hash = plumbing.NewHash(extension.Version)
		} else {
			hash = tagRef.Hash()
		}
		err = worktree.Checkout(&git.CheckoutOptions{
			Hash: hash,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
