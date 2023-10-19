# Check if "docker-compose" is available, and if not, use "docker compose."
ifeq (, $(shell which docker-compose))
    DOCKER_COMPOSE_CMD := docker compose up --build
else
    DOCKER_COMPOSE_CMD := docker-compose up --build
endif

run:
	$(DOCKER_COMPOSE_CMD)

# Check if "docker-compose" is available, and if not, use "docker compose."
ifeq (, $(shell which docker-compose))
    DOCKER_COMPOSE_STOP := docker compose stop
else
    DOCKER_COMPOSE_STOP := docker-compose stop
endif
stop:
	$(DOCKER_COMPOSE_STOP)

mod:
	go mod tidy && go mod vendor

mocks:
	# Run mockgen to generate mock interfaces
	mockgen -source=./api/handlers.go -destination=mocks/mock_handlers.go -package=mocks
	mockgen -source=./services/articles/articles_service.go -destination=mocks/mock_service.go -package=mocks
	mockgen -source=./api/routes.go -destination=mocks/mock_routes.go -package=mocks
	# Print a message indicating the process is complete
	echo "Mock interfaces generated successfully."


test:
	go test ./...

cover:
	go test ./... -cover