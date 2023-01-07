package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "server/db"
	todo "server/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = db.Connect()

func home() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var todos []primitive.M

	for curr.Next(context.Background()) {
		var todo primitive.M
		err := curr.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	defer curr.Close(context.Background())
	return todos
}

func createTodo(todo todo.Todo) {
	inserted, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", inserted.InsertedID)
}

func updateTodo(todoId string) {
	id, _ := primitive.ObjectIDFromHex(todoId)
	filter := bson.M{"_id": id}
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	completed, ok := result["completed"].(bool)
	if !ok {
		log.Fatal("Completed field is not a boolean")
	}
	result["completed"] = !completed
	updateDoc, err := bson.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	rawDoc := bson.Raw(updateDoc)
	update := bson.D{{"$set", rawDoc}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated a single document: ", updated.ModifiedCount)
}

func deleteTodo(todoId string) {
	id, _ := primitive.ObjectIDFromHex(todoId)
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted a single document: ", deleted.DeletedCount)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	todos := home()
	json.NewEncoder(w).Encode(todos)
}
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	var todo todo.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	createTodo(todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	params := mux.Vars(r)
	updateTodo(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	params := mux.Vars(r)
	deleteTodo(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
