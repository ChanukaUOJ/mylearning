#### Instruction

#### Download protoc package (this gives the ability generate protobuf)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

#### Download grpc package (this gives the ability generate grpc)
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#### Generate protobuf files
protoc \
--go_out=./proto_gen --go_opt=paths=source_relative \
--go-grpc_out=./proto_gen --go-grpc_opt=paths=source_relative \
coffee_shop.proto
