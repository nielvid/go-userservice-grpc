
plugin1:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
plugin2:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

genpb:
	protoc  --go_out=. --go-grpc_out=. proto/*.proto

genpb2:
	protoc --go_out=. --go_opt=module=github.com/nielvid/go-userservice-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/nielvid/go-userservice-grpc proto/greet.proto

grpcurl:
	grpcurl -plaintext -d '{"name": "Paul", "message": "how are you"}' localhost:8090 chat.ChatService.Greet

grpcui:
	grpui -plaintext localhost:8090