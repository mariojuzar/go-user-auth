package mongodb

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/pkg/constant"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

func MongoDbConn(ctx context.Context, config config.MainConfig) *mongo.Database {
	client := mongodbClient(ctx, config)

	db := client.Database(config.MongoDB.Database)

	logrus.Info("Connected to MongoDB!")

	return db
}

func mongodbClient(ctx context.Context, config config.MainConfig) *mongo.Client {
	cred := getCredential(config.MongoDB)
	clientOptions := options.Client().ApplyURI(GetMongoDbStringConnURI(config))

	if cred != nil {
		clientOptions = clientOptions.SetAuth(*cred)
	}

	newClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logrus.Fatal(err)
	}

	err = newClient.Ping(ctx, nil)

	if err != nil {
		logrus.Fatal(err)
	}

	return newClient
}

func GetMongoDbStringConnURI(config config.MainConfig) string {
	cfgDB := config.MongoDB
	return "mongodb://" + getUserPass(cfgDB) + getHost(cfgDB) + "/" + cfgDB.Database
}

func getCredential(cfgDB config.MongoDBConfig) *options.Credential {
	if getUserPass(cfgDB) != constant.EmptyString {
		return &options.Credential{
			Username:   cfgDB.User,
			Password:   cfgDB.Password,
			AuthSource: cfgDB.Database,
		}
	}
	return nil
}

func getUserPass(cfgDB config.MongoDBConfig) string {
	if cfgDB.User != constant.EmptyString && cfgDB.Password != constant.EmptyString {
		return cfgDB.User + ":" + cfgDB.Password + "@"
	}
	return constant.EmptyString
}

func getHost(cfgDB config.MongoDBConfig) string {
	// host already contain port
	if strings.Contains(cfgDB.Host, ":") {
		return cfgDB.Host
	}
	return cfgDB.Host + ":" + strconv.FormatUint(uint64(cfgDB.Port), 10)
}
