build: parser.go
	go build ./...

parser.go: parser.go.y
	go tool yacc -o parser.go parser.go.y

test: parser.go
	go test -v ./...
