
# Docker commands

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

# Main

run: 
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test: 
	go test -cover ./...
