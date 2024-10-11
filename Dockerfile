FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Install PostgreSQL client tools
RUN apk add --no-cache postgresql-client

# Copy the initialization script
COPY init-db.sh /app/init-db.sh
RUN chmod +x /app/init-db.sh

# Use the initialization script as the entrypoint
ENTRYPOINT ["/app/init-db.sh"]
CMD ["go", "run", "cmd/server/main.go"]

