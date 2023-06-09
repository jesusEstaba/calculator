Arithmetic Calculator REST API
------
API Rest build on go, who provides services for a simple calculator functionality, where each operation
has a separate cost per request.

**Operations:**
- Addition: `a+b`
- Subtraction: `a-b`
- Multiplication: `a*b`
- Division: `a/b`
- Square root: `√a`
- Random string generation

## System Requirements
- Go (1.9 or superior)
- MongoDB
- Docker (just in case You want to use a dockerized MongoDB)

## Instructions
Install local tools:
```shell
go install github.com/swaggo/swag/cmd/swag@v1.8.9
go install github.com/golangci/golangci-lint/cmd/golangci-lint
go install github.com/go-courier/husky/cmd/husky
```

Install git hooks:
```shell
husky init
```
> thanks to a husky script, when We make a `commit`, this run the following commands:
> - run test `go test ./...`
> - swagger generation `swag init -g cmd/api/main.go -o docs/`
> - linter `golangci-lint run ./...`
> 
> this allows us to automate these processes

Copy environment file
```shell
cp .env .env.example
```

Run a MongoDB container
```shell
docker-compose up -d
```

Install dependencies
```shell
go get
```

Run seeders
```shell
go run cmd/seed/main.go
```

Run project
```shell
go run cmd/api/main.go
```

## Swagger Documentation

Open the following URL in your browser to watch the documentation:

| Environment | Url                                                     |
|-------------|---------------------------------------------------------|
| local       | http://localhost:8080/docs/index.html                   |
| live        | https://estaba-calculator.herokuapp.com/docs/index.html |

## Deploy

This process it's perform thanks to [Heroku](https://heroku.com) automatically when We push changes to `master` branch

> live version host https://estaba-calculator.herokuapp.com