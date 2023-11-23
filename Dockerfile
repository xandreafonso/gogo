FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go
CMD ["./api"]

FROM scratch
COPY --from=builder /app/api /
CMD ["./api"]