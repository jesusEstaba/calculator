Calculator
------

### Development
Install local tools
```shell
go install github.com/swaggo/swag/cmd/swag@v1.8.9
go install github.com/golangci/golangci-lint/cmd/golangci-lint
go install github.com/go-courier/husky/cmd/husky
```

Install git hooks
```shell
husky init
```
> thanks to husky pre-commit allow us to run commands like swagger generation files and the linter

Open the following URL in your browser to watch the documentation:

[http://localhost:8080/api/calculator/v1/docs/index.html](http://localhost:8080/api/calculator/v1/docs/index.html)