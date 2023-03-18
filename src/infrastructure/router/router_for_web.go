package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"my-judgment/infrastructure/customcontext"
	"my-judgment/infrastructure/middleware"
	"my-judgment/infrastructure/rdb/persistence"
	"my-judgment/interfaces/handler/userhandler"
	"my-judgment/usecase/userusecase"
)

func RouterForWebGroup(e *echo.Echo, db *gorm.DB) {
	// 認証用トークン不要なルート
	mjRouterWithoutTokenAuth := e.Group(
		"/mj",
		middleware.LogMiddleware(),
		middleware.CustomContext(),
		middleware.DBMiddleware(db),
		middleware.VerifyApiVersion(),
	)

	// ユーザー系
	userRouterForWeb(mjRouterWithoutTokenAuth)

	//// 認証用トークン必須なルート
	//mjRouter := e.Group(
	//	"/mj",
	//	middleware.LogMiddleware(),
	//	middleware.CustomContext(),
	//	middleware.DBMiddleware(db),
	//	middleware.VerifyApiVersion(),
	//	middleware.VerifyToken(),
	//)
}

func userRouterForWeb(mjRouter *echo.Group) {
	userRepository := persistence.NewUserRepository()

	// ユーザー新規登録
	createUserUsecase := userusecase.NewCreateUserUsecase(userRepository)
	createUserHandler := userhandler.NewCreateUserHandler(createUserUsecase)

	mjRouter.POST("/users", func(c echo.Context) error {
		return createUserHandler.CreateUser(c.(*customcontext.CustomContext))
	})
}
