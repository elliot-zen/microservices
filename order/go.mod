module github.com/elliot-zen/microservices/order

go 1.23.7

require (
	github.com/elliot-zen/microservices-proto/golang/order v0.0.4
	github.com/elliot-zen/microservices-proto/golang/payment v0.0.4
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.1
	github.com/sirupsen/logrus v1.9.3
	github.com/sony/gobreaker v1.0.0
	github.com/stretchr/testify v1.10.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463
	google.golang.org/grpc v1.71.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
