package config

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Uses Go `embed` package to embed templates in binary
//go:embed templates/*
var templatesFS embed.FS

const (
	CONFIG_DIR       = ".heb-digital"
	CONFIG_FILE_NAME = "gitlab-buddy"
	CONFIG_FILE_EXT  = "yml"
)

type ConfigFile struct {
	Dir      string
	Filename string
	Ext      string
}

func NewConfigFile(path string) (ConfigFile, error) {
	dir, file := filepath.Split(path)
	if file == "" {
		return ConfigFile{}, errors.Unwrap(fmt.Errorf("Invalid file path: %v", path))
	}

	fparts := strings.Split(file, ".")

	// Enforce '.yml' or '.yaml' file extension
	if len(fparts) < 2 || (fparts[1] != "yml" && fparts[1] != "yaml") {
		return ConfigFile{}, errors.Unwrap(fmt.Errorf("Config file must have '.yml' extension: %v", file))
	}

	return ConfigFile{
		Dir:      dir,
		Filename: fparts[0],
		Ext:      fparts[1],
	}, nil
}

func ConfigExists() (bool, error) {
	fileName := strings.Join([]string{CONFIG_FILE_NAME, CONFIG_FILE_EXT}, ".")
	configPath := filepath.Join(ConfigDir, fileName)
	return ExistsOnDisk(configPath)
}

//TODO: Implement config override
func Load() error {

	configFile := ConfigFile{
		Dir:      ConfigDir,
		Filename: CONFIG_FILE_NAME,
		Ext:      CONFIG_FILE_EXT,
	}

	return LoadFile(configFile)
}

func LoadFile(configFile ConfigFile) error {
	// pre-set viper config values
	viper.AddConfigPath(configFile.Dir)
	viper.SetConfigName(configFile.Filename)
	viper.SetConfigType(configFile.Ext)

	dirExists, fileExists := configExists(configFile)

	// Attempt to create a new config file if it doesn't exists
	if !fileExists {
		if !dirExists {
			err := os.Mkdir(configFile.Dir, 0777)
			if err != nil {
				log.Printf("Error creating config directory: %v\n", err)
			}
		}

		err := CreateConfig(configFile)
		if err != nil {
			return err
		}
	} else {
		// If a config file is found, read it in.
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// If only one host is defined, set it as default
	hosts := viper.GetStringMap("hosts")
	if len(hosts) == 1 {
		for key := range hosts {
			viper.Set("hosts.default", key)
		}
	}

	return nil
}

// Create a new config file and read values into viper
func CreateConfig(config ConfigFile) error {

	//TODO: Possibly add support for different config types?
	tpl := template.Must(template.ParseFS(templatesFS, "templates/yaml-config.tmpl"))

	settings := viper.GetViper().AllSettings()

	var tplBuffer bytes.Buffer
	err := tpl.Execute(&tplBuffer, settings)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(config.Dir, config.Filename+"."+config.Ext)
	e := viper.WriteConfigAs(fullPath)
	if e != nil {
		return e
	}

	return nil
}

// TODO: Remove and replace with ExistsOnDisk method in pkg/config/dir.go
func configExists(file ConfigFile) (bool, bool) {
	var dirExists bool
	var fileExists bool

	// Check if config dir exists.
	if _, err := os.Stat(file.Dir); !errors.Is(err, os.ErrNotExist) {
		dirExists = true
	}

	// Check if config file exists
	fullPath := filepath.Join(file.Dir, file.Filename+"."+file.Ext)
	if _, err := os.Stat(fullPath); !errors.Is(err, os.ErrNotExist) {
		fileExists = true
	}

	return dirExists, fileExists
}
