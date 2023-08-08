package db

import (
	"context"
	"fmt"
	"rest-api/internal/user"
	"rest-api/pkg/logging"

	"github.com/pelletier/go-toml/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// Create implements user.Storage.
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	d.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (*db) FindOne(ctx context.Context, id string) (u user.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to onjectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user (id: %s) from DB due to error: %v", id, err)
	}
	
	return u, nil
}

func (*db) Update(ctx context.Context, user user.User) error {
	panic("unimplemented")
}

func (*db) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
