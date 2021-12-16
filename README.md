# GRPC server with http gateway

## Validation and openapi v2 schema generator included!

### Prepare
```
$ go mod vendor

$ sh generate.sh
```

### Run
```
$ go run main.go
```

### Test
```
$ http :8081/swagger/service/addons_v1.swagger.json

$ http :8081/v1/addons/sections 

$ http PUT :8081/v1/addons/section/okay          
```
