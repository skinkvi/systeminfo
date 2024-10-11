#!/bin/sh

# Wait for the database to be ready
until pg_isready -h db -p 5432 -U postgres; do
  echo "Waiting for the database to be ready..."
  sleep 2
done

# Set the password for the postgres user
export PGPASSWORD=postgres

# Check if the database exists and create it if it doesn't
psql -h db -p 5432 -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'system_infodb'" | grep -q 1 || psql -h db -p 5432 -U postgres -c "CREATE DATABASE system_infodb"

# Run migrations
goose -dir=./migrations postgres "postgres://postgres:postgres@db:5432/system_infodb?sslmode=disable" up

# Start the server
exec "$@"

