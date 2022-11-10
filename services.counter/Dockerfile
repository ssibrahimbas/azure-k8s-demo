FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy
COPY . .
RUN go build -o main ./src/main.go

CMD ["./main"]