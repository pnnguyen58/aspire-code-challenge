
protoc:
	protoc -I . \
		--go_out=./pkg \
		--go_opt=module=github.com/pnnguyen58/aspire-code-challenge \
		--go-grpc_out=./pkg \
		--go-grpc_opt=module=github.com/pnnguyen58/aspire-code-challenge \
		--grpc-gateway_out=./pkg \
		--grpc-gateway_opt module=github.com/pnnguyen58/aspire-code-challenge \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt generate_unbound_methods=true \
		./pkg/proto/*.proto