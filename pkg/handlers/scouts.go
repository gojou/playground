package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gojou/playground/pkg/svc/scout"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Scouts displays full list of Scouts
func Scouts(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	scouts, err := getScouts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message: "` + err.Error() + `"}"`))
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(scouts)

}

func getClient() mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return *client
}

func getScouts() ([]scout.Scout, error) {
	var scouts []scout.Scout
	client := getClient()
	collection := client.Database("goscouting").Collection("scouts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(`{ "message: "` + err.Error() + `"}"`))
		return scouts, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var scout scout.Scout
		cursor.Decode(&scout)
		scouts = append(scouts, scout)
	}
	if err := cursor.Err(); err != nil {

		return scouts, err
	}
	return scouts, nil

}
