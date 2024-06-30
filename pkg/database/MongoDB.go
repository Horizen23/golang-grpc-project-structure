package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"golang-grpc-project-structure/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func openClient(config *config.Configuration) (*mongo.Client, error) {
	clientOpts := options.Client().SetHosts(
		[]string{config.MongoDB.IP + ":" + strconv.Itoa(config.MongoDB.PORT)},
	).SetAuth(
		options.Credential{
			AuthSource:    "admin",
			AuthMechanism: "SCRAM-SHA-256",
			Username:      config.MongoDB.USERNAME,
			Password:      config.MongoDB.PASSWORD,
		},
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.MongoDB.CONN_TIME_OUT)*time.Second)
	defer cancel()
	MongoConnection, err := mongo.Connect(ctx, clientOpts)
	return MongoConnection, err
}
func ConnectMongoDB(config *config.Configuration) *mongo.Client {
	mongoClient, err := openClient(config)
	if err != nil {
		fmt.Println("Error of mongo connection: ", err.Error())
		panic(err)
	}
	fmt.Printf("MONGO         connected : %s:%d @user: %s \n", config.MongoDB.IP, config.MongoDB.PORT, config.MongoDB.USERNAME)

	return mongoClient
}
