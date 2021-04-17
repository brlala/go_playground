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

# SSL Encryption
1. Step 1: Generate Certificate Authority + Trust Certificate(ca.crt)
   1. Generate Certificate Authority(ca.key) + Trust certificate(ca.crt)
   2. The trust certificate should be shared with all clients for them to verify the CA 
2. Step 2: Generate the Server Private Key (server.key)
3. Step 3: Get a certificate signing request from the CA (server.csr)
   1. Sign the server key to claim that the servername is what we defined, in example is `localhost`
   2. Server will output a server signing request(server.csr)
4. Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
   1. In real life we will send the server.csr to whoever is the CA
   2. They will sign it for a certain amount of days, using the files, server.csr, ca.crt and ca.key to output server.crt.
5. Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC

# Doing reflection with Evans CLI
1. `evans -r repl` - Start the EVANS CLI
2. `show package` - show all package available
3. `show service` - show all RPC method and request response type
4. `desc GreetRequest` - see the implementation and description of the GreetRequest
4. `package greet` - go into package Greet
5. `show service` - show all service 
6. `service GreetService` - go into GreetService
7. `call Greet` - call the method Greet