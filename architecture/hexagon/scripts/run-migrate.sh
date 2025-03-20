#!/bin/bash

CURRENT_DIR=$(pwd)

# Read env in .env file
export $(cat .env | xargs)

cd internal/adapters/framework/right/database/sql/schema

DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

echo "Running migrate up on $DB_URL"

goose postgres "$DB_URL" up

cd $CURRENT_DIR
