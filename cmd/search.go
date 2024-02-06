package cmd

import (
	"encoding/json"

	gh "github.com/cli/go-gh"
	"github.com/koki-develop/go-fzf"
)

type GHExtension struct {
	Description string `json:"description"`
	FullName    string `json:"fullName"`
}

func search() (GHExtension, error) {
	selected := GHExtension{}
	args := []string{"extension", "search", "--json", "description,fullName"}
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return selected, err
	}
	extensions := []GHExtension{}
	json.Unmarshal(stdOut.Bytes(), &extensions)
	f, err := fzf.New()
	if err != nil {
		return selected, err
	}
	idx, err := f.Find(extensions,
		func(i int) string { return extensions[i].FullName },
		fzf.WithPreviewWindow(func(i, _, __ int) string {
			return extensions[i].Description
		}))
	if err != nil {
		return selected, err
	}
	return extensions[idx[0]], nil
}
