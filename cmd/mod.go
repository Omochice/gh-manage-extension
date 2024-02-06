package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install some extension",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "uninstall",
				Usage: "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:   "sync",
				Usage:  "Sync extensions",
				Action: sync,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
	// log.Fatalf("end")
	// configPath, err := getConfigPath()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(configPath)
	// config, err := loadConfig(configPath)
	// if err != nil {
	// 	return err
	// }

	// re := regexp.MustCompile(`/(.+)$`)

	// extensionPath, err := getExtensionPath()
	// if err != nil {
	// 	return err
	// }
	// for _, extension := range config.Extensions {
	// 	matched := re.FindSubmatch([]byte(extension.Repo))
	// 	if len(matched) < 2 {
	// 		return fmt.Errorf("repo is invalid")
	// 	}
	// 	repoName := string(matched[1])
	// 	// TODO: use goroutine
	// 	url := fmt.Sprintf("https://github.com/%s", extension.Repo)
	// 	cloneTo := filepath.Join(extensionPath, repoName)
	// 	err = clone(url, cloneTo)
	// 	fmt.Println(extension.As == "", url, extensionPath, extension.Repo, repoName)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// return nil
}
