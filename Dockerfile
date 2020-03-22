FROM golang:alpine as builder 

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o book-service cmd/main.go

# Running the book-service image container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

EXPOSE 8082

COPY --from=builder /app/book-service .
# COPY --from=builder /app/build/.env ./build/

CMD ["./book-service"]