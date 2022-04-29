package restapi

import (
	// handlerAuth "golang_service/belum_pindah/auth/handler"

	"github.com/labstack/echo/v4"
)

func setRouterForUser(e *echo.Echo) {

	// handlerAuth.SetRouter(
	// e,
	// logicAuth.New(
	// 	repositoryUser.New(postgresd.GetDB()),
	// ),
	// postgresd.GetDB(),
	// )

	// handlerUser.SetRouter(
	// 	e,
	// 	logicUser.New(
	// 		repositoryUser.New(postgresd.GetDB()),
	// 		logicAnalytic.New(
	// 			repositoryStore.New(postgresd.GetDB()),
	// 			time.Second*60,
	// 			postgresd.GetDB(),
	// 		),
	// 		repositoryStore.New(postgresd.GetDB()),
	// 		time.Second*60,
	// 		postgresd.GetDB(),
	// 	),
	// 	postgresd.GetDB(),
	// 	middlewareIsUserAuthorized,
	// )

}
