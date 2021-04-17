#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet\greetpb\greet.proto
#protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative blog\blogpb\blog.proto