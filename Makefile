.PHONY: build clean docker down run stop test

build:
	go build -o server ./cmd/server

clean:
	rm -f server

docker:
	docker build -t server .

down:
	docker compose down

run:
	docker compose up -d

stop:
	docker compose down

test:
	go test ./...