FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN apk add --no-cache postgresql-client

COPY init-db.sh /app/init-db.sh
RUN chmod +x /app/init-db.sh

ENTRYPOINT ["/app/init-db.sh"]
CMD ["go", "run", "cmd/server/main.go"]

