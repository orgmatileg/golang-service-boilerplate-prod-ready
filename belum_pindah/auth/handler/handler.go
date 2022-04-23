package handler

import (
	"database/sql"
	"golang_service/belum_pindah/auth/dto"
	"golang_service/exception"
	"net/http"

	// logicUser "golang_service/module/auth/logic"
	"golang_service/pkg/logger"
	"golang_service/util"

	"github.com/labstack/echo/v4"
)

type handler struct {
	// logic logicUser.IAuthLogic
	db *sql.DB
}

func SetRouter(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	h := &handler{
		// logic: l,
		db: db,
	}

	// Login With PIN
	e.POST("/auth/login/pin", h.AuthLoginWithPIN)

	// Refresh Token
	e.POST("/auth/refresh-token", h.AuthRefreshToken)

}

// AuthRefreshToken ...
func (h *handler) AuthRefreshToken(c echo.Context) error {
	var (
		requestID = c.Request().Header.Get("X-Request-ID")
		// ctx       = c.Request().Context()
		reqData = new(dto.AuthRefreshTokenRequest)
		expt    *exception.Exception
	)

	err := c.Bind(reqData)
	if err != nil {
		logger.Infof("AuthRefreshToken | RequestID=%s | could not c.Bind(reqData): %s", requestID, err.Error())
		return err
	}

	reqData.RequestID = requestID

	// data, expt := h.logic.AuthRefreshToken(ctx, reqData)
	if expt != nil {
		logger.Infof("AuthRefreshToken | RequestID=%s | Exception=%s | could not AuthRefreshToken", requestID, expt.ErrorCode)
		return util.ResponseError(c, expt)
	}

	return util.ResponseSuccess(c, nil, http.StatusOK)
}

// AuthLoginWithPIN ...
func (h *handler) AuthLoginWithPIN(c echo.Context) error {
	var (
		requestID = c.Request().Header.Get("X-Request-ID")
		// ctx       = c.Request().Context()
		reqData = new(dto.AuthLoginWithPINRequest)
		expt    *exception.Exception
	)

	err := c.Bind(reqData)
	if err != nil {
		logger.Infof("AuthLoginWithPIN | RequestID=%s | could not read request data: %s", requestID, err.Error())
		return err
	}

	reqData.RequestID = requestID

	// resp, expt := h.logic.AuthLoginWithPIN(ctx, reqData)
	if expt != nil {
		logger.Infof("AuthLoginWithPIN | RequestID=%s | Exception=%s | could not auth login with phone number", requestID, expt.ErrorCode)
		return util.ResponseError(c, expt)
	}

	return util.ResponseSuccess(c, nil, http.StatusOK)
}
