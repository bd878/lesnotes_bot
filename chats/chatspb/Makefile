all : api.proto
	protoc \
		--proto_path=./ \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=import \
		--go-grpc_opt=paths=import \
		--go_opt=module="github.com/bd878/lesnotes_bot/chats/chatspb" \
		--go-grpc_opt=module="github.com/bd878/lesnotes_bot/chats/chatspb" \
		api.proto