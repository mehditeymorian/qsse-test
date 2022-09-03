FROM golang:alpine AS BUILDER

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY internal .
COPY cmd .
COPY main.go .

RUN go build -o /qsse-test

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /qsse-test .

EXPOSE 8080
EXPOSE 4242

ENTRYPOINT ["./qsse-test"]



