package logic

// import (
// 	"fmt"
// 	"golang_service/config"
// 	"golang_service/database/redisd"
// 	"golang_service/enum"
// 	"golang_service/exception"
// 	"golang_service/external/qiscus"
// 	"golang_service/external/xendit"
// 	"golang_service/model"
// 	logicAnalytic "golang_service/module/analytic/logic"
// 	repositoryCoaBalance "golang_service/module/coaBalance/repository"
// 	repositoryStore "golang_service/module/store/repository"
// 	"golang_service/module/user/dto"
// 	repositoryUser "golang_service/module/user/repository"
// 	repositoryUserDelete "golang_service/module/userDelete/repository"

// 	"context"
// 	"database/sql"
// 	"golang_service/pkg/logger"
// 	"golang_service/util"
// 	"strconv"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"golang.org/x/crypto/bcrypt"
// )

// type IUserLogic interface {
// 	Create(ctx context.Context, requestID string, reqData *dto.CreateRequest) (int64, *exception.Exception)
// 	FindByID(ctx context.Context, requestID string, id int64) (*dto.FindByIDResponse, *exception.Exception)
// 	UpdateByID(ctx context.Context, requestID string, id int64, reqData *dto.UpdateByIDRequest, tx *sql.Tx) (int64, *exception.Exception)
// 	DeleteByID(ctx context.Context, requestID string, id int64) (int64, *exception.Exception)
// 	FindAll(ctx context.Context, requestID string, lastID, limit int64) (dto.FindAllResponse, *exception.Exception)
// 	RequestUpdatePhoneNumber(ctx context.Context, reqData *dto.RequestUpdatePhoneNumber) *exception.Exception
// 	UpdatePhoneNumberOTP(ctx context.Context, reqData *dto.UpdatePhoneNumberOTP, tx *sql.Tx) (*int64, *exception.Exception)
// 	PINCheck(ctx context.Context, reqData *dto.PINCheckRequest) (*bool, *exception.Exception)
// 	DeleteAccountOTP(ctx context.Context, reqData *dto.DeleteAccountOTPRequest) *exception.Exception
// 	DeleteAccount(ctx context.Context, reqData *dto.DeleteAccountRequest, tx *sql.Tx) *exception.Exception
// }

// type userLogic struct {
// 	db                   *sql.DB
// 	repositoryUser       repositoryUser.IUserRepository
// 	repositoryCoaBalance repositoryCoaBalance.ICoaBalanceRepository
// 	repositoryUserDelete repositoryUserDelete.IUserDeleteRepository
// 	repositoryStore      repositoryStore.IStoreRepository
// 	logicAnalytic        logicAnalytic.IAnalyticLogic
// 	contextTimeout       time.Duration
// }

// // New ...
// func New(ru repositoryUser.IUserRepository,
// 	la logicAnalytic.IAnalyticLogic,
// 	repoCoaBalance repositoryCoaBalance.ICoaBalanceRepository,
// 	repoUserDelete repositoryUserDelete.IUserDeleteRepository,
// 	repoStore repositoryStore.IStoreRepository,
// 	timeout time.Duration, db *sql.DB) IUserLogic {
// 	return &userLogic{
// 		db:                   db,
// 		repositoryUser:       ru,
// 		repositoryCoaBalance: repoCoaBalance,
// 		repositoryUserDelete: repoUserDelete,
// 		repositoryStore:      repoStore,
// 		logicAnalytic:        la,
// 		contextTimeout:       timeout,
// 	}
// }

// // FindByID ...
// func (u *userLogic) FindAll(ctx context.Context, requestID string, lastID, limit int64) (dto.FindAllResponse, *exception.Exception) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	// make dto response
// 	listUserModel, err := u.repositoryUser.FindAll(ctx, requestID, lastID, limit)
// 	if err != nil {
// 		logger.Infof("FindAll | RequestID=%s | could not user find all: %s", requestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return nil, exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	return listUserModel, nil
// }

// // UpdateByID ...
// func (u *userLogic) UpdateByID(ctx context.Context, requestID string, id int64, reqData *dto.UpdateByIDRequest, tx *sql.Tx) (int64, *exception.Exception) {

// 	reqData.ID = id

// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	var (
// 		err error
// 		m   = new(model.User)
// 	)

// 	err = reqData.Validate()
// 	if err != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | %#v | error when validate request data: %s", requestID, reqData, err.Error())
// 		return -1, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	usr, err := u.repositoryUser.FindByID(ctx, requestID, id)
// 	if err != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | could not user find by id: %s", requestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return -1, exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	m.ID = int(reqData.ID)
// 	m.FullName = reqData.FullName
// 	m.Status = usr.Status
// 	m.Email = usr.Email
// 	m.CreatedAt = usr.CreatedAt
// 	m.UpdatedAt = usr.UpdatedAt

// 	if usr.Email == nil {
// 		_, err = u.repositoryUser.FindByEmail(ctx, requestID, reqData.Email)
// 		if err == nil {
// 			return -1, exception.GetException(exception.USER_EMAIL_ALREADY_EXIST)
// 		} else if err != nil && err != sql.ErrNoRows {
// 			logger.Infof("UpdateByID | RequestID=%s | could not find user by email: %s", requestID, err.Error())
// 		}
// 		m.Email = &reqData.Email
// 	} else if reqData.Email != "" && usr.Email != nil && reqData.Email != *usr.Email {
// 		_, err = u.repositoryUser.FindByEmail(ctx, requestID, reqData.Email)
// 		if err == nil {
// 			return -1, exception.GetException(exception.USER_EMAIL_ALREADY_EXIST)
// 		} else if err != nil && err != sql.ErrNoRows {
// 			logger.Infof("UpdateByID | RequestID=%s | could not find user by email: %s", requestID, err.Error())
// 		}
// 		m.Email = &reqData.Email
// 	}

// 	insertedID, err := u.repositoryUser.UpdateByID(ctx, requestID, m, tx)
// 	if err != nil {
// 		logger.Infof("UpdateByID | RequestID=%s | could not user delete by id: %s", requestID, err.Error())
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	// Call Logic Analytics for Update Store to Clever Tap
// 	ra := u.logicAnalytic.CleverTapUpdateUser(ctx, requestID, *m)
// 	if ra {
// 		logger.Infof("UpdateByID | RequestID=%s | could not cleverTapUpdateStore: %s", requestID, ra)
// 	}

// 	return insertedID, nil
// }

// // FindByID ...
// func (u *userLogic) FindByID(ctx context.Context, requestID string, id int64) (*dto.FindByIDResponse, *exception.Exception) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	userModel, err := u.repositoryUser.FindByID(ctx, requestID, id)
// 	if err != nil {
// 		logger.Infof("FindByID | RequestID=%s | could not user find by id: %s", requestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return nil, exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	resp := dto.FindByIDResponse{
// 		User: *userModel,
// 	}

// 	return &resp, nil
// }

// // Create ...
// func (u *userLogic) Create(ctx context.Context, requestID string, reqData *dto.CreateRequest) (int64, *exception.Exception) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	var err error

// 	err = reqData.Validate()
// 	if err != nil {
// 		logger.Infof("Create | RequestID=%s | %#v | error when validate request data: %s", requestID, reqData, err.Error())
// 		return -1, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	_, err = u.repositoryUser.FindByEmail(ctx, requestID, reqData.Email)
// 	if err == nil {
// 		return -1, exception.GetException(exception.USER_EMAIL_ALREADY_EXIST)
// 	} else if err != nil && err != sql.ErrNoRows {
// 		logger.Infof("Create | RequestID=%s | could not find user by email: %s", requestID, err.Error())
// 	}

// 	_, err = u.repositoryUser.FindByPhoneNumber(ctx, requestID, reqData.PhoneNumber)
// 	if err == nil {
// 		return -1, exception.GetException(exception.USER_PHONE_NUMBER_ALREADY_EXIST)
// 	} else if err != nil && err != sql.ErrNoRows {
// 		logger.Infof("Create | RequestID=%s | could not find user by phone number: %s", requestID, err.Error())
// 	}

// 	tx, errTx := u.db.Begin()
// 	if errTx != nil {
// 		logger.Infof("Create | RequestID=%s | could not begin transaction to db: %s", requestID, errTx.Error())
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	defer func() {
// 		var errTx error
// 		if err != nil {
// 			errTx = tx.Rollback()
// 		} else {
// 			errTx = tx.Commit()
// 		}

// 		if errTx != nil {
// 			logger.Infof("Create | RequestID=%s | could not commit/rollback to db: %s", requestID, errTx.Error())
// 		}
// 	}()

// 	model := model.User{
// 		FullName:    reqData.FullName,
// 		PhoneNumber: reqData.PhoneNumber,
// 		Email:       &reqData.Email,
// 		Status:      &reqData.Status,
// 	}

// 	insertedID, err := u.repositoryUser.Create(ctx, requestID, &model, tx)
// 	if err != nil {
// 		logger.Infof("Create | RequestID=%s | could not user delete by id: %s", requestID, err.Error())
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	return insertedID, nil
// }

// // DeleteByID ...
// func (u *userLogic) DeleteByID(ctx context.Context, requestID string, id int64) (int64, *exception.Exception) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	var err error

// 	_, err = u.repositoryUser.FindByID(ctx, requestID, id)
// 	if err != nil {
// 		logger.Infof("DeleteByID | RequestID=%s | could not user find by id: %s", requestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return -1, exception.GetException(exception.GENERAL_NOT_FOUND)
// 		}
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	tx, errTx := u.db.Begin()
// 	if errTx != nil {
// 		logger.Infof("DeleteByID | RequestID=%s | could not begin transaction to db: %s", requestID, errTx.Error())
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	defer func() {
// 		var errTx error
// 		if err != nil {
// 			errTx = tx.Rollback()
// 		} else {
// 			errTx = tx.Commit()
// 		}

// 		if errTx != nil {
// 			logger.Infof("DeleteByID | RequestID=%s | could not commit/rollback to db: %s", requestID, errTx.Error())
// 		}
// 	}()

// 	totalDeleted, err := u.repositoryUser.DeleteByID(ctx, requestID, id, tx)
// 	if err != nil {
// 		logger.Infof("DeleteByID | RequestID=%s | could not user delete by id: %s", requestID, err.Error())
// 		return -1, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	return totalDeleted, nil
// }

// // RequestUpdatePhoneNumber ...
// func (u *userLogic) RequestUpdatePhoneNumber(ctx context.Context, reqData *dto.RequestUpdatePhoneNumber) *exception.Exception {

// 	var err error

// 	err = reqData.Validate()
// 	if err != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | %#v | error when validate request data: %s", reqData.RequestID, reqData, err.Error())
// 		return exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	_, err = u.repositoryUser.FindByID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	existPhoneNumber, err := u.repositoryUser.FindByPhoneNumber(ctx, reqData.RequestID, reqData.PhoneNumber)
// 	if err != nil && err != sql.ErrNoRows {
// 		logger.Infof("RequestUpdatePhoneNumber | RequestID=%s | could not find user by phone number: %s", reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if existPhoneNumber != nil && existPhoneNumber.ID == int(reqData.UserID) {
// 		return exception.GetException(exception.USER_PHONE_NUMBER_ALREADY_USED)
// 	} else if existPhoneNumber != nil && existPhoneNumber.ID != int(reqData.UserID) {
// 		return exception.GetException(exception.USER_PHONE_NUMBER_ALREADY_EXIST)
// 	}

// 	rdb := redisd.GetDB()
// 	key := util.RedisCreateUpdatePhoneNumberUserOTP(reqData.PhoneNumber)
// 	keyFailure := util.RedisCreateFailureUpdatePhoneNumberUserOTP(reqData.PhoneNumber)

// 	failureCount, err := rdb.Get(ctx, keyFailure).Result()
// 	if err != nil && err != redis.Nil {
// 		logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if failureCount != "" {
// 		failureCountInt, err := strconv.Atoi(failureCount)
// 		if err != nil {
// 			logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 			return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 		}

// 		if failureCountInt >= config.Get().FailureLimitGeneral {
// 			logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | request exceeds - too many: %s", keyFailure, reqData.RequestID)
// 			return exception.GetException(exception.GENERAL_TOO_MANY_REQUEST)
// 		}
// 	}

// 	//Check if exist OTP
// 	_, err = rdb.Get(ctx, key).Result()
// 	if err == nil {
// 		logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | otp still in redis, someone trying abuse otp retry", key, reqData.RequestID)
// 		return exception.GetException(exception.AUTH_OTP_STILL_VALID)
// 	}

// 	//Create new OTP
// 	OTPNumber := util.GetRandomSixNumber()
// 	_, err = rdb.Set(ctx, key, OTPNumber, time.Second*300).Result()
// 	if err != nil {
// 		logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | could not store otp cashout to redis: %s", key, reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if config.Get().Env == enum.CONFIG_ENV_PROD {
// 		q := qiscus.SendOTPRequest{
// 			To:      reqData.PhoneNumber,
// 			OTPCode: OTPNumber,
// 		}

// 		if _, err = q.Do(ctx, reqData.RequestID); err != nil {
// 			rdb.Del(ctx, key).Result()
// 			logger.Infof("RequestUpdatePhoneNumber | %s | RequestID=%s | qiscus.SendOTPRequest: %s", key, reqData.RequestID, err.Error())
// 			return exception.GetExceptionCustomMessage(exception.GENERAL_INTERNAL_SERVER_ERROR, err.Error())
// 		}
// 	}

// 	return nil
// }

// // UpdatePhoneNumberOTP ...
// func (u *userLogic) UpdatePhoneNumberOTP(ctx context.Context, reqData *dto.UpdatePhoneNumberOTP, tx *sql.Tx) (*int64, *exception.Exception) {

// 	var err error

// 	err = reqData.Validate()
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | %#v | error when validate request data: %s", reqData.RequestID, reqData, err.Error())
// 		return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	detailUser, err := u.repositoryUser.FindByID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return nil, exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	existPhoneNumber, err := u.repositoryUser.FindByPhoneNumber(ctx, reqData.RequestID, reqData.PhoneNumber)
// 	if err != nil && err != sql.ErrNoRows {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | could not find user by phone number: %s", reqData.RequestID, err.Error())
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if existPhoneNumber != nil && existPhoneNumber.ID == int(reqData.UserID) {
// 		return nil, exception.GetException(exception.USER_PHONE_NUMBER_ALREADY_USED)
// 	} else if existPhoneNumber != nil && existPhoneNumber.ID != int(reqData.UserID) {
// 		return nil, exception.GetException(exception.USER_PHONE_NUMBER_ALREADY_EXIST)
// 	}

// 	rdb := redisd.GetDB()
// 	key := util.RedisCreateUpdatePhoneNumberUserOTP(reqData.PhoneNumber)
// 	keyFailure := util.RedisCreateFailureUpdatePhoneNumberUserOTP(reqData.PhoneNumber)

// 	otpVal, err := rdb.Get(ctx, key).Result()
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not get otp from redis: %s", key, reqData.RequestID, err.Error())
// 		return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_NOT_FOUND), "otp expired, please request new otp")
// 	}
// 	if config.Get().Env == enum.CONFIG_ENV_PROD {
// 		if otpVal != reqData.OTP {

// 			failureCount, err := rdb.Get(ctx, keyFailure).Result()
// 			if err != nil && err != redis.Nil {
// 				logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 				return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 			}

// 			if failureCount != "" {
// 				failureCountInt, err := strconv.Atoi(failureCount)
// 				if err != nil {
// 					logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}

// 				if failureCountInt >= config.Get().FailureLimitGeneral {
// 					logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | request exceeds - too many: %s", keyFailure, reqData.RequestID)
// 					return nil, exception.GetException(exception.GENERAL_TOO_MANY_REQUEST)
// 				}

// 				_, err = rdb.Incr(ctx, keyFailure).Result()
// 				if err != nil {
// 					logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not increment failure to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}
// 			} else {
// 				_, err = rdb.Set(ctx, keyFailure, "1", time.Hour).Result()
// 				if err != nil {
// 					logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not set failure otp to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}
// 			}

// 			logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | OTP is not same, OTP_FROM_REQ=%s OTP_FROM_REDIS=%s", reqData.RequestID, reqData.OTP, otpVal)
// 			return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), "otp not matched")
// 		}
// 	} else {
// 		if reqData.OTP != "111111" {
// 			logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | OTP is not same, OTP_FROM_REQ=%s DEFAULT_OTP=%s err=%s", reqData.RequestID, reqData.OTP, "111111")
// 			return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), "otp not matched")
// 		}
// 	}

// 	rowAffected, err := u.repositoryUser.UpdatePhoneByID(ctx, reqData.RequestID, &model.User{
// 		ID:          int(reqData.UserID),
// 		PhoneNumber: reqData.PhoneNumber,
// 	}, tx)
// 	if err != nil && err != sql.ErrNoRows {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | error when repositoryUser.UpdatePhoneByID: %s", reqData.RequestID, err.Error())
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	// remove key value from redis when success
// 	_, err = rdb.Del(ctx, key, keyFailure).Result()
// 	if err != nil {
// 		logger.Infof("UpdatePhoneNumberOTP | %s | RequestID=%s | could not delete key on redis: %s", key, reqData.RequestID, err.Error())
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	// Call Logic Analytics for Update Store to Clever Tap
// 	detailUser.PhoneNumber = reqData.PhoneNumber
// 	ra := u.logicAnalytic.CleverTapUpdateUser(ctx, reqData.RequestID, *detailUser)
// 	if ra {
// 		logger.Infof("UpdatePhoneNumberOTP | RequestID=%s | could not cleverTapUpdateStore: %s", reqData.RequestID, ra)
// 	}

// 	return &rowAffected, nil
// }

// func (u *userLogic) PINCheck(ctx context.Context, reqData *dto.PINCheckRequest) (*bool, *exception.Exception) {
// 	var err error

// 	err = reqData.Validate()
// 	if err != nil {
// 		logger.Infof("PINCheck | RequestID=%s | %#v | error when validate request data: %s", reqData.RequestID, reqData, err.Error())
// 		return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	detailUser, err := u.repositoryUser.FindByID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("PINCheck | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return nil, exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if detailUser.PIN == nil {
// 		return nil, exception.GetException(exception.AUTH_PIN_NOT_SET)
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(*detailUser.PIN), []byte(reqData.PIN))
// 	if err != nil {
// 		return nil, exception.GetException(exception.AUTH_PIN_NOT_MATCH)
// 	}
// 	return util.NewBooleanPointer(true), nil
// }

// // DeleteAccountOTP ...
// func (u *userLogic) DeleteAccountOTP(ctx context.Context, reqData *dto.DeleteAccountOTPRequest) *exception.Exception {

// 	userDetail, err := u.repositoryUser.FindByID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("DeleteAccountOTP | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if userDetail.IsDelete != nil && *userDetail.IsDelete {
// 		return exception.GetException(exception.USER_STATUS_DELETE)
// 	}

// 	rdb := redisd.GetDB()
// 	key := util.RedisCreateDeleteAccountOTP(userDetail.PhoneNumber)
// 	keyFailure := util.RedisCreateFailureDeleteAccountOTP(userDetail.PhoneNumber)

// 	failureCount, err := rdb.Get(ctx, keyFailure).Result()
// 	if err != nil && err != redis.Nil {
// 		logger.Infof("DeleteAccountOTP | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if failureCount != "" {
// 		failureCountInt, err := strconv.Atoi(failureCount)
// 		if err != nil {
// 			logger.Infof("DeleteAccountOTP | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 			return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 		}

// 		if failureCountInt >= config.Get().FailureLimitGeneral {
// 			logger.Infof("DeleteAccountOTP | %s | RequestID=%s | request exceeds - too many: %s", keyFailure, reqData.RequestID)
// 			return exception.GetException(exception.GENERAL_TOO_MANY_REQUEST)
// 		}
// 	}

// 	//Check if exist OTP
// 	_, err = rdb.Get(ctx, key).Result()
// 	if err == nil {
// 		logger.Infof("DeleteAccountOTP | %s | RequestID=%s | otp still in redis, someone trying abuse otp retry", key, reqData.RequestID)
// 		return exception.GetException(exception.AUTH_OTP_STILL_VALID)
// 	}

// 	//Create new OTP
// 	OTPNumber := util.GetRandomSixNumber()
// 	_, err = rdb.Set(ctx, key, OTPNumber, time.Second*300).Result()
// 	if err != nil {
// 		logger.Infof("DeleteAccountOTP | %s | RequestID=%s | could not store otp delete account to redis: %s", key, reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	if config.Get().Env == enum.CONFIG_ENV_PROD {
// 		q := qiscus.SendOTPRequest{
// 			To:      userDetail.PhoneNumber,
// 			OTPCode: OTPNumber,
// 		}

// 		if _, err = q.Do(ctx, reqData.RequestID); err != nil {
// 			rdb.Del(ctx, key).Result()
// 			logger.Infof("DeleteAccountOTP | %s | RequestID=%s | qiscus.SendOTPRequest: %s", key, reqData.RequestID, err.Error())
// 			return exception.GetExceptionCustomMessage(exception.GENERAL_INTERNAL_SERVER_ERROR, err.Error())
// 		}
// 	}

// 	return nil
// }

// func (u *userLogic) DeleteAccount(ctx context.Context, reqData *dto.DeleteAccountRequest, tx *sql.Tx) *exception.Exception {
// 	if err := reqData.Validate(); err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | %#v | error when validate request data: %s", reqData.RequestID, reqData, err.Error())
// 		return exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	detailUser, err := u.repositoryUser.FindByID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return exception.GetException(exception.USER_NOT_FOUND)
// 		}
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}
// 	if detailUser.IsDelete != nil && *detailUser.IsDelete {
// 		return exception.GetException(exception.USER_STATUS_DELETE)
// 	}

// 	rdb := redisd.GetDB()
// 	key := util.RedisCreateDeleteAccountOTP(detailUser.PhoneNumber)
// 	keyFailure := util.RedisCreateFailureDeleteAccountOTP(detailUser.PhoneNumber)

// 	otpVal, err := rdb.Get(ctx, key).Result()
// 	if err != nil {
// 		logger.Infof("DeleteAccount | %s | RequestID=%s | could not get otp from redis: %s", key, reqData.RequestID, err.Error())
// 		return exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_NOT_FOUND), "otp expired, please request new otp")
// 	}
// 	if config.Get().Env == enum.CONFIG_ENV_PROD {
// 		if otpVal != reqData.OTP {

// 			failureCount, err := rdb.Get(ctx, keyFailure).Result()
// 			if err != nil && err != redis.Nil {
// 				logger.Infof("DeleteAccount | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 				return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 			}

// 			if failureCount != "" {
// 				failureCountInt, err := strconv.Atoi(failureCount)
// 				if err != nil {
// 					logger.Infof("DeleteAccount | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}

// 				if failureCountInt >= config.Get().FailureLimitGeneral {
// 					logger.Infof("DeleteAccount | %s | RequestID=%s | request exceeds - too many: %s", keyFailure, reqData.RequestID)
// 					return exception.GetException(exception.GENERAL_TOO_MANY_REQUEST)
// 				}

// 				_, err = rdb.Incr(ctx, keyFailure).Result()
// 				if err != nil {
// 					logger.Infof("DeleteAccount | %s | RequestID=%s | could not increment failure to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}
// 			} else {
// 				_, err = rdb.Set(ctx, keyFailure, "1", time.Hour).Result()
// 				if err != nil {
// 					logger.Infof("DeleteAccount | %s | RequestID=%s | could not set failure otp to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 					return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 				}
// 			}

// 			logger.Infof("DeleteAccount | RequestID=%s | OTP is not same, OTP_FROM_REQ=%s OTP_FROM_REDIS=%s", reqData.RequestID, reqData.OTP, otpVal)
// 			return exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), "otp not matched")
// 		}
// 	} else {
// 		if reqData.OTP != "111111" {
// 			logger.Infof("DeleteAccount | RequestID=%s | OTP is not same, OTP_FROM_REQ=%s DEFAULT_OTP=%s err=%s", reqData.RequestID, reqData.OTP, "111111")
// 			return exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), "otp not matched")
// 		}
// 	}

// 	storesUser, err := u.repositoryStore.FindAllStoreByUserID(ctx, reqData.RequestID, reqData.UserID)
// 	if err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | could not user find by id: %s", reqData.RequestID, err.Error())
// 		if err == sql.ErrNoRows {
// 			return exception.GetException(exception.STORE_NOT_FOUND)
// 		}
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	//Check saldo digital
// 	for _, v := range storesUser {
// 		getDigitalBalanceFromXendit := xendit.CheckBalanceRequest{
// 			XenPlatformAccountID: v.XenPlatformAccountID,
// 		}
// 		xenditDigitalBalance, err := getDigitalBalanceFromXendit.Do(ctx, reqData.RequestID)
// 		if err != nil {
// 			logger.Infof("DeleteAccount | RequestID=%s | xendit.CheckBalanceRequest: %s", reqData.RequestID, err.Error())
// 			return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 		}
// 		if *xenditDigitalBalance > 0 {
// 			return exception.GetExceptionCustomMessage(exception.GENERAL_BAD_REQUEST, "Anda masih memiliki saldo digital")
// 		}
// 	}

// 	//Update current info users
// 	if _, err := u.repositoryUser.DeleteAccount(ctx, reqData.RequestID, &model.User{
// 		ID:          detailUser.ID,
// 		PhoneNumber: fmt.Sprintf("%d%s", detailUser.ID, util.GetRandomPhoneNumber()),
// 		Email:       nil,
// 	}, tx); err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | repositoryUser.DeleteAccount: %s", reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	//Create backup user delete
// 	if _, err := u.repositoryUserDelete.Create(ctx, reqData.RequestID, &model.UserDelete{
// 		UserID:      detailUser.ID,
// 		PhoneNumber: detailUser.PhoneNumber,
// 		Email:       detailUser.Email,
// 		Reason:      reqData.Reason,
// 		Questions:   reqData.Questions,
// 	}, tx); err != nil {
// 		logger.Infof("DeleteAccount | RequestID=%s | repositoryUserDelete.Create: %s", reqData.RequestID, err.Error())
// 		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	return nil
// }
