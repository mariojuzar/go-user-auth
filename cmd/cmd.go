package cmd

import (
	"github.com/mariojuzar/go-user-auth/cmd/migration"
	"github.com/mariojuzar/go-user-auth/cmd/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "User API",
}

func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	rootCmd.AddCommand(server.ServeHttp())
	rootCmd.AddCommand(migration.RunMigration())

	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}
