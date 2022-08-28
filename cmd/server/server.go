package server

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/handler"
	"github.com/mariojuzar/go-user-auth/internal/infrastructures/mongodb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ServeHttp() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run the REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load()
			if err != nil {
				logrus.WithError(err).Fatalf("Config error: %v", cfg)
			}

			mongoDb := mongodb.MongoDbConn(context.Background(), *cfg)

			if err := handler.NewHttpHandler(cfg, mongoDb).Start(); err != nil {
				logrus.WithError(err).Fatalf("Error starting API")
			}
		},
	}
}
