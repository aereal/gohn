build: parser/parser.go
	go build ./...

parser/parser.go: parser/parser.go.y
	goyacc -o parser/parser.go parser/parser.go.y

test: parser/parser.go
	go test -v ./...
