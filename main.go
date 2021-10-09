package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.org/crypto/bcrypt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Connection mongoDB with helper class
var collection = connection.ConnectDB()
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUserId).Methods("GET")
	r.HandleFunc("/api/posts", createPost).Methods("POST")
	r.HandleFunc("/api/post/{id}", getPost).Methods("GET")
	r.HandleFunc("/api/users/{id}", getAllPosts).Methods("GET")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our user model.
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		connection.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getUserId(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		connection.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&post)

	// insert our post model.
	result, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		connection.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		connection.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(post)
}



func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Post array
	var post []models.Post

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var post models.Post
		// & character returns the memory address of the following variable.
		err := cur.Decode(&post) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(posts) // encode similar to serialize process.
}