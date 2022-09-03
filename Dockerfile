FROM golang:alpine AS BUILDER

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY internal .
COPY cmd .
COPY main.go .

RUN go build main.go -o /main

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /main .

EXPOSE 8080
EXPOSE 4242

ENTRYPOINT ["./main"]



