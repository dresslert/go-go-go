FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache git bash gcc libc-dev

RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]
