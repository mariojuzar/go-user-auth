package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/mariojuzar/go-user-auth/docs"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/handler/middleware"
	"github.com/mariojuzar/go-user-auth/internal/interfaces/dao"
	"github.com/mariojuzar/go-user-auth/internal/usecases"
	"github.com/mariojuzar/go-user-auth/internal/usecases/auth"
	"github.com/mariojuzar/go-user-auth/internal/usecases/user"
	"github.com/swaggo/fiber-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	db     *mongo.Database
	engine *fiber.App
	cfg    config.MainConfig
	userUc usecases.UserUseCase
	authUc usecases.AuthUseCase
	mw     middleware.Interceptor
}

func NewAPI(engine *fiber.App, db *mongo.Database, cfg *config.MainConfig) *API {
	api := &API{
		db:     db,
		engine: engine,
		cfg:    *cfg,
	}
	api.setup()
	api.setupRoutes()
	return api
}

func (a *API) Start() error {
	return a.engine.Listen(fmt.Sprintf("%s:%d", a.cfg.ApiConfig.Host, a.cfg.ApiConfig.Port))
}

func (a *API) setup() {
	userRepo := dao.NewUserRepository(a.db)
	rpRepo := dao.NewRolePermissionRepository(a.db)
	a.userUc = user.New(&user.Opts{
		UserRepo:           userRepo,
		RolePermissionRepo: rpRepo,
	})
	a.authUc = auth.New(&auth.Opts{
		UserRepo:   userRepo,
		AuthConfig: a.cfg.AuthConfig,
	})

	a.mw = middleware.NewMiddleware(&middleware.Opts{
		CfgAuth:            a.cfg.AuthConfig,
		RolePermissionRepo: rpRepo,
	})
}

func (a *API) setupRoutes() {
	router := a.engine

	router.Use(recover.New())

	router.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	router.Get("/docs/*", fiberSwagger.FiberWrapHandler())

	v1Router := router.Group("/v1")
	userRouter := v1Router.Group("/users")
	// using Auth middleware for authentication
	userRouter.Get("/user/:user_id", a.mw.Auth(a.FindUserById))
	userRouter.Get("/me", a.mw.Auth(a.FindMyUser))
	userRouter.Get("/", a.mw.Auth(a.FindAllUser))
	userRouter.Post("/", a.mw.Auth(a.CreateUser))
	userRouter.Patch("/", a.mw.Auth(a.UpdateUser))
	userRouter.Delete("/:user_id", a.mw.Auth(a.DeleteUser))

	authRouter := v1Router.Group("/auth")
	authRouter.Post("/login", a.LoginUser)
	authRouter.Post("/refresh", a.RefreshToken)
}
