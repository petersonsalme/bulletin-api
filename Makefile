# Deploy
dependencies:
	go mod download

build-api:
	go build -o ./bin/bulletin-api api/main.go
	# go build -tags $(_ENV) -o ./bin/api api/main.go

build-docker-services:
	docker-compose build

kompose:
	kompose -f docker-compose.yml convert -o k8s/

build: dependencies build-api build-docker-services kompose

# Dev
clean:
	rm -rf bin/*

run:
	go run api/main.go

# compile:
# 	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
# 	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go