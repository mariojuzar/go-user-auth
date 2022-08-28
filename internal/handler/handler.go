package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/handler/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type MainHandlerHttp interface {
	Start() error
}

type httpServer struct {
	db  *mongo.Database
	cfg config.MainConfig
	api *controller.API
}

func (h *httpServer) Start() error {
	return h.api.Start()
}

func NewHttpHandler(cfg *config.MainConfig, db *mongo.Database) MainHandlerHttp {
	server := &httpServer{
		db:  db,
		cfg: *cfg,
		api: controller.NewAPI(fiber.New(), db, cfg),
	}
	return server
}
