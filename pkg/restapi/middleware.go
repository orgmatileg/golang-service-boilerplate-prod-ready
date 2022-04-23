package restapi

import (
	"context"
	"golang_service/config"
	"golang_service/exception"
	"golang_service/pkg/jwtd"
	"golang_service/pkg/logger"
	"golang_service/util"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	datadogEchoTracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func setMiddleware(e *echo.Echo) {
	e.Use(datadogEchoTracer.Middleware())
	e.Use(middleware.BodyLimit("10M"))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Pre(setXRequestIDIfNotPresent)
	e.Pre(setXRequestIDToContext)
	e.Use(middleware.BodyDump(middlewareDumpBodyRequest))
	e.Use(setUserIDToCtx)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"Authorization",
			"X-Request-ID",
			"X-API-KEY",
		},
	}))
}

func setXRequestIDIfNotPresent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-Request-ID") == "" {
			c.Request().Header.Set("X-Request-ID", util.GetRandomString())
		}
		return next(c)
	}
}

func setXRequestIDToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Request().Context()
			key   = "X-Request-ID"
			value = c.Request().Header.Get("X-Request-ID")
		)

		ctx = context.WithValue(ctx, key, value)
		httpReq := c.Request().WithContext(ctx)
		c.SetRequest(httpReq)
		return next(c)
	}
}

// middlewareDumpBodyRequest ...
func middlewareDumpBodyRequest(c echo.Context, reqBody, resBody []byte) {
	requestID := c.Request().Header.Get("X-Request-ID")
	appMobileAppVersionCode := c.Request().Header.Get("CMA-Version-Code")
	appMobileAppVersionName := c.Request().Header.Get("CMA-Version-Name")

	logger.Infof("middlewareDumpBodyRequest | RequestID=%s | app Mobile App Version Code: %s", requestID, appMobileAppVersionCode)
	logger.Infof("middlewareDumpBodyRequest | RequestID=%s | app Mobile App Version Name: %s", requestID, appMobileAppVersionName)
	logger.Infof("middlewareDumpBodyRequest | RequestID=%s | Request URL: %s", requestID, c.Request().RequestURI)
	logger.Infof("middlewareDumpBodyRequest | RequestID=%s | Request Body: %s", requestID, string(reqBody))
	logger.Infof("middlewareDumpBodyRequest | RequestID=%s | Response Body: %s", requestID, string(resBody))

}

// middlewareIsUserAuthorized ...
func middlewareIsUserAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			logger.Infof("middlewareIsUserAuthorized | RequestID=%s | header Authorization is empty", c.Request().Header.Get("X-Request-ID"))
			return util.ResponseError(c, exception.GetException(exception.AUTH_EMPTY_AUTHORIZATION))
		}

		authSplited := strings.Split(auth, " ")

		if len(authSplited) < 2 || len(authSplited) > 2 {
			logger.Infof("middlewareIsUserAuthorized | RequestID=%s | header Authorization is not valid: len not valid", c.Request().Header.Get("X-Request-ID"))
			return util.ResponseError(c, exception.GetException(exception.AUTH_NOT_VALID))
		}

		if authSplited[0] != "Bearer" {
			logger.Infof("middlewareIsUserAuthorized | RequestID=%s | header Authorization is not valid: Bearer not found", c.Request().Header.Get("X-Request-ID"))
			return util.ResponseError(c, exception.GetException(exception.AUTH_NOT_VALID))
		}

		if authSplited[1] == "" {
			logger.Infof("middlewareIsUserAuthorized | RequestID=%s | header Authorization is not valid: token not found", c.Request().Header.Get("X-Request-ID"))
			return util.ResponseError(c, exception.GetException(exception.AUTH_NOT_VALID))
		}

		isValid, err := jwtd.IsValidToken(authSplited[1], config.Get().JWT.SecretAccessToken)
		if err != nil {
			errJwt, _ := err.(*jwt.ValidationError)
			if errJwt.Errors == jwt.ValidationErrorExpired {
				return util.ResponseError(c, exception.GetException(exception.AUTH_EXPIRED))
			}

			if errJwt.Errors != jwt.ValidationErrorSignatureInvalid {
				logger.Infof("middlewareIsUserAuthorized | RequestID=%s | header Authorization is not valid: %s", c.Request().Header.Get("X-Request-ID"), err.Error())
				return util.ResponseError(c, exception.GetException(exception.AUTH_NOT_VALID))
			}

		}

		if isValid {
			return next(c)
		}

		logger.Infof("middlewareIsUserAuthorized | RequestID=%s | somehow token not valid", c.Request().Header.Get("X-Request-ID"), err.Error())
		return util.ResponseError(c, exception.GetException(exception.AUTH_NOT_VALID))
	}
}

func setUserIDToCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		auth := c.Request().Header.Get("Authorization")
		authSplited := strings.Split(auth, " ")

		if len(authSplited) > 1 {
			tc, err := jwtd.ParseClaim(authSplited[1], config.Get().JWT.SecretAccessToken)
			if err != nil {
				logger.Infof("setUserIDToHeader | RequestID=%s | could not set user id to header from token: %s", c.Request().Header.Get("X-Request-ID"), err.Error())
			} else {
				ctx := c.Request().Context()
				ctx = context.WithValue(ctx, "User-ID", strconv.Itoa(tc.Data.UserID))
				ctx = context.WithValue(ctx, "User-ID-Int", tc.Data.UserID)
				httpReq := c.Request().WithContext(ctx)
				c.SetRequest(httpReq)
			}
		}

		return next(c)
	}
}
