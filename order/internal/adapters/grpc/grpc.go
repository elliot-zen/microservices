// grpc.go is for defining the handlders;
package grpc

import (
	"context"

	"github.com/elliot-zen/microservices-proto/golang/order"
	"github.com/elliot-zen/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	result, err := a.api.Get(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}
	var orderItems []*order.OrderItem
	for _, item := range result.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	return &order.GetOrderResponse{UserId: result.CustomerID, OrderItems: orderItems}, nil
}
