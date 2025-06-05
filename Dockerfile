FROM golang:1.23

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

RUN go build -o walletcore cmd/walletcore/main.go

CMD ["./walletcore"]

# CMD ["tail", "-f", "/dev/null"]

