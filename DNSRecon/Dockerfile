FROM golang:latest
WORKDIR /dnsrecon
COPY . /DNSRecon
RUN go build -o ./build/dnsrecon ./cmd/main.go; mv /dnsrecon/build/dnsrecon /usr/local/bin/dnsrecon;
ENTRYPOINT ["/opt/local/bin/dnsrecon"]