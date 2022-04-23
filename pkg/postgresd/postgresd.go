package postgresd

import (
	"database/sql"
	"fmt"
	"golang_service/config"
	"golang_service/pkg/logger"
	"time"

	"github.com/jackc/pgx/v4/stdlib"
	sqlDatadogTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormDatadogTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// GetGormDB get db
func GetGormDB() *gorm.DB {

	if gormDB != nil {
		return gormDB
	}

	err := createConnectionGorm()
	if err != nil {
		logger.Fatalf("postgresd | GetGormDB() | could not createConnection(): %s", err.Error())
	}

	return gormDB
}

// GetDB get db
func GetDB() *sql.DB {

	if gormDB != nil {

		dbCon, err := gormDB.DB()
		if err != nil {
			logger.Fatalf("postgresd | GetGormDB() | could not createConnection(): %s", err.Error())
		}

		return dbCon
	}

	err := createConnectionGorm()
	if err != nil {
		logger.Fatalf("postgresd | GetGormDB() | could not createConnection(): %s", err.Error())
	}

	dbCon, err := gormDB.DB()
	if err != nil {
		logger.Fatalf("postgresd | GetGormDB() | could not createConnection(): %s", err.Error())
	}

	return dbCon
}

// createConnectionGorm open new connection Postgres
func createConnectionGorm() error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		config.Get().Postgres.Host,
		config.Get().Postgres.Port,
		config.Get().Postgres.Username,
		config.Get().Postgres.Password,
		config.Get().Postgres.DBName,
		config.Get().Postgres.SSLMode,
	)

	sqlDatadogTrace.Register("pgx", &stdlib.Driver{},
		sqlDatadogTrace.WithServiceName("app_backend_postgres_master"),
		sqlDatadogTrace.WithAnalyticsRate(1),
		sqlDatadogTrace.WithAnalytics(true),
	)

	sqlCon, err := sqlDatadogTrace.Open("pgx", dsn)
	if err != nil {
		logger.Errorf("could not connect to database: %s", err.Error())
		return err
	}

	gormDBCon, err := gormDatadogTrace.Open(
		postgres.New(
			postgres.Config{
				Conn: sqlCon,
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		logger.Errorf("could not ping postgres database: %s", err.Error())
		return err
	}

	sqldb, err := gormDBCon.DB()
	if err != nil {
		logger.Errorf("could not ping postgres database: %s", err.Error())
		return err
	}

	err = sqldb.Ping()
	if err != nil {
		logger.Errorf("could not ping postgres database: %s", err.Error())
		return err
	}

	sqldb.SetMaxOpenConns(100)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxIdleTime(300 * time.Second)
	sqldb.SetConnMaxLifetime(time.Duration(300 * time.Second))

	logger.Infof("database postgres: Connected!")

	gormDB = gormDBCon
	return nil
}
