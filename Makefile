build: parser.go
	go build ./...

parser.go: parser.go.y
	goyacc -o parser.go parser.go.y

test: parser.go
	go test -v ./...
