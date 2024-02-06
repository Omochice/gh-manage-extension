package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pelletier/go-toml/v2"
)

type Extension struct {
	Repo    string
	Version string
	As      string
}

type Configuration struct {
	Extensions []Extension
}

func loadConfig(path string) (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(path)
	if err != nil {
		return configuration, err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return configuration, err
	}
	err = toml.Unmarshal(content, &configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}

func getConfigPath() (string, error) {
	// NOTE: support xdg
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// NOTE: support DI
	return filepath.Join(home, ".config", "gh-manage-extension", "config.toml"), nil
}

func normalizeRepoName(extension *Extension) string {
	allowdRegex := "[a-zA-Z0-9]([a-zA-Z0-9]?|[\\-]?([a-zA-Z0-9])){0,38}"
	re := regexp.MustCompile(fmt.Sprintf("^%s/%s$", allowdRegex, allowdRegex))
	if re.Match([]byte(extension.Repo)) {
		return fmt.Sprintf("https://github.com/%s", extension.Repo)
	}
	return extension.Repo
}

func getCloneTo(extension *Extension) string {
	if extension.As != "" {
		return extension.As
	}
	re := regexp.MustCompile("/(.+)$")
	b := re.FindSubmatch([]byte(extension.Repo))
	return string(b[1]) // first marched one
}
