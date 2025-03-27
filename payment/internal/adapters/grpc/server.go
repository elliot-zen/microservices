package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/elliot-zen/microservices-proto/golang/payment"
	"github.com/elliot-zen/microservices/payment/config"
	"github.com/elliot-zen/microservices/payment/internal/ports"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func getTlsCredentials() (credentials.TransportCredentials, error) {
	serverCert, serverCertErr := tls.LoadX509KeyPair("cert/server.crt", "cert/server.key")
	if serverCertErr != nil {
		return nil, fmt.Errorf("failed to load TLS cerdentials, err: %w", serverCertErr)
	}
	certPool := x509.NewCertPool()
	caCert, caCertErr := os.ReadFile("cert/ca.crt")
	if caCertErr != nil {
		return nil, fmt.Errorf("failed to load CA credentials, err: %w", caCertErr)
	}
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append the CA certs;")
	}
	return credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequestClientCert,
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
	}), nil

}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error; %v", a.port, err)
	}
  logrus.Infof("=> Loading TLS config...")
	tlsCrendentials, tlsCrendentialsErr := getTlsCredentials()
	if tlsCrendentialsErr != nil {
		log.Fatalf("failed to load TLS config; err: %v ", tlsCrendentialsErr)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(tlsCrendentials))
	grpcServer := grpc.NewServer(opts...)
	payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed serve grpc on port")
	}

}
