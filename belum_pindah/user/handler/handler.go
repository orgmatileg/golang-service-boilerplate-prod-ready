package handler

// import (
// 	"database/sql"
// 	"net/http"
// 	"golang_service/exception"
// 	"golang_service/module/user/dto"
// 	"golang_service/module/user/logic"
// 	"golang_service/pkg/logger"
// 	"golang_service/util"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type handler struct {
// 	logic logic.IUserLogic
// 	db    *sql.DB
// }

// func SetRouter(e *echo.Echo, l logic.IUserLogic, db *sql.DB, m ...echo.MiddlewareFunc) {
// 	h := &handler{
// 		logic: l,
// 		db:    db,
// 	}

// 	g := e.Group("/users", m...)

// 	// g.POST("", h.Create, m...)	// Cek apakah dipake? kalau engga hapus aja
// 	g.GET("", h.FindAll, m...)        // Cek apakah dipake? kalau engga hapus aja
// 	g.GET("/:id", h.FindByID, m...)   // Cek apakah dipake? kalau engga hapus aja
// 	g.PUT("/:id", h.UpdateByID, m...) // Cek apakah dipake? kalau engga hapus aja
// 	g.POST("/phone", h.RequestUpdatePhoneNumber, m...)
// 	g.POST("/phone/otp", h.UpdatePhoneNumberOTP, m...)
// 	g.POST("/pin/check", h.PINCheck, m...)
// 	g.DELETE("/delete", h.DeleteAccount, m...)
// 	g.POST("/delete/otp", h.DeleteAccountOTP, m...)

// }

// // FindAll ...
// func (h *handler) FindAll(c echo.Context) error {

// 	var (
// 		lastID    = c.QueryParam("last_id")
// 		limit     = c.QueryParam("limit")
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 	)

// 	lastIDInt, err := strconv.Atoi(lastID)
// 	if err != nil {
// 		logger.Infof("FindByID | RequestID=%s | could not parse last_id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	limitInt, err := strconv.Atoi(limit)
// 	if err != nil {
// 		logger.Infof("FindByID | RequestID=%s | could not parse limit to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	resp, expt := h.logic.FindAll(ctx, requestID, int64(lastIDInt), int64(limitInt))
// 	if expt != nil {
// 		logger.Infof("FindByID | RequestID=%s | Exception=%s | could not find user by id logic", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, resp, http.StatusOK)
// }

