package main

import (
	"log"

	"github.com/elliot-zen/microservices/payment/config"
	"github.com/elliot-zen/microservices/payment/internal/adapters/db"
	"github.com/elliot-zen/microservices/payment/internal/adapters/grpc"
	"github.com/elliot-zen/microservices/payment/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatasouceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database, error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
