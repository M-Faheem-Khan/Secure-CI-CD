FROM golang:latest
WORKDIR /dnsrecon
COPY . .
RUN go build -o ./build/dnsrecon ./cmd/main.go; mv ./build/dnsrecon /usr/local/bin/dnsrecon;
ENTRYPOINT ["/opt/local/bin/dnsrecon"]