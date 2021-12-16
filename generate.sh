#!/bin/env/sh
protoc \
	-I ./vendor/github.com/googleapis \
	-I ./vendor/github.com/grpc-ecosystem/grpc-gateway/v2 \
	-I ./vendor/github.com/envoyproxy/protoc-gen-validate \
	-I ./proto \
	--go_out . --go_opt paths=source_relative \
    	--go-grpc_out . --go-grpc_opt paths=source_relative \
    	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
    	--openapiv2_out ./swagger \
	proto/*/**.proto
