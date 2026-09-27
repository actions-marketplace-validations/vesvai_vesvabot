vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o bin/vesvabot cmd/vesvabot/main.go

run:
	./bin/vesvabot

clean:
	rm -rf bin