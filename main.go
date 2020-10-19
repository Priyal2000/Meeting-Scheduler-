package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("The API is running")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/meetings", MeetingScheduler)
	http.HandleFunc("/articles/", Participants)
	http.HandleFunc("/meeting/", GetMeeting)
	http.ListenAndServe(":12345", nil)
}
