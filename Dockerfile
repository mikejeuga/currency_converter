FROM golang:1.19

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
copy ./ ./

RUN go build -o main ./cmd/main.go

EXPOSE 8002

CMD ["./main"]