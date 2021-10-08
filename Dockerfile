FROM golang:1.17-alpine
WORKDIR app
COPY . .
RUN go build ./cmd/main/main.go
CMD ["./main"]