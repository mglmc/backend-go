package repository

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := os.Getenv("MONGODB_URI")

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.Background())

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
}

// func GetAllTasks() ([]model.Task, error) {
// 	var tasks []model.Task
// 	collection := db.client.Database("test").Collection("tasks")
// 	cur, err := collection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cur.Close(context.Background())
// 	for cur.Next(context.Background()) {
// 		var task model.Task
// 		err := cur.Decode(&task)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil, err
// 	}
// 	return tasks, nil
// }

// func GetTaskByID(id int) (*model.Task, error) {
// 	var task model.Task
// 	collection := db.client.Database("test").Collection("tasks")
// 	err := collection.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&task)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &task, nil
// }

// func CreateTask(task *model.Task) error {
// 	collection := db.client.Database("test").Collection("tasks")
// 	_, err := collection.InsertOne(context.Background(), task)
// 	return err
// }

// func UpdateTask(task *model.Task) error {
// 	collection := db.client.Database("test").Collection("tasks")
// 	_, err := collection.UpdateOne(context.Background(), bson.D{{"id", task.ID}}, bson.D{{"$set", task}})
// 	return err
// }

// func DeleteTask(id int) error {
// 	collection := db.client.Database("test").Collection("tasks")
// 	_, err := collection.DeleteOne(context.Background(), bson.D{{"id", id}})
// 	return err
// }
