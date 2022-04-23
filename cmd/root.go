package cmd

import (
	"golang_service/config"
	"golang_service/pkg/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang_service",
	Short: "A Backend App",
	Long:  `A Backend App`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	// Init Logger
	err := logger.NewLogger(logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Info,
		ConsoleJSONFormat: true,
		EnableFile:        false,
		FileLevel:         logger.Info,
		FileJSONFormat:    false,
		FileLocation:      "refit-backend.log",
		ServiceName:       "app_backend",
	}, logger.InstanceZapLogger)
	if err != nil {
		logger.Fatalf("Could not instantiate log %s", err.Error())
	}

	config.ParseConfigFromENV()

}
