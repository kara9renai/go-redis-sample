FROM golang:1.19 as dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "*.go"]