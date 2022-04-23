package restapi

import (
	"github.com/labstack/echo/v4"
)

func setRouter(e *echo.Echo) {
	setRouterForUser(e)

}
