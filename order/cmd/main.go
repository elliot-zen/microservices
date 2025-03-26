package main

import (
	"log"

	"github.com/elliot-zen/microservices/order/config"
	"github.com/elliot-zen/microservices/order/internal/adapters/db"
	"github.com/elliot-zen/microservices/order/internal/adapters/grpc"
	"github.com/elliot-zen/microservices/order/internal/adapters/payment"
	"github.com/elliot-zen/microservices/order/internal/application/core/api"
)

// As always, main.go is the place where we handle dependency injections;
func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatasouceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database, error: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to initializee payment stub. error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
