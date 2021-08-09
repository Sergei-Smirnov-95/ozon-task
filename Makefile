.PHONY: build,run,test

build:
	GOOS=linux go build -o ./bin/main ./

run:
	go run ./

test:
	go test ./ -cover