// // UpdateByID ...
// func (h *handler) UpdateByID(c echo.Context) error {
// 	var (
// 		requestID                      = c.Request().Header.Get("X-Request-ID")
// 		ctx                            = c.Request().Context()
// 		reqData                        = new(dto.UpdateByIDRequest)
// 		expt      *exception.Exception = nil
// 	)

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("UpdateByID | RequestID=%s | cannot get UserID from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	tx, errTx := h.db.Begin()
// 	if errTx != nil {
// 		logger.Errorf("UpdateByID | RequestID=%s | could not begin transaction to db: %s", requestID, errTx.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR))
// 	}

// 	defer func() {
// 		var errTx error
// 		if expt != nil {
// 			errTx = tx.Rollback()
// 		} else {
// 			errTx = tx.Commit()
// 		}

// 		if errTx != nil {
// 			logger.Infof("UpdateByID | RequestID=%s | could not commit/rollback to db: %s", requestID, errTx.Error())
// 		}
// 	}()

// 	_, expt = h.logic.UpdateByID(ctx, requestID, int64(userIDInt), reqData, tx)
// 	if expt != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | Exception=%s | could not update user by id", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusCreated)
// }

// // FindByID ...
// func (h *handler) FindByID(c echo.Context) error {

// 	var (
// 		id        = c.Param("id")
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 	)

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		logger.Infof("FindByID | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	resp, expt := h.logic.FindByID(ctx, requestID, int64(idInt))
// 	if expt != nil {
// 		logger.Infof("FindByID | RequestID=%s | Exception=%s | could not find user by id logic", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, resp, http.StatusOK)
// }

// // Create ...
// func (h *handler) Create(c echo.Context) error {
// 	var (
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 		reqData   = new(dto.CreateRequest)
// 	)

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("Create | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	_, expt := h.logic.Create(ctx, requestID, reqData)
// 	if expt != nil {
// 		logger.Infof("Create | RequestID=%s | Exception=%s | could not create user", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusCreated)
// }

// // DeleteAccount ...
// func (h *handler) DeleteAccount(c echo.Context) error {

// 	var (
// 		requestID                      = c.Request().Header.Get("X-Request-ID")
// 		ctx                            = c.Request().Context()
// 		reqData                        = new(dto.DeleteAccountRequest)
// 		expt      *exception.Exception = nil
// 	)

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("DeleteAccount | RequestID=%s | cannot get UserID from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	idInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}
// 	reqData.UserID = int64(idInt)
// 	reqData.RequestID = requestID

// 	tx, errTx := h.db.Begin()
// 	if errTx != nil {
// 		logger.Errorf("Create | RequestID=%s | could not begin transaction to db: %s", requestID, errTx.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR))
// 	}

// 	defer func() {
// 		var errTx error
// 		if expt != nil {
// 			errTx = tx.Rollback()
// 		} else {
// 			errTx = tx.Commit()
// 		}

// 		if errTx != nil {
// 			logger.Infof("Create | RequestID=%s | could not commit/rollback to db: %s", requestID, errTx.Error())
// 		}
// 	}()

// 	if expt := h.logic.DeleteAccount(ctx, reqData, tx); expt != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | Exception=%s | could not delete user by id logic", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusOK)
// }

// // RequestUpdatePhoneNumber ...
// func (h *handler) RequestUpdatePhoneNumber(c echo.Context) error {
// 	var (
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 		reqData   = new(dto.RequestUpdatePhoneNumber)
// 	)

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | cannot get Role from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}
// 	reqData.UserID = int64(userIDInt)
// 	reqData.RequestID = requestID

// 	if expt := h.logic.RequestUpdatePhoneNumber(ctx, reqData); expt != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | Exception=%s | could not update user by id", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusOK)
// }

// // UpdatePhoneNumberOTP ...
// func (h *handler) UpdatePhoneNumberOTP(c echo.Context) error {
// 	var (
// 		requestID                      = c.Request().Header.Get("X-Request-ID")
// 		ctx                            = c.Request().Context()
// 		reqData                        = new(dto.UpdatePhoneNumberOTP)
// 		expt      *exception.Exception = nil
// 	)

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | cannot get Role from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}
// 	reqData.UserID = int64(userIDInt)
// 	reqData.RequestID = requestID

// 	tx, errTx := h.db.Begin()
// 	if errTx != nil {
// 		logger.Errorf("Create | RequestID=%s | could not begin transaction to db: %s", requestID, errTx.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR))
// 	}

// 	defer func() {
// 		var errTx error
// 		if expt != nil {
// 			errTx = tx.Rollback()
// 		} else {
// 			errTx = tx.Commit()
// 		}

// 		if errTx != nil {
// 			logger.Infof("Create | RequestID=%s | could not commit/rollback to db: %s", requestID, errTx.Error())
// 		}
// 	}()

// 	if _, expt := h.logic.UpdatePhoneNumberOTP(ctx, reqData, tx); expt != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | Exception=%s | could not update user by id", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusOK)
// }

// // PINCheck ...
// func (h *handler) PINCheck(c echo.Context) error {
// 	var (
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 		reqData   = new(dto.PINCheckRequest)
// 	)

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("PINCheck | RequestID=%s | cannot get Role from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("PINCheck | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("PINCheck | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}
// 	reqData.UserID = int64(userIDInt)
// 	reqData.RequestID = requestID

// 	resp, expt := h.logic.PINCheck(ctx, reqData)
// 	if expt != nil {
// 		logger.Infof("PINCheck | RequestID=%s | Exception=%s | could not update user by id", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, resp, http.StatusOK)
// }

// // DeleteAccountOTP ...
// func (h *handler) DeleteAccountOTP(c echo.Context) error {
// 	var (
// 		requestID = c.Request().Header.Get("X-Request-ID")
// 		ctx       = c.Request().Context()
// 		reqData   = new(dto.DeleteAccountOTPRequest)
// 	)

// 	userID, ok := ctx.Value("User-ID").(string)
// 	if !ok {
// 		logger.Infof("DeleteAccountOTP | RequestID=%s | cannot get Role from context: %s", requestID)
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}

// 	err := c.Bind(reqData)
// 	if err != nil {
// 		logger.Infof("DeleteAccountOTP | RequestID=%s | could not read request data: %s", requestID, err.Error())
// 		return err
// 	}

// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		logger.Infof("DeleteAccountOTP | RequestID=%s | could not parse id to int: %s", requestID, err.Error())
// 		return util.ResponseError(c, exception.GetException(exception.GENERAL_BAD_REQUEST))
// 	}
// 	reqData.UserID = int64(userIDInt)
// 	reqData.RequestID = requestID

// 	if expt := h.logic.DeleteAccountOTP(ctx, reqData); expt != nil {
// 		logger.Infof("DeleteAccountOTP | RequestID=%s | Exception=%s | could not update user by id", requestID, expt.ErrorCode)
// 		return util.ResponseError(c, expt)
// 	}

// 	return util.ResponseSuccess(c, nil, http.StatusOK)
// }
