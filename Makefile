# Makefile for development;
ENV = "development"
DATA_SOURCE_URL = "root:root@tcp(127.0.0.1:3306)/$(mod)"
APPLICATION_PORT = $(port)
PAYMENT_SERVICE_URL= "localhost:3001"
DEV_ENVS=ENV=$(ENV) DATA_SOURCE_URL=$(DATA_SOURCE_URL) APPLICATION_PORT=$(APPLICATION_PORT) PAYMENT_SERVICE_URL=$(PAYMENT_SERVICE_URL)
run:
	@cd $(mod) && \
		go mod tidy && \
		$(DEV_ENVS) go run cmd/main.go
