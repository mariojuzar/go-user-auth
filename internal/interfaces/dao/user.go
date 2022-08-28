package dao

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/internal/domain/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) repository.UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (repo *UserRepository) FindById(ctx context.Context, userId primitive.ObjectID) (*model.User, error) {
	user := &model.User{}

	filter := bson.M{
		"_id":        userId,
		"is_deleted": false,
	}

	err := repo.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(ctx context.Context, user *model.User) error {
	if user == nil {
		return ErrNilParam
	}
	_, err := repo.collection.InsertOne(ctx, *user)
	if err != nil {
		logrus.WithError(err).Error("Failed to insert user")
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateById(ctx context.Context, user *model.UserUpdate, userId primitive.ObjectID) error {
	if user == nil {
		return ErrNilParam
	}
	filter := bson.M{
		"_id": userId,
	}

	update := bson.M{
		"$set": *user,
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logrus.WithError(err).Error("Failed to update user")
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(ctx context.Context, userId primitive.ObjectID) error {
	filter := bson.M{
		"_id": userId,
	}

	_, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete user")
		return err
	}
	return nil
}

func (repo *UserRepository) FindAll(ctx context.Context, page, size int) ([]model.User, error) {
	skip := page * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))

	result, err := repo.collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var users []model.User
	err = result.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}

	filter := bson.M{
		"username":   username,
		"is_deleted": false,
	}

	err := repo.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}
