# Zachary Perry
# compiles problem_1 and problem_2 into respective bin/ files
.DEFAULT_GOAL := build

fmt: 
	@go fmt ./...

lint: fmt
	@golint ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o bin/problem_1 cmd/problem_1/problem_1.go
	@go build -o bin/problem_2a cmd/problem_2a/problem_2a.go
	@go build -o bin/problem_2b cmd/problem_2b/problem_2b.go
	@go build -o bin/primes cmd/prime_generator/prime_generator.go

clean:
	@go clean
	rm bin/*

test:
	@go test ./...

.PHONY: fmt lint vet build clean
