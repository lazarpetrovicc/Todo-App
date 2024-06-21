package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodos(w)
	case "POST":
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/"):]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		getTodoByID(w, objID)
	case "PUT":
		updateTodo(w, r, objID)
	case "DELETE":
		deleteTodoByID(w, objID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter) {
	mutex.Lock()
	defer mutex.Unlock()

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	var todos []Todo
	if err = cursor.All(context.TODO(), &todos); err != nil {
		http.Error(w, "Error decoding todos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	result, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(todo)
}

func getTodoByID(w http.ResponseWriter, id primitive.ObjectID) {
	mutex.Lock()
	defer mutex.Unlock()

	var todo Todo
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTodo.Title},
			{Key: "completed", Value: updatedTodo.Completed},
		}},
	})

	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	updatedTodo.ID = id
	json.NewEncoder(w).Encode(updatedTodo)
}

func deleteTodoByID(w http.ResponseWriter, id primitive.ObjectID) {
	mutex.Lock()
	defer mutex.Unlock()

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
