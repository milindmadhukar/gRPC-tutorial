watch:
	@go run github.com/cosmtrek/air@latest --build.cmd "go build -o bin/api main.go" --build.bin "./bin/api"

proto:
	rm -f pb/*.go
	rm docs/swagger/api.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
			--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
			--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=api \
			proto/*.proto

evans:
	evans --host 127.0.0.1 --port 9090 -r repl && package pb && service API

.PHONY: watch proto evans
