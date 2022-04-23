module golang_service

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/jackc/pgx/v4 v4.15.0
	github.com/labstack/echo/v4 v4.7.2
	github.com/labstack/gommon v0.3.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/xhit/go-str2duration/v2 v2.0.0
	go.uber.org/fx v1.17.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e // indirect
	golang.org/x/tools v0.1.7 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.37.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)
