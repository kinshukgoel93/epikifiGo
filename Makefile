include .env


up:
	@echo "Starting mongodb container..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping mongodb container..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/	

start :	
	./${BINARY}	

restart : build start
