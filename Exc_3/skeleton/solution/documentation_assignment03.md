# Solution steps for assignment 3

## Step 1
open wsl and enter
```bash
 cd /mnt/c/Public/ais/sbd3/SBD-AIS-Exercise/Exc_3/skeleton
```
this is the path I'm on, and it is different from my windows path because linux require it differently
all commands unless specified otherwise were run here

## Step 2
create a persistent volume for postgres data
```bash
docker volume create pgdata_sbd_ais
```

## Step 3
running postgres 18 (detached), using our env file, host network, and the storage path from the readme
```bash
docker run -d --name pg-ais \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  --network host \
  -v pgdata_ais:/var/lib/postgresql/18/docker \
  postgres:18
```

(sanity checks:
````bash
docker ps
docker logs -f pg-ais 
````

## Step 4
make Dockerfile

## Step 5
change the appropriate part in db.go (func prepopulate)

## Step 6
Build image
```bash
docker build -t orderservice .
```
this gave me an error at first and in the end I fixed it by executing
```bash
 sudo apt-get update && sudo apt-get install -y dos2unix
 
 dos2unix scripts/*.sh
```

after this, it worked and the container was built appropriately.

## Step 7
restart DB and app
```bash
docker rm -f pg-ais orderservice 2>/dev/null || true

docker run -d --name pg-ais \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  --network host \
  -v pgdata_ais:/var/lib/postgresql/18/docker \
  postgres:18

docker run -d --name orderservice \
  --env-file debug.env \
  --network host \
  orderservice

# check logs
docker logs -f orderservice
```

## Step 8
Testing via
```bash
curl -s http://localhost:3000/api/menu | jq
```
now at this point it turned out that windows doesn't really like supporting the --network host so a couple of changes were in order

## Step 9
clean up briefly
```bash
docker rm -f orderservice pg-ais 2>/dev/null || true
docker network rm aisnet 2>/dev/null || true
```

## Step 10
create a network
```bash
docker network create aisnet
```

## Step 11
Prepare env for containers
the debug.env was changed to
POSTGRES_DB=order
POSTGRES_USER=docker
POSTGRES_PASSWORD=docker
POSTGRES_TCP_PORT=5432
DB_HOST=pg-ais

[change at DB_HOST]

**recreated the container** at this step so it contains the changed file

## Step 12
Finally cleaning everything and making sure it runs

```bash

# make sure the network exists
docker network create aisnet 2>/dev/null || true

# stop any old containers
docker rm -f orderservice pg-ais 2>/dev/null || true

# start Postgres (on the same network)
docker run -d --name pg-ais \
  --network aisnet \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  -v pgdata_ais:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18

# start your app and JOIN THE SAME NETWORK
docker run -d --name orderservice \
  --network aisnet \
  --env-file debug.env \
  -p 3000:3000 \
  orderservice

```

## Step 13

tried to connect it to the database on GoLand
somehow the user wasn't created so I did
```bash
docker exec -e PGPASSWORD=docker -it pg-ais \
  psql -U docker -d order -c "\du"
```

to initialize the user

## Step 14
run
```bash
./scripts/run.sh
```



