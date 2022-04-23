package echo

import (
	"context"
	"golang_service/exception"
	"golang_service/util"
	"log"
	"net/http"

	labstackEcho "github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEcho(lc fx.Lifecycle) *labstackEcho.Echo {

	e := labstackEcho.New()

	e.IPExtractor = labstackEcho.ExtractIPDirect()
	e.HTTPErrorHandler = customHTTPErrorHandler

	// setMiddleware(e)

	e.GET("/health", func(c labstackEcho.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting HTTP server.")
			return e.Start("0.0.0.0:8080")
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server.")
			return e.Shutdown(ctx)
		},
	})

	return e
}

func customHTTPErrorHandler(err error, c labstackEcho.Context) {
	if err != nil {
		// logger.Errorf("there is an error unhandled: %s", err.Error())
	}
	util.ResponseError(c, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR))
}
