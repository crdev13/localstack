# localstack-golang-s3

Golang example 

# Requirements
1. Docker
2. Docker-compose

# Steps
1. `$ docker-compose up`
2. `$ aws --endpoint-url=http://localhost:4566 --no-sign-request s3 mb s3://test-dev`
3. `$ go run main.go`