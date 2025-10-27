#!/bin/sh

# todo
# docker build
# docker run db
# docker run orderservice

#!/bin/sh
set -e #exit immediately on error

#building the GO app image
docker build -t orderservice .

#create network if not exists
docker network create aisnet 2>/dev/null || true

#stopping and removing existing containers
docker rm -f ordersystem-db orderservice 2>/dev/null || true

#starting the PostgreSQL database
docker volume create pgdata_ais
docker run -d \
  --name pg-ais \
  --network aisnet \
  --env-file debug.env \
  -v pgdata_ais:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18


#waiting a few seconds for DB to initialize
echo "Waiting 5 seconds for Postgres to start..."
sleep 5


#starting the GO backend
docker run -d --name orderservice \
  --network aisnet \
  --env-file debug.env \
  -p 3000:3000 \
  orderservice

#Final message
echo "All containers started! Visit http://localhost:3000"