build: dependencies build-api build-docker-services

# Deploy
dependencies:
	go mod download

build-api:
	go build -o ./bin/bulletin-api api/main.go

build-docker-services:
	docker-compose build

kompose:
	kompose -f docker-compose.yml convert -o k8s/

# Dev
clean:
	rm -rf bin/*

run:
	docker-compose up -d