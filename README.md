# mocking http outbound requests in go

This repo demonstrates a few ways to mock outbound HTTP requests in Go.

## Installation

```
git clone https://github.com/sonya/go-http-mocking
cd go-http-mocking
cd approaches
go mod download`
```

## Running the unit tests

From the root of this repo, run:
```
cd approaches
go test ./...
```

## Running the integration tests

From the root of this repo, run:
```
docker-compose up -d
cd approaches
go test --tags=integration ./... # one of these should fail
docker-compose down
```
