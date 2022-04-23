package delivery

import (
	"context"
	"database/sql"

	"go.uber.org/fx"
)

type UserDelivery interface {
	CreateUser(ctx context.Context) error
}

func UserDeliveryNew(lc fx.Lifecycle) UserDelivery {
	return &userDelivery{}
}

type userDelivery struct {
	dbMaster *sql.DB
	dbSlave  *sql.DB
}

func (u *userDelivery) CreateUser(ctx context.Context) error {

	return nil
}
