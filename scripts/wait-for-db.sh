#!/bin/sh

set -e

DB_HOST="${1}"
DB_PORT="${2}"
TIMEOUT="${3}"
shift 3

echo "Waiting for Postgres at $DB_HOST:$DB_PORT (timeout ${TIMEOUT}s)..."

start_ts=$(date +%s)
while ! nc -z "$DB_HOST" "$DB_PORT"; do
    sleep 1
    now_ts=$(date +%s)
    if [ $((now_ts - start_ts)) -ge "$TIMEOUT" ]; then
        echo "Timeout reached waiting for Postgres at $DB_HOST:$DB_PORT"
        exit 1
    fi
done

echo "Postgres is ready! Running command: $*"
exec "$@"