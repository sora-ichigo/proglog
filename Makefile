all: tools gen

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.22.0

gen: tools
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.
test:
	go test -v -race ./...
