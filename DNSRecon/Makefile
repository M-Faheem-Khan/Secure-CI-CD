PROJECT_NAME := dnsrecon

build:
	go build -o ./build/$(PROJECT_NAME) ./cmd/main.go

clean:
	rm -rf ./build/
	mkdir build

docker:
	docker build -t $(PROJECT_NAME) .