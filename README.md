# Grpc Golang Course

This is a companion repository for my [GRPC Golang course](http://bit.ly/grpc-golang-github)

# Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB

# Installation
https://grpc.io/docs/languages/go/quickstart/
1. create a protobuff file, and go.mod file for dependency management
2. `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet\greetpb\greet.proto`