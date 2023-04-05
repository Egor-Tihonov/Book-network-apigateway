proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/auth-service/pb/auth.proto pkg/user-service/pb/user.proto pkg/mail-service/pb/mail.proto pkg/book-service/pb/book.proto

server:
	go run cmd/main.go