APP_NAME=merch-shop
DOCKER_COMPOSE=docker-compose
DOCKER_COMPOSE_FILE=docker-compose.yml

build:
	@echo "Building the application..."
	go build -o $(APP_NAME) ./cmd/main.go

run: build
	@echo "Running the application with Docker..."
	$(DOCKER_COMPOSE) up -d
	./$(APP_NAME)

test:
	@echo "Running tests..."
	go test ./tests -v

docker-down:
	@echo "Stopping and removing containers..."
	$(DOCKER_COMPOSE) down

clean:
	@echo "Cleaning up the project..."
	rm -f $(APP_NAME)

restart: docker-down run

install:
	@echo "Installing dependencies..."
	go mod tidy


.PHONY: build run test docker-down clean restart install
