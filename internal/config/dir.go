package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

const (
	CONFIG_DIR_NAME = ".heb-digital"
)

var HomeDir string
var ConfigDir string

func init() {
	HomeDir = GetHomeDir()
	ConfigDir = filepath.Join(HomeDir, CONFIG_DIR_NAME)
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to find user's home directory: %+v", err)
	}
	return homeDir
}

// TODO: Move to tools package
// Checks if the given filepath exists on local disk
func ExistsOnDisk(filePath string) (bool, error) {
	zap.S().Infof("Checking if path exists: %v", filePath)
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	return true, nil
}
