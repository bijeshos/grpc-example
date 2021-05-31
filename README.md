# gRPC Go example

A simple gRPC client-server example in Go

# Prerequisites
- Install Protocol Buffer
    - Protocol Buffer (version 3 - proto3) compiler
        - manually download pre-compiled binary from here (for your operating system): https://github.com/protocolbuffers/protobuf/releases (fe.g.: protoc-3.17.0-linux-x86_64.zip)
        - unzip the file
        - update path to include `protoc`
            - `$export PATH="$PATH:$(go env GOPATH)/bin"`
        - check installation by running
            - `protoc --version`
- Install Go plugins
    - $ go install google.golang.org/protobuf/cmd/protoc-gen-go 
    - $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Clone Repo
- Clone this repo
- Run the following command
    - `$ go get`

# Run Server
Run the following to start server
- `$ go run main.go server`

# Run Client
Run the following to start client
- `$ go run main.go client`

# Re-generate Proto files (Optional)
- If any changes are made in proto service definition files, regenerate the proto files
    - `$ protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative protof/product.proto`
- Make respective changes in server/client implementations
- Rerun server/client