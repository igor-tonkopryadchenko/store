module github.com/igor-tonkopryadchenko/store

replace github.com/igor-tonkopryadchenko/store => ./

go 1.17

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/googleapis/googleapis v0.0.0-20211216002249-82a542279650 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
