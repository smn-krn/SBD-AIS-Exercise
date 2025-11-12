#!/bin/sh
# wait-for-postgres.sh
set -e

host="$DB_HOST" # host from env variables
port="$POSTGRES_TCP_PORT" # port from env variables

echo "Waiting for Postgres at $host:$port..." # status message

until nc -z "$host" "$port"; do # checks whether the port is reachable
  echo "Postgres is unavailable - sleeping" 
  sleep 1  # repitition until response is recieved
done

echo "Postgres is up - executing ordersystem"
exec /app/ordersystem
