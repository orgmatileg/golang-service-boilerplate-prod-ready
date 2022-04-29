package delivery

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserDelivery interface {
	CreateUser(ctx context.Context) error
}

func UserDeliveryNew(lc fx.Lifecycle) UserDelivery {
	return &userDelivery{}
}

type userDelivery struct {
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
}

func (u *userDelivery) CreateUser(ctx context.Context) error {

	return nil
}
