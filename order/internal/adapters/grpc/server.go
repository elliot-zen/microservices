package grpc

import (
	"github.com/elliot-zen/microservices-proto/golang/order"
	"github.com/elliot-zen/microservices/order/internal/ports"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}
