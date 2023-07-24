- go mod init "module_name"

# Create grpc related stuff
- protoc --go_out=. --go-grpc_out=. proto/greet.proto

# Install missing dependencies
- go mod tidy

# Run 

- cd server -> go run *.go
- cd client -> go run *.go