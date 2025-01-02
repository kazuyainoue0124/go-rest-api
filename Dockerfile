##
## build
##
FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

##
## test
##
FROM build-stage AS run-test-stage
RUN go test -v ./...

##
## deploy
##
FROM gcr.io/distroless/base-debian12 AS release-stage

WORKDIR /app

COPY --from=build-stage /app/server /app/server

EXPOSE 8080

USER nonroot:nonroot

CMD ["./server"]
