package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo-list-api/domains"
	"todo-list-api/infrastructures/mongo"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func (tr *taskRepository) Add(c context.Context, task *domains.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	return err
}

func (tr *taskRepository) FetchByOwner(c context.Context, owner string) ([]domains.Task, error) {
	collection := tr.database.Collection(tr.collection)
	cursor, err := collection.Find(c, bson.M{"owner": owner}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	var tasks []domains.Task
	err = cursor.All(c, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) FetchById(c context.Context, id string) (*domains.Task, error) {
	collection := tr.database.Collection(tr.collection)
	task := domains.Task{}
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *taskRepository) EditById(c context.Context, id string, task *domains.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.UpdateOne(c, bson.M{"_id": id}, task)
	return err
}

func (tr *taskRepository) DeleteById(c context.Context, id string) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}

func (tr *taskRepository) MarkAsCompleted(c context.Context, id string) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.UpdateOne(c, bson.M{"_id": id}, bson.M{"isCompleted": true})
	return err
}

func (tr *taskRepository) FetchCompleted(c context.Context, owner string) ([]domains.Task, error) {
	collection := tr.database.Collection(tr.collection)
	opts := options.Find().SetSort(bson.D{{Key: "isCompleted", Value: 1}})
	cursor, err := collection.Find(c, bson.M{"owner": owner}, opts)
	if err != nil {
		return nil, err
	}
	tasks := make([]domains.Task, 0)
	err = cursor.All(c, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func NewTaskRepository(db mongo.Database, collection string) domains.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}
