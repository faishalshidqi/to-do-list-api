package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"todo-list-api/domains"
	"todo-list-api/infrastructures/mongo"
)

type refreshAuthnRepository struct {
	database   mongo.Database
	collection string
}

func (rr *refreshAuthnRepository) Add(c context.Context, refreshRequest domains.RefreshAuthnRequest) error {
	collection := rr.database.Collection(rr.collection)
	_, err := collection.InsertOne(c, refreshRequest)
	return err
}

func (rr *refreshAuthnRepository) Fetch(c context.Context, token string) (domains.RefreshAuthnRequest, error) {
	collection := rr.database.Collection(rr.collection)
	tokenRequest := domains.RefreshAuthnRequest{}
	err := collection.FindOne(c, bson.M{"refreshToken": token}).Decode(&tokenRequest)
	if err != nil {
		return domains.RefreshAuthnRequest{}, err
	}
	return tokenRequest, nil
}

func (rr *refreshAuthnRepository) DeleteToken(c context.Context, token string) error {
	collection := rr.database.Collection(rr.collection)
	refreshRequest := domains.RefreshAuthnRequest{}
	err := collection.FindOne(c, bson.M{"refreshToken": token}).Decode(&refreshRequest)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"refreshToken": refreshRequest.RefreshToken})
	if err != nil {
		return err
	}
	return nil
}

func NewRefreshAuthnRepository(db mongo.Database, collection string) domains.RefreshAuthnRepository {
	return &refreshAuthnRepository{
		database:   db,
		collection: collection,
	}
}
