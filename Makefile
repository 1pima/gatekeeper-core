# Переменные
APP_NAME := gatekeeper
MAIN_PATH := ./cmd/$(APP_NAME)/main.go
PROTO_SRC_DIR := ./proto-contracts/proto
GO_GEN_DIR := ./pkg/proto/$(APP_NAME)
PYTHON_GEN_DIR := ./pkg/proto/python

.PHONY: all generate-contracts clean-contracts tidy build run

# Основные команды
all: tidy generate-contracts build

# Работа с зависимостями
tidy:
	@echo "Running go mod tidy..."
	go mod tidy
	go mod vendor

# Генерация контрактов
generate-contracts:
	@echo "Generating gRPC contracts..."
	mkdir -p $(GO_GEN_DIR)
	# Go генерация
	protoc -I=$(PROTO_SRC_DIR) \
		--go_out=$(GO_GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(shell find $(PROTO_SRC_DIR) -name "*.proto")
	@echo "Go generation OK!"

# Билд приложения
build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

# Запуск
run:
	go run $(MAIN_PATH)

# Очистка
clean-contracts:
	@echo "Cleaning generated files..."
	rm -rf $(GO_GEN_DIR)
	rm -rf $(PYTHON_GEN_DIR)