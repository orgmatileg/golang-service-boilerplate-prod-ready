package cmd

import (
	"context"
	"fmt"
	"golang_service/config"
	"golang_service/pkg/echo"
	"golang_service/pkg/logger"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func init() {
	rootCmd.AddCommand(serveApiCMD)
}

var serveApiCMD = &cobra.Command{
	Use:   "serve-api",
	Short: "Start Serving REST API app",
	Long:  `Start Serving REST API app`,
	Run: func(cmd *cobra.Command, args []string) {

		err := logger.NewLogger(logger.Configuration{
			EnableConsole: true,
			ConsoleLevel:  logger.Info,
			ServiceName:   "app_backend",
		}, logger.InstanceZapLogger)
		if err != nil {
			logger.Fatalf("Could not instantiate log %s", err.Error())
		}

		logger.Infof("Starting server...")

		config.ParseConfigFromENV()

		ctx, cancel := context.WithCancel(context.Background())
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		go func() {
			<-quit
			cancel()
		}()

		app := fx.New(
			// fx.Provide(labstackEcho.Echo),
			fx.Invoke(echo.NewEcho),
		)
		if err := app.Start(ctx); err != nil {
			fmt.Println(err)
		}

	},
}
