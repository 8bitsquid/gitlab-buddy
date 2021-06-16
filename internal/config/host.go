package config

import (
	"errors"
	"net/url"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	GITLAB_API_PATH   = "/api/v4/"
	GITLAB_URL        = "https://gitlab.com"
	DEFAULT_HOST_TYPE = "gitlab"
)

type Host struct {
	APIPath string `yaml:"apipath" mapstructure:"apipath"`
	URL     string `yaml:"url" mapstructure:"url"`
	Token   string `yaml:"token" mapstructure:"token"`
	Type    string `yaml:"type" mapstructure:"type"`
}

func GetHost(name string) (*Host, error) {

	if name == "" {
		name = GetDefaultHost()
		if name == "" {
			return nil, errors.New("at least one host must be defined in config file")
		}
	}
	
	hostKey := "hosts." + name

	hostConfig := viper.Sub(hostKey)
	if !hostConfig.IsSet("token") {
		zap.S().Panicw("Token not provided for host", "host", hostConfig.AllSettings())
	}
	if !hostConfig.IsSet("url") {
		hostConfig.Set("url", GITLAB_URL)
	}
	if !hostConfig.IsSet("apipath") {
		hostConfig.Set("apipath", GITLAB_API_PATH)
	}
	if !hostConfig.IsSet("type") {
		hostConfig.Set("type", DEFAULT_HOST_TYPE)
	}

	viper.Set(hostKey, hostConfig.AllSettings())

	var host Host
	viper.UnmarshalKey(hostKey, &host)

	return &host, nil
}

func GetDefaultHost() string {
	return viper.GetString("hosts.default")
}

func (h *Host) GetBaseURL() *url.URL {

	hostURL, err := parseHost(h.URL)
	if err != nil {
		zap.S().Panicw("Invalid host url", "host", h.URL)
	}
	apiPath, err := url.Parse(h.APIPath)
	if err != nil {
		zap.S().Panicw("Invalid API Path", "api_path", h.APIPath)
	}

	return hostURL.ResolveReference(apiPath)

}

//helper functions
func parseHost(hostURL string) (*url.URL, error) {
	hostURL = strings.Replace(hostURL, "www.gitlab", "gitlab", 1)
	hostURL = strings.Replace(hostURL, "http:", "https:", 1)

	return url.Parse(hostURL)
}
