package app

import (
	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/cmd/config"
	"github.com/suosi-inc/go-demo/cmd/pkg/log"
)

var (
	cfg = config.Cfg
)

// NewApp New app
func NewApp() error {
	// Bootstrap app
	bootstrap()

	// Setup service and set di
	setupDi()

	return nil
}

// bootstrap Bootstrap app
func bootstrap() {
	// Config map into struct
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic("Unable to decode config into struct: ")
	}
	log.Info("Config into struct", log.Any("cfg", cfg))
}

// setupService Setup service and set di
func setupDi() {

}
