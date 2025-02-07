package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
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

func (tr *taskRepository) FetchByOwner(c context.Context, owner string, page, size string) ([]domains.Task, error) {
	collection := tr.database.Collection(tr.collection)
	convPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	convSize, err := strconv.Atoi(size)
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetSkip(int64(convPage * convSize)).SetLimit(int64(convSize))
	oid, err := primitive.ObjectIDFromHex(owner)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(c, bson.M{"owner": oid}, opts)
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
	oid, err := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(c, bson.M{"_id": oid}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *taskRepository) EditById(c context.Context, id string, task *domains.Task) error {
	collection := tr.database.Collection(tr.collection)
	oid, err := primitive.ObjectIDFromHex(id)
	_, err = collection.UpdateOne(
		c,
		bson.M{"_id": oid},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{"title", task.Title},
					{"description", task.Description},
					{"isCompleted", task.IsCompleted},
					{"updatedAt", task.UpdatedAt},
				},
			},
		},
	)
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
