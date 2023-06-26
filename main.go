package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/cmd/internal"
	"github.com/suosi-inc/go-demo/cmd/pkg"
	"log"
	"os"
)

const (
	AppName      = "cmd"
	AppShortDesc = "cmd entry"
	AppLongDesc  = "cmd entry is cmd app"
)

var (
	cfgFile string

	AppCmd = &cobra.Command{
		Use:          AppName,
		Short:        AppShortDesc,
		Long:         AppLongDesc,
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Println("App preRun")

			// Init viper config
			initConfig()

			// Init zap logger and set to global log
			pkg.InitZapLogger()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// New app and run
			return internal.NewApp()
		},
	}
)

func main() {
	Execute()
}

// init
func init() {
	// Bind command flags
	AppCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

// initConfig Init config
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Load default config file
		viper.SetConfigName(AppName)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")
	}

	// Load environment variables
	viper.AutomaticEnv()

	// Default config
	defaultConfig()

	// Read config file into viper config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic("Config file not found")
		} else {
			log.Panic(err)
		}
	}

	log.Println("App using config file:", viper.ConfigFileUsed())

	mergeEnvConfig()
}

// defaultConfig Default config
func defaultConfig() {
	viper.SetDefault("app.name", AppName)

	viper.SetDefault("logger.file", AppName+".log")
	viper.SetDefault("logger.format", "text")
	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.maxsize", 128)
	viper.SetDefault("logger.maxage", 3)
	viper.SetDefault("logger.maxbackups", 7)
}

// Merge in environment specific config
func mergeEnvConfig() {
	configFilePath := ""
	env := []string{"product", "develop"}

out:
	for _, e := range env {
		configName := AppName + "-" + e + ".yaml"
		configPaths := []string{configName, "./config/" + configName}

		for _, path := range configPaths {
			if _, err := os.Stat(path); err == nil {
				configFilePath = path
				break out
			}
		}
	}

	if configFilePath == "" {
		return
	}
	configBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("Could not read config file: %s \n", err))
	}
	err = viper.MergeConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		panic(fmt.Errorf("Merge config file error: %s \n", err))
	}

	log.Println("App using merge config file:", configFilePath)
}

func Execute() {
	if err := AppCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
