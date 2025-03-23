### Simple Go gRPC microservice

Follow run.sh to set up local env or documentation below

##### 1. Install Go

Go, also known as Golang, is the programming language of choice for building gRPC microservices. If you haven't already installed Go, follow these steps to get started:

Download Go: Visit the official Go website (https://golang.org/dl/) and download the Go installer for your operating system.
Install Go: Run the installer and follow the on-screen instructions to install Go on your system. Ensure that Go's bin directory is added to your system's PATH variable.
Verify Installation: Open a terminal and run the following command to verify that Go is correctly installed:
go version
This command should display the installed Go version.

##### 2. Install gRPC Tools

To work with gRPC, you'll need to install the gRPC tools, which include the Protocol Buffers (Protobuf) compiler and the Go plugin. Follow these steps to install the gRPC tools:

Install Protobuf Compiler: Depending on your operating system, download the Protobuf compiler from the official repository (https://github.com/protocolbuffers/protobuf/releases).
Install the Go Protobuf Plugin: Use Go's package manager, go get, to install the Go Protobuf plugin:
```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
```
Install gRPC Tools: Install the gRPC tools, including the protoc-gen-go-grpc plugin, which is used for generating gRPC code:
```shell
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

##### 3. Verify gRPC Tools Installation

To ensure that the gRPC tools are correctly installed, run the following command:
```shell
protoc-gen-go-grpc --version
```
This command should display the installed version of the gRPC tools.

##### 4. Compile Service Stubs

```shell
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    shipping.proto
```

##### 5. Install dependencies

```shell
go mod tidy
```

##### 6. Run server and client

```shell
./run_client_server.sh
```
