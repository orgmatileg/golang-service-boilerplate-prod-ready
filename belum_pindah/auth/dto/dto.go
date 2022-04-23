package dto

import (
	"errors"
)

type AuthRefreshTokenRequest struct {
	RefreshToken string
	RequestID    string
}

func (c *AuthRefreshTokenRequest) Validate() error {
	if c.RefreshToken == "" {
		return errors.New("refresh token tidak boleh kosong")
	}
	return nil
}

type AuthRefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthLoginWithPINRequest struct {
	RequestID   string
	PhoneNumber string
	PIN         string
}

func (c *AuthLoginWithPINRequest) Validate() error {
	if len(c.PIN) != 6 {
		return errors.New("pin harus berjumlah 6 karakter")
	}
	return nil
}

type AuthLoginWithPINResponse struct {
	AccessToken  string
	RefreshToken string
}
