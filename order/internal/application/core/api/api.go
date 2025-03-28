package api

import (
	"strings"

	"github.com/elliot-zen/microservices/order/internal/application/core/domain"
	"github.com/elliot-zen/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		var allErrors []string
    if status.Code(paymentErr) != codes.InvalidArgument {
      allErrors = append(allErrors, paymentErr.Error())
    }
		stat := status.Convert(paymentErr)
		for _, detail := range stat.Details() {
			switch errType := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range errType.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetail, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetail.Err()
	}
	return order, nil
}
