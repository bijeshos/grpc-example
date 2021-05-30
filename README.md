# gRPC Go example

A simple gRPC example in Go

# Prerequisites
- Protocol Buffer Installation
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


# Generate Proto files
- `$ protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative protof/message.proto`

# Run Server
- `$ go run main.go server`

# Run Client
- `$ go run main.go client`