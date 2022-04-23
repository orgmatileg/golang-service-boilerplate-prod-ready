package logic

// import (
// 	"context"
// 	"database/sql"
// 	"golang_service/config"
// 	"golang_service/exception"
// 	"golang_service/module/auth/dto"
// 	"golang_service/pkg/jwtd"
// 	"golang_service/pkg/logger"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// type IAuthLogic interface {
// 	// Login with PIN
// 	AuthLoginWithPIN(ctx context.Context, reqData *dto.AuthLoginWithPINRequest) (*dto.AuthLoginWithPINResponse, *exception.Exception)
// 	// Refresh Token
// 	AuthRefreshToken(ctx context.Context, reqData *dto.AuthRefreshTokenRequest) (*dto.AuthRefreshTokenResponse, *exception.Exception)
// }

// type authLogic struct {
// 	db             *sql.DB
// 	contextTimeout time.Duration
// }

// // New ...
// func New(ru repositoryUser.IUserRepository, timeout time.Duration, db *sql.DB) IAuthLogic {
// 	return &authLogic{
// 		db:             db,
// 		contextTimeout: timeout,
// 	}
// }

// func (u *authLogic) AuthRefreshToken(ctx context.Context, reqData *dto.AuthRefreshTokenRequest) (*dto.AuthRefreshTokenResponse, *exception.Exception) {

// 	err := reqData.Validate()
// 	if err != nil {
// 		return nil, exception.GetExceptionCustomMessage(exception.GENERAL_BAD_REQUEST, err.Error())
// 	}

// 	isValid, err := jwtd.IsValidToken(reqData.RefreshToken, config.Get().JWT.SecretRefreshToken)
// 	if err != nil {
// 		logger.Infof("AuthRefreshToken | RequestID=%s | could not validate token refresh: %s", reqData.RequestID, err.Error())

// 		errJwt, _ := err.(*jwt.ValidationError)
// 		if errJwt.Errors == jwt.ValidationErrorExpired {
// 			return nil, exception.GetException(exception.AUTH_REFRESH_TOKEN_EXPIRED)
// 		}
// 		return nil, exception.GetExceptionCustomMessage(exception.GENERAL_BAD_REQUEST, err.Error())
// 	}

// 	if !isValid {
// 		logger.Errorf("AuthRefreshToken | RequestID=%s | Somehow token is not valid")
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	claim, err := jwtd.ParseClaim(reqData.RefreshToken, config.Get().JWT.SecretRefreshToken)
// 	if err != nil {
// 		return nil, exception.GetExceptionCustomMessage(exception.GENERAL_BAD_REQUEST, err.Error())
// 	}

// 	userID := claim.Data.UserID

// 	res := dto.AuthRefreshTokenResponse{}

// 	t := time.Now()
// 	accessToken, err := jwtd.GenerateToken(jwtd.Claim{
// 		Data: jwtd.ClaimData{
// 			UserID: userID,
// 		},
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:    config.Get().AppURL,
// 			Audience:  "Owner",
// 			IssuedAt:  t.Unix(),
// 			ExpiresAt: t.Add(config.Get().Auth.ExpireAccessTokenDuration).Unix(),
// 		},
// 	}, config.Get().JWT.SecretAccessToken)
// 	if err != nil {
// 		logger.Infof("AuthRefreshToken | RequestID=%s | could not generate access token: %s", reqData.RequestID, err.Error())
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	res.AccessToken = accessToken

// 	refreshToken, err := jwtd.GenerateToken(jwtd.Claim{
// 		Data: jwtd.ClaimData{
// 			UserID: userID,
// 		},
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:    config.Get().AppURL,
// 			Audience:  "Owner",
// 			IssuedAt:  t.Unix(),
// 			ExpiresAt: t.Add(config.Get().Auth.ExpireRefreshTokenDuration).Unix(),
// 		},
// 	}, config.Get().JWT.SecretRefreshToken)
// 	if err != nil {
// 		logger.Infof("AuthRefreshToken | RequestID=%s | could not generate refresh token: %s", reqData.RequestID, err.Error())
// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	}

// 	res.RefreshToken = refreshToken

// 	return &res, nil
// }

// // AuthLoginWithPIN ...
// func (u *authLogic) AuthLoginWithPIN(ctx context.Context, reqData *dto.AuthLoginWithPINRequest) (*dto.AuthLoginWithPINResponse, *exception.Exception) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
// 	defer cancel()

// 	if err := reqData.Validate(); err != nil {
// 		logger.Infof("AuthLoginWithPIN | RequestID=%s | %#v | error when validate request data: %s", reqData.RequestID, reqData, err.Error())
// 		return nil, exception.GetExceptionCustomMessage(exception.Code(exception.GENERAL_BAD_REQUEST), err.Error())
// 	}

// 	// currentUserModel, err := u.repositoryUser.FindByPhoneNumber(ctx, reqData.RequestID, reqData.PhoneNumber)
// 	// if err != nil {
// 	// 	logger.Infof("AuthLoginWithPIN | RequestID=%s | could not user find by phone: %s", reqData.PhoneNumber, err.Error())
// 	// 	if err == sql.ErrNoRows {
// 	// 		return nil, exception.GetException(exception.USER_NOT_FOUND)
// 	// 	}
// 	// 	return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// }

// 	// if currentUserModel.PIN == nil {
// 	// 	return nil, exception.GetException(exception.AUTH_PIN_NOT_SET)
// 	// }

// 	// if currentUserModel.Status != nil && *currentUserModel.Status == enum.UserStatusNotVerified {
// 	// 	logger.Infof("AuthLoginWithPIN | RequestID=%s | phone number found but need verified %s", reqData.RequestID, reqData.PhoneNumber)
// 	// 	return nil, exception.GetException(exception.USER_STATUS_NOT_ACTIVE)
// 	// }

// 	// rdb := redisd.GetDB()
// 	// keyFailure := util.RedisCreateKeyFailureLoginWithPIN(reqData.PhoneNumber)

// 	// err = bcrypt.CompareHashAndPassword([]byte(*currentUserModel.PIN), []byte(reqData.PIN))
// 	// if err != nil {

// 	// 	failureCount, err := rdb.Get(ctx, keyFailure).Result()
// 	// 	if err != nil && err != redis.Nil {
// 	// 		logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 	// 		return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// 	}

// 	// 	if failureCount != "" {
// 	// 		failureCountInt, err := strconv.Atoi(failureCount)
// 	// 		if err != nil {
// 	// 			logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | could not get otp from redis: %s", keyFailure, reqData.RequestID, err.Error())
// 	// 			return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// 		}

// 	// 		if failureCountInt >= config.Get().FailureLimitGeneral {
// 	// 			logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | request exceeds - too many: %s", keyFailure, reqData.RequestID)
// 	// 			return nil, exception.GetException(exception.GENERAL_TOO_MANY_REQUEST)
// 	// 		}

// 	// 		_, err = rdb.Incr(ctx, keyFailure).Result()
// 	// 		if err != nil {
// 	// 			logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | could not increment failure to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 	// 			return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// 		}

// 	// 	} else {
// 	// 		_, err = rdb.Set(ctx, keyFailure, "1", time.Hour).Result()
// 	// 		if err != nil {
// 	// 			logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | could not set failure otp to redis: %s", keyFailure, reqData.RequestID, err.Error())
// 	// 			return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// 		}
// 	// 	}

// 	// 	return nil, exception.GetException(exception.AUTH_PIN_NOT_MATCH)
// 	// }

// 	// // Generate Auth & Refresh Token
// 	// var resp dto.AuthLoginWithPINResponse

// 	// t := time.Now()
// 	// accessToken, err := jwtd.GenerateToken(jwtd.Claim{
// 	// 	Data: jwtd.ClaimData{
// 	// 		UserID: currentUserModel.ID,
// 	// 	},
// 	// 	StandardClaims: jwt.StandardClaims{
// 	// 		Issuer:    config.Get().AppURL,
// 	// 		Audience:  "Owner",
// 	// 		IssuedAt:  t.Unix(),
// 	// 		ExpiresAt: t.Add(config.Get().Auth.ExpireAccessTokenDuration).Unix(),
// 	// 	},
// 	// }, config.Get().JWT.SecretAccessToken)
// 	// if err != nil {
// 	// 	logger.Infof("AuthLoginWithPIN | RequestID=%s | could not generate access token: %s", reqData.RequestID, err.Error())
// 	// 	return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// }

// 	// resp.AccessToken = accessToken

// 	// refreshToken, err := jwtd.GenerateToken(jwtd.Claim{
// 	// 	Data: jwtd.ClaimData{
// 	// 		UserID: currentUserModel.ID,
// 	// 	},
// 	// 	StandardClaims: jwt.StandardClaims{
// 	// 		Issuer:    config.Get().AppURL,
// 	// 		Audience:  "Owner",
// 	// 		IssuedAt:  t.Unix(),
// 	// 		ExpiresAt: t.Add(config.Get().Auth.ExpireRefreshTokenDuration).Unix(),
// 	// 	},
// 	// }, config.Get().JWT.SecretRefreshToken)
// 	// if err != nil {
// 	// 	logger.Infof("AuthLoginWithPIN | RequestID=%s | could not generate refresh token: %s", reqData.RequestID, err.Error())
// 	// 	return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// }

// 	// resp.RefreshToken = refreshToken

// 	// // remove key value from redis when success
// 	// _, err = rdb.Del(ctx, keyFailure).Result()
// 	// if err != nil {
// 	// 	logger.Infof("AuthLoginWithPIN | %s | RequestID=%s | could not delete key on redis: %s", keyFailure, reqData.RequestID, err.Error())
// 	// 	return nil, exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
// 	// }

// 	return nil, nil

// }
