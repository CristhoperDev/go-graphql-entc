package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const configFile = ".env"

func ReadConfig() error {
	return readConfig()
}

func readConfig() error {
	viper.SetConfigFile(configFile)
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		return err
	}

	return nil
}
