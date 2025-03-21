default: out/example

clean:
	rm -rf out

test: *.go
	go test ./...

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -buildvcs=false -o out/example ./cmd/example
