# Makefile for development;
ENV = development
DATA_SOURCE_URL = root:root@tcp(127.0.0.1:3306)/$(mod)
APPLICATION_PORT = $(port)
PAYMENT_SERVICE_URL="localhost:3001";

run:
	cd $(mod) && \
		go mod tidy && \
		go run cmd/main.go
