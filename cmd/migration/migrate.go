package migration

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	infraMongo "github.com/mariojuzar/go-user-auth/internal/infrastructures/mongodb"
	"github.com/mariojuzar/go-user-auth/internal/interfaces/dao"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func RunMigration() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run migration database",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load()
			if err != nil {
				logrus.WithError(err).Fatalf("Config error: %v", cfg)
			}

			mongoDb := infraMongo.MongoDbConn(context.Background(), *cfg)

			cfgMongo := mongodb.Config{
				DatabaseName: cfg.MongoDB.Database,
			}
			driver, err := mongodb.WithInstance(mongoDb.Client(), &cfgMongo)
			if err != nil {
				logrus.WithError(err).Fatal("Failed connect mongo driver")
			}

			migration, err := migrate.NewWithDatabaseInstance("file://database/migrations", cfg.MongoDB.Database, driver)
			if err != nil {
				logrus.WithError(err).Fatal("Failed connect migration driver")
			}

			if err := migration.Up(); err != nil {
				logrus.WithError(err).Error("Error running migration schema")
			}

			userRepo := dao.NewUserRepository(mongoDb)
			permRepo := dao.NewRolePermissionRepository(mongoDb)

			rolePermissions := []model.RolePermission{
				model.RolePermission{
					Role:       model.BasicUser,
					Permission: "read-my-user",
					AccessAPI:  "/v1/users/me",
					Method:     http.MethodGet,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "read-my-user",
					AccessAPI:  "/v1/users/me",
					Method:     http.MethodGet,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "read-user",
					AccessAPI:  "/v1/users/user/:user_id",
					Method:     http.MethodGet,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "read-users",
					AccessAPI:  "/v1/users",
					Method:     http.MethodGet,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "create-users",
					AccessAPI:  "/v1/users",
					Method:     http.MethodPost,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "update-users",
					AccessAPI:  "/v1/users",
					Method:     http.MethodPatch,
				},
				model.RolePermission{
					Role:       model.UserAdmin,
					Permission: "delete-user",
					AccessAPI:  "/v1/users/:user_id",
					Method:     http.MethodDelete,
				},
			}

			passwordAdmin, _ := bcrypt.GenerateFromPassword([]byte("superadmin"), 5)

			// seed database

			// create super admin
			superadmin := model.User{
				Username:    "superadmin",
				FirstName:   "super",
				LastName:    "admin",
				Password:    string(passwordAdmin),
				Role:        model.SuperAdmin,
				Permissions: []string{},
				CreatedAt:   time.Now(),
				CreatedBy:   "system",
				UpdatedAt:   time.Now(),
				UpdatedBy:   "system",
				IsDeleted:   false,
			}
			_ = userRepo.Create(context.Background(), &superadmin)

			// seed default permission
			for _, rolePermission := range rolePermissions {
				_ = permRepo.InsertRolePermission(context.Background(), rolePermission)
			}

			logrus.Infof("Complete migrate and seed database")
		},
	}
}
