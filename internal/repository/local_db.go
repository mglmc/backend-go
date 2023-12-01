package repository

import (
	"backend-go/internal/model"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *LocalDB) ConnectDB() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := os.Getenv("MONGODB_URI")

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (db *LocalDB) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	collection := db.client.Database("test").Collection("tasks")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var task model.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (db *LocalDB) GetTaskByID(id int) (*model.Task, error) {
	var task model.Task
	collection := db.client.Database("test").Collection("tasks")
	err := collection.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (db *LocalDB) CreateTask(task *model.Task) error {
	collection := db.client.Database("test").Collection("tasks")
	_, err := collection.InsertOne(context.Background(), task)
	return err
}

func (db *LocalDB) UpdateTask(task *model.Task) error {
	collection := db.client.Database("test").Collection("tasks")
	_, err := collection.UpdateOne(context.Background(), bson.D{{"id", task.ID}}, bson.D{{"$set", task}})
	return err
}

func (db *LocalDB) DeleteTask(id int) error {
	collection := db.client.Database("test").Collection("tasks")
	_, err := collection.DeleteOne(context.Background(), bson.D{{"id", id}})
	return err
}
