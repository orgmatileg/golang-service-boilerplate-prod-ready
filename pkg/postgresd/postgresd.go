package postgresd

import (
	"fmt"
	"golang_service/config"
	"golang_service/pkg/logger"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormMaster *gorm.DB

type GormSlave *gorm.DB

// GetGormDB get db
func ProvideGormMaster(lc fx.Lifecycle) *gorm.DB {

	// if gormDB != nil {
	// 	return gormDB
	// }

	// err := createConnectionGorm()
	// if err != nil {
	// 	logger.Fatalf("postgresd | GetGormDB() | could not createConnection(): %s", err.Error())
	// }

	// return gormDB
	return nil
}

// createConnectionGormMaster open new connection Postgres
func createConnectionGormMaster() error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		config.Get().PostgreMaster.Host,
		config.Get().PostgreMaster.Port,
		config.Get().PostgreMaster.Username,
		config.Get().PostgreMaster.Password,
		config.Get().PostgreMaster.DBName,
		config.Get().PostgreMaster.SSLMode,
	)

	gormDBCon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("could not open connection to postgres database: %s", err.Error())
		return err
	}

	return nil
}
