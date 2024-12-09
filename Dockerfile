FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/server /app/
EXPOSE 8080
ENTRYPOINT ["./server"]
