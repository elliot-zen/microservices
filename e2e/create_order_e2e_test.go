package e2e

import (
	"context"
	"log"
	"testing"

	"github.com/elliot-zen/microservices-proto/golang/order"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateOrderTestSuited struct {
	suite.Suite
	stack *compose.DockerCompose
}

func (c *CreateOrderTestSuited) SetupSuite() {
	ctx := context.Background()
	composeFile := compose.WithStackFiles("resources/compose.yaml")
	stack, err := compose.NewDockerComposeWith(composeFile)
	if err != nil {
		log.Fatalf("Could not read compose file; err: %v", err)
	}
	c.stack = stack
	execError := stack.WithEnv(map[string]string{}).Up(ctx, compose.Wait(true))
	if execError != nil {
		log.Fatalf("Could not run compose stack; err: %v", err)
	}
}

func (c *CreateOrderTestSuited) TearDownSuite() {
	// err := c.stack.Down(context.Background(), compose.RemoveOrphans(true), compose.RemoveVolumes(true), compose.RemoveImagesLocal)
	// if err != nil {
	// 	log.Fatalf("could not shutdown compose stack: %v", err)
	// }
}

func (c *CreateOrderTestSuited) Test_Should_Create_Order() {
	var opts []grpc.DialOption
	ctx := context.Background()
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("failed to connect order service. Err: %v", err)
	}
	defer conn.Close()

	orderClient := order.NewOrderClient(conn)
	createOrderResponse, errCreate := orderClient.Create(ctx, &order.CreateOrderRequest{
		UserId: 23,
		OrderItems: []*order.OrderItem{
			{
				ProductCode: "CAM123",
				Quantity:    3,
				UnitPrice:   1.23,
			},
		},
	})
	c.Nil(errCreate)

	getOrderResponse, errGet := orderClient.Get(ctx, &order.GetOrderRequest{OrderId: createOrderResponse.OrderId})
	c.Nil(errGet)
	c.Equal(int64(23), getOrderResponse.UserId)
	orderItem := getOrderResponse.OrderItems[0]
	c.Equal(float32(1.23), orderItem.UnitPrice)
	c.Equal(int32(3), orderItem.Quantity)
	c.Equal("CAM123", orderItem.ProductCode)
}

func TestCreateORderTestSuite(t *testing.T) {
	suite.Run(t, new(CreateOrderTestSuited))
}
