FROM golang:latest
WORKDIR /dnsrecon
COPY . .
RUN go build -o ./build/dnsrecon ./cmd/main.go
ENTRYPOINT ["/dnsrecon/build/dnsrecon"]