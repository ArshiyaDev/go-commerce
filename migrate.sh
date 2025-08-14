#!/bin/bash
set -e

# Load environment variables from .env
if [ ! -f .env ]; then
  echo ".env file not found!"
  exit 1
fi

export $(grep -v '^#' .env | xargs)

# Build database URL
DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

echo "Running migrations on database: $DB_NAME at $DB_HOST:$DB_PORT"

# Check if migrate CLI exists
if ! command -v migrate &> /dev/null
then
    echo "migrate CLI could not be found. Install it with:"
    echo "go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Run all up migrations
migrate -path ./migrations -database "$DB_URL" up || echo "No new migrations to run"

echo "Migrations completed successfully!"
