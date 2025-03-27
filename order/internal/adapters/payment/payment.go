package payment

import (
	"context"
	"log"
	"time"

	"github.com/elliot-zen/microservices-proto/golang/payment"
	"github.com/elliot-zen/microservices/order/internal/application/core/domain"
	"github.com/elliot-zen/microservices/order/internal/utils"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Adapter struct {
	payment payment.PaymentClient
}

func CircuitBreakerClientInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, cbErr := cb.Execute(func() (any, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}
			return nil, nil
		})
		return cbErr
	}
}

func NewAdapter(paymentSerivceURL string) (*Adapter, error) {
	logrus.Info("=> Initialize payment adapter...")
	var opts []grpc.DialOption
	logrus.Info("=> Load payment circuit breaker config....")
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "demo",
		MaxRequests: 3,               // Allowed number of requests for a half-open circuit;
		Timeout:     3 * time.Second, // Timeout for an open to haf-open transtion;
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// This function decide on if the circuit will be open;
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			// This function is Executed on each state change;
			logrus.Warnf("Circuit Breaker: %s, changed from %v to %v", name, from, to)
		},
	})
	logrus.Info("=> Load TLS config....")
	tlsCredentials, tlsCredentialsErr := utils.GetTlsCredentials()
	if tlsCredentialsErr != nil {
		log.Fatalf("failed to load TLS config; err: %v", tlsCredentialsErr)
	}
	// Use go-grpc-middleware to apply retry logic to gRPC calls;
	opts = append(opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(3),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)))
	opts = append(opts, grpc.WithUnaryInterceptor(CircuitBreakerClientInterceptor(cb)))
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.NewClient(paymentSerivceURL, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
