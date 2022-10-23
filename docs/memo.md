- Initialize the go module
```
$ go mod init learn-go-with-tests
```

- Install go packages
```
$ go install golang.org/x/tools/cmd/godoc@latest
$ go install github.com/rakyll/gotest@latest
```

- Run go test files
```
$ gotest ./...
$ go test ./...
```