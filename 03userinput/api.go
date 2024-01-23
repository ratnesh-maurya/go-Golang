package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type CVE struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	// Add other fields as needed
}

func init() {
	// Initialize MongoDB connection
	mongoURI := "mongodb+srv://ratnesh:ratnesh@cluster0.3ka0uom.mongodb.net/"
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
}

func getCVEs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Connect to the "cve_db" database and "cve_list" collection
	collection := client.Database("cve_db").Collection("cve_list")

	// Query MongoDB for CVEs (you may need to adjust the query based on your data model)
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var cves []CVE
	for cur.Next(context.Background()) {
		var cve CVE
		err := cur.Decode(&cve)
		if err != nil {
			log.Fatal(err)
		}
		cves = append(cves, cve)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Convert CVEs to JSON and send the response
	json.NewEncoder(w).Encode(cves)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cves", getCVEs).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
