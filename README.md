# User MicroService built with golang.
# Technology
 - Communication Protocol - grpc
 - Database                 - Mongodb

# Generate protobuf
Run the following command to generate your protobuf files in golang.
Make sure you have protoc in stalled on your system and the golang plugin.
Incase you don't have the plugin install it using the command below:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Then run this command to generate your protoco buffer (protobuf)
```sh
	protoc  --go_out=. --go-grpc_out=. proto/*.proto
```
or 
```sh
	protoc --go_out=. --go_opt=module=github.com/nielvid/go-userservice-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/nielvid/go-userservice-grpc proto/greet.proto
```
