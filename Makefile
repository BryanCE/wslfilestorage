build:
	@go build -buildvcs=false -o bin/fs

run: build
	@./bin/fs

test:
	@go test -v ./...