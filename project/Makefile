SHELL=cmd.exe
FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

start:
	@echo "Starting front end"
	cd ../front-end && go run ./cmd/web &
	@echo "Started front end!"

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Stopped front end!"