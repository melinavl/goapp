.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -v ./... -bench=. -run=xxx -benchmem

.PHONY: clean
clean:
	go clean
	rm -f bin/*
