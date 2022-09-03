FROM golang:alpine AS BUILDER

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY internal .
COPY cmd .
COPY main.go .

RUN go mod download


RUN go build -o /qsse-test main.go

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /qsse-test .

EXPOSE 8080
EXPOSE 4242

ENTRYPOINT ["./qsse-test"]



