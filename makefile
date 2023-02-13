proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/auth/pb/auth.proto pkg/user/pb/user.proto pkg/mail/pb/mail.proto

server:
	go run cmd/main.go