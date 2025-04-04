package ports

import "github.com/elliot-zen/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
