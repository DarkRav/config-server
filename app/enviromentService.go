package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Application struct {
	Name       string            `json:"name"`
	Profile    string            `json:"profile"`
	Properties map[string]string `json:"properties"`
}

type Environment struct {
	Application Application `json:"application"`
}

type EnvironmentService struct {
	mongoClient *mongo.Client
}

func createEnvironmentService() *EnvironmentService {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	environmentService := new(EnvironmentService)
	environmentService.mongoClient = client
	return environmentService
}

func (environmentService EnvironmentService) getConfigs(application string, profile string) (string, error) {
	collection := environmentService.mongoClient.Database("config-server").Collection("environments")

	query := bson.M{
		"application.name":    application,
		"application.profile": profile,
	}
	var environment Environment
	err := collection.FindOne(context.TODO(), query).Decode(&environment)
	if err != nil {
		return "", err
	}
	environmentJson, _ := json.Marshal(environment)
	log.Print(string(environmentJson))
	return string(environmentJson), nil
}
