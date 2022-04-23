package dto

// import (
// 	"errors"
// 	"golang_service/enum"
// 	"golang_service/model"
// 	"golang_service/util"

// 	validation "github.com/go-ozzo/ozzo-validation/v4"
// 	"github.com/go-ozzo/ozzo-validation/v4/is"
// )

// type CreateRequest struct {
// 	FullName    string
// 	PhoneNumber string
// 	Email       string
// 	Status      int
// }

// func (c *CreateRequest) Validate() error {
// 	if !util.IsValidPhoneNumber(c.PhoneNumber) {
// 		return errors.New(util.PhoneNumberNotValid)
// 	}
// 	if err := util.ValidationIsIntContainOnSlice("Status", c.Status, []int{
// 		int(enum.UserStatusNotVerified),
// 		int(enum.UserStatusVerified),
// 		int(enum.UserStatusBlocked),
// 	}); err != nil {
// 		return err
// 	}
// 	if err := util.ValidationLenStringLowerThan("Nama", 5, c.FullName); err != nil {
// 		return err
// 	}
// 	if c.Email != "" && validation.Validate(c.Email, is.Email) != nil {
// 		return errors.New(util.EmailNotValid)
// 	}
// 	return nil
// }

// type UpdateByIDRequest struct {
// 	ID          int64
// 	FullName    string
// 	PhoneNumber string
// 	Email       string
// 	Status      int
// }

// func (c *UpdateByIDRequest) Validate() error {
// 	if !util.IsValidPhoneNumber(c.PhoneNumber) {
// 		return errors.New(util.PhoneNumberNotValid)
// 	}

// 	if err := util.ValidationIsIntContainOnSlice("Status", c.Status, []int{
// 		int(enum.UserStatusNotVerified),
// 		int(enum.UserStatusVerified),
// 		int(enum.UserStatusBlocked),
// 	}); err != nil {
// 		return err
// 	}

// 	if err := util.ValidationLenStringLowerThan("Nama", 5, c.FullName); err != nil {
// 		return err
// 	}
// 	if c.Email != "" && validation.Validate(c.Email, is.Email) != nil {
// 		return errors.New(util.EmailNotValid)
// 	}
// 	return nil
// }

// type FindByIDResponse struct {
// 	model.User
// }

// type FindAllResponse []*model.User

// type RequestUpdatePhoneNumber struct {
// 	RequestID   string
// 	UserID      int64
// 	PhoneNumber string
// }

// func (c *RequestUpdatePhoneNumber) Validate() error {
// 	if !util.IsValidPhoneNumber(c.PhoneNumber) {
// 		return errors.New(util.PhoneNumberNotValid)
// 	}
// 	return nil
// }

// type UpdatePhoneNumberOTP struct {
// 	UserID      int64
// 	RequestID   string
// 	PhoneNumber string
// 	OTP         string
// }

// func (c *UpdatePhoneNumberOTP) Validate() error {
// 	if !util.IsValidPhoneNumber(c.PhoneNumber) {
// 		return errors.New(util.PhoneNumberNotValid)
// 	}
// 	return nil
// }

// type PINCheckRequest struct {
// 	UserID    int64
// 	RequestID string
// 	PIN       string
// }

// func (c *PINCheckRequest) Validate() error {
// 	if len(c.PIN) != 6 {
// 		return errors.New("pin harus berjumlah 6 karakter")
// 	}

// 	return nil
// }

// type DeleteAccountOTPRequest struct {
// 	UserID    int64
// 	RequestID string
// }

// type DeleteAccountRequest struct {
// 	RequestID string
// 	UserID    int64
// 	IsAgree   bool
// 	OTP       string
// 	Reason    *string
// 	Questions []int
// }

// func (c *DeleteAccountRequest) Validate() error {
// 	if !c.IsAgree {
// 		return errors.New("Anda harus menyetujui segala ketentuan penghapusan akun")
// 	}
// 	if len(c.OTP) != 6 {
// 		return errors.New("OTP harus berjumlah 6 karakter")
// 	}
// 	return nil
// }
