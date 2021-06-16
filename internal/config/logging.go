package config

import (
	"github.com/spf13/viper"
)

// Ok, ok this is overkill, but it allows sharing config value keys to configure loggers in differe packages
const (
	// Key used in the config got logging options
	CONFIG_KEY_LOG = "logging"

	// Config default values
	LOG_LEVEL = "debug"
	LOG_OUPUT = "stderr"

	// Config log keys used to define options
	LOG_KEY_OUTPUT    = "output"
	LOG_KEY_MESSAGE   = "message"
	LOG_KEY_LEVEL     = "level"
	LOG_KEY_CALLER    = "caller"
	LOG_KEY_TIMESTAMP = "timestamp"
)

func init() {
	viper.SetDefault(CONFIG_KEY_LOG+"."+LOG_KEY_OUTPUT, []string{LOG_OUPUT})
	viper.SetDefault(CONFIG_KEY_LOG+"."+LOG_KEY_LEVEL, LOG_LEVEL)
}
