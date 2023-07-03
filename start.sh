#!/bin/sh

set -e

# Run database migration
echo "Running database migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# Start the application
echo "Starting the application"
exec "$@"
