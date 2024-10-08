package app

import (
	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/cmd/pkg/log"
	"github.com/x-funs/go-fun"
)

// NewApp New app
func NewApp() error {
	// Bootstrap app
	bootstrap()

	// Setup service and set di
	setupDi()

	log.Infof("DateTime : %s", fun.Date())

	return nil
}

// bootstrap Bootstrap app
func bootstrap() {
	// Config map into struct
	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic("Unable to decode config into struct: ")
	}
	log.Info("Config into struct", log.Any("cfg", Cfg))
}

// setupService Setup service and set di
func setupDi() {

}
