#!/bin/bash

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && cd .. && pwd )"
DB_PATH="$ROOT_DIR/database/main.sqlite"
SEED_PATH="$ROOT_DIR/database/seed.sql"

sqlite3 "$DB_PATH" < "$SEED_PATH"

echo "Database seeded"
