package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"my-judgment/infrastructure/auth/jwttoken"
	"my-judgment/infrastructure/customcontext"
	"my-judgment/infrastructure/middleware"
	"my-judgment/infrastructure/rdb/persistence"
	"my-judgment/interfaces/handler/tokenhandler"
	"my-judgment/interfaces/handler/userhandler"
	"my-judgment/usecase/tokenusecase"
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

	// トークン系
	tokenRouterForWeb(mjRouterWithoutTokenAuth)

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

func tokenRouterForWeb(mjRouter *echo.Group) {
	tokenService := jwttoken.NewJwtTokenService()
	userRepository := persistence.NewUserRepository()

	// アクセストークン取得
	generateWebTokenUsecase := tokenusecase.NewGenerateWebTokenUsecase(tokenService, userRepository)
	generateWebTokenHandler := tokenhandler.NewGenerateWebTokenHandler(generateWebTokenUsecase)

	mjRouter.GET("/token", func(c echo.Context) error {
		return generateWebTokenHandler.GenerateWebToken(c.(*customcontext.CustomContext))
	})
}

func userRouterForWeb(mjRouter *echo.Group) {
	userRepository := persistence.NewUserRepository()

	// ユーザー新規登録
	createUserUsecase := userusecase.NewCreateUserUsecase(userRepository)
	createUserHandler := userhandler.NewCreateUserHandler(createUserUsecase)

	mjRouter.POST("/users", func(c echo.Context) error {
		return createUserHandler.CreateUser(c.(*customcontext.CustomContext))
	})

	// ユーザー単体取得
	fetchUserUsecase := userusecase.NewFetchUserUsecase(userRepository)
	fetchUserHandler := userhandler.NewFetchUserHandler(fetchUserUsecase)

	mjRouter.GET("/users/:userID", func(c echo.Context) error {
		return fetchUserHandler.FetchUser(c.(*customcontext.CustomContext))
	})

	// ユーザー情報更新
	updateUserUsecase := userusecase.NewUpdateUserUsecase(userRepository)
	updateUserHandler := userhandler.NewUpdateUserHandler(updateUserUsecase)

	mjRouter.PUT("/users/:userID", func(c echo.Context) error {
		return updateUserHandler.UpdateUser(c.(*customcontext.CustomContext))
	})
}
