package cmd

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func clone(url string) error {
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}
