postgres:
	docker run -d --name my-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=url_redirector -p 5432:5432 postgres:latest

server:
	go run ./cmd/server/main.go

client:
	go run ./cmd/main.go

gen-auth:
	protoc -I=./pkg/auth/pb --go_out=./ --go-grpc_out=./ ./pkg/auth/pb/*.proto

gen-url:
	protoc -I=./pkg/url/pb --go_out=./ --go-grpc_out=./ ./pkg/url/pb/*.proto

start:
	docker start my-postgres

mock_storage:
	mockgen -destination=pkg/mocks/mock_storage.go --build_flags=--mod=mod -package=mocks name-counter-auth/pkg/db Storage

test:
	go test -v -cover ./...