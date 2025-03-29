package ports

import (
	"context"

	"github.com/elliot-zen/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	Get(ctx context.Context, id int64) (domain.Order, error)
}
