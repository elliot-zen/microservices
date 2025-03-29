package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/elliot-zen/microservices/order/internal/application/core/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DataSourceUrl string
}

func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()

	mysqlContainer, err := mysql.Run(ctx,
		"mysql:8.0.36",
		mysql.WithDatabase("orders"),
		mysql.WithUsername("root"),
		mysql.WithPassword("s3cr3t"),
	)
	if err != nil {
		logrus.Fatalf("failed to start container: %s", err)
		return
	}
	endpoint, _ := mysqlContainer.Endpoint(ctx, "")
	o.DataSourceUrl = fmt.Sprintf("root:s3cr3t@tcp(%s)/orders?charset=utf8mb4&parseTime=True&loc=Local", endpoint)
	logrus.Infof("=> Get a test connect: %s", o.DataSourceUrl)
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := NewAdapter(o.DataSourceUrl)
	o.Nil(err)
	saveErr := adapter.Save(context.Background(), &domain.Order{})
	o.Nil(saveErr)
}

func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
  ctx := context.Background()
	adapter, _ := NewAdapter(o.DataSourceUrl)
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "CAM",
			Quantity:    5,
			UnitPrice:   1.32,
		},
	})
	adapter.Save(ctx, &order)
	ord, _ := adapter.Get(ctx, order.ID)
	o.Equal(int64(2), ord.CustomerID)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
