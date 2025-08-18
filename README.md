## Example AI service for SNET Decentralized AI Platform

* Golang 1.24.6
* Simple calls
* Training support
* Service type grpc
* Default port 5001
* Minimum requirements

## Proto generate or update

1. ```cd service```

2. ```protoc -I . *.proto --go-grpc_out=. --go_out=.```

## Run with go

```go run main.go```

