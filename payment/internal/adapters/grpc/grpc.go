package grpc

import (
	"context"
	"fmt"

	"github.com/elliot-zen/microservices-proto/golang/payment"
	"github.com/elliot-zen/microservices/payment/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, fmt.Errorf("failed to charge. err: %v", err)
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
