#!/bin/bash

# Generate .proto code
protoc --go_out=plugins=grpc:.\greet .\greet\greetpb\greet.proto
protoc --go_out=plugins=grpc:.\calculator .\calculator\calcpb\calc.proto