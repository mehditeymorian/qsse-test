FROM golang:1.18.5 AS BUILDER

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY internal ./internal
COPY cmd ./cmd
COPY main.go .

RUN go mod download


RUN CGO_ENABLED=0 go build -o /qsse-test main.go

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /qsse-test .

EXPOSE 8080
EXPOSE 4242

ENTRYPOINT ["./qsse-test"]



