// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-grpc-project-structure/internal/rpci"
	"golang-grpc-project-structure/internal/services"
	"golang-grpc-project-structure/pkg/config"
	"golang-grpc-project-structure/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Injectors from app.go:

func InitApp() (*App, error) {
	configuration, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	listener, err := NewListener(configuration)
	if err != nil {
		return nil, err
	}
	server := NewGRPCServer()
	db := database.ConnectSQL(configuration)
	client := database.ConnectMongoDB(configuration)
	connected := Connected{
		sqlDB:   db,
		mongoDB: client,
	}
	greeterService := services.NewGreeterService()
	rpcIGreeterServe := rpci.NewGreeterServer(greeterService, server)
	rpcImpl := RpcImpl{
		RpcIGreeterServer: rpcIGreeterServe,
	}
	app := &App{
		listener:  listener,
		gsrv:      server,
		cfg:       configuration,
		connected: connected,
		rpcImpl:   rpcImpl,
	}
	return app, nil
}

// app.go:

// App contains minimal list of dependencies to be able to start an application.
type App struct {
	// listener is a TCP listener which is used by gRPC server.
	listener net.Listener
	// gRPC serer itself.
	gsrv *grpc.Server
	// gRPC server implementation. It's not used here directly, but it must be
	// initialized for registering. gRPC server.
	cfg       *config.Configuration
	connected Connected
	rpcImpl   RpcImpl
}

type Connected struct {
	sqlDB   *sql.DB
	mongoDB *mongo.Client
}

type RpcImpl struct {
	RpcIGreeterServer rpci.RpcIGreeterServe
}

func NewListener(config2 *config.Configuration,
) (net.Listener, error) {
	port := flag.String("port", config2.Server.PORT, "The server port")
	host := config2.Server.HOST
	address := fmt.Sprintf("%s:%s", host, *port)
	return net.Listen("tcp", address)
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func (a App) Start() error {
	return a.gsrv.Serve(a.listener)
}

func (a App) Addr() net.Addr {
	return a.listener.Addr()
}

func (a App) Config() *config.Configuration {
	return a.cfg
}

func (a App) App() App {
	return a
}

func (a App) Name() string {
	return "GRPS Server"
}

func (a App) Disconnect() {
	if err := a.connected.mongoDB.Disconnect(context.TODO()); err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB successfully")
	}

	if err := a.connected.sqlDB.Close(); err != nil {
		log.Printf("Failed to close SQL database connection: %v", err)
	} else {
		log.Println("Closed SQL database connection successfully")
	}
}