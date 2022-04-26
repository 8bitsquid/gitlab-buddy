package pkg

import (
	"github.com/spf13/viper"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {

	// Get logging config values
	var zapConfig zap.Config
	if viper.GetString(config.CONFIG_KEY_LOG+"."+config.LOG_KEY_LEVEL) == "debug" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.OutputPaths = viper.GetStringSlice(config.CONFIG_KEY_LOG + "." + config.LOG_KEY_OUTPUT)

	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: config.LOG_KEY_MESSAGE,

		LevelKey:    config.LOG_KEY_LEVEL,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,

		TimeKey:    config.LOG_KEY_TIMESTAMP,
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    config.LOG_KEY_CALLER,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}
