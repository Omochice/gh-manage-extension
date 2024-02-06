package cmd

import (
	"os"
	"path/filepath"
)

func getExtensionPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// TODO: support DI
	// TODO: use xdg
	// TODO: support windows
	// extensionPath := filepath.Join(home, ".local", "share", "gh", "extenions")
	extensionPath := filepath.Join(home, ".local", "share", "gh", "extensions")
	if err = os.MkdirAll(extensionPath, 0777); err != nil {
		return "", err
	}
	return extensionPath, nil
}
