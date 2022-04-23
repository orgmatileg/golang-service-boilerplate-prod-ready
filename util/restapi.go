package util

import (
	"golang_service/exception"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseError ...
func ResponseError(c echo.Context, e *exception.Exception) error {

	res := make(map[string]string)
	res["ErrorCode"] = e.ErrorCode
	res["Message"] = e.Message

	return c.JSON(e.StatusCode, res)
}

// ResponseSuccess ...
func ResponseSuccess(c echo.Context, data interface{}, statusCode int) error {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	if data == nil {
		res := make(map[string]string)
		res["Message"] = "success"
		return c.JSON(statusCode, res)
	}

	return c.JSON(statusCode, data)
}
