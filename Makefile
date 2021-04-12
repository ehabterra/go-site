

start-db:
	./run.sh

run:
	go run ./cmd/site

build:
	go build -o ./bin/go-site ./cmd/site