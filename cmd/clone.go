package cmd

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func clone(url string, cloneTo string) error {
	_, err := git.PlainClone(cloneTo, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}
