package server

import (
	"github.com/mariojuzar/go-user-auth/internal/config"
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

			logrus.Infoln("Server called")
		},
	}
}
