package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo-list-api/domains"
	"todo-list-api/infrastructures/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (ur userRepository) Add(c context.Context, user *domains.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

func (ur userRepository) Fetch(c context.Context) ([]domains.User, error) {
	collection := ur.database.Collection(ur.collection)
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	users := make([]domains.User, 0)
	err = cursor.All(c, &users)
	if err != nil {
		return []domains.User{}, err
	}
	return users, nil
}

func (ur userRepository) GetByEmail(c context.Context, email string) (domains.User, error) {
	collection := ur.database.Collection(ur.collection)
	user := domains.User{}
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (ur userRepository) GetByID(c context.Context, id primitive.ObjectID) (domains.User, error) {
	collection := ur.database.Collection(ur.collection)
	user := domains.User{}
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func NewUserRepository(db mongo.Database, collection string) domains.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
