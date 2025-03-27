package utils

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

// GetTlsCredentials: This function is used to send a request to payment service;
func GetTlsCredentials() (credentials.TransportCredentials, error) {
	clientCert, clientCertErr := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if clientCertErr != nil {
		return nil, fmt.Errorf("failed to load TLS cerdentials, err: %w", clientCertErr)
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
    ServerName: "localhost",
		Certificates: []tls.Certificate{clientCert},
		ClientCAs:    certPool,
	}), nil

}
