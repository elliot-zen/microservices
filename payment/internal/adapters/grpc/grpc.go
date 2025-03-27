package grpc

import (
	"context"
	"fmt"

	"github.com/elliot-zen/microservices-proto/golang/payment"
	"github.com/elliot-zen/microservices/payment/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.UserId < 1 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "user_id",
			Description: "user id cannot be less than 1",
		})
	}
	if request.OrderId < 1 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "order_id",
			Description: "order id cannot be less than 1",
		})
	}
	if len(validationErrors) > 0 {
		stat := status.New(400, "invalid order request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, fmt.Errorf("failed to charge. err: %v", err)
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
