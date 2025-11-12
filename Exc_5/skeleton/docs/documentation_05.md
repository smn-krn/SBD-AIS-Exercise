# Exercise 5 fixing docker shenanigans

## Final accessible URLs:
| Service                      |URL|
|------------------------------|----|
| Traefik Dashboard            | http://localhost:8080/dashboard |
| Backend (OrderService)       | http://orders.localhost/openapi/index.html |
| Frontend (Static Web Server) |http://localhost/|

## Final commands
make sure to delete everthing beforehand
```bash
docker compose down --volumes  # stoping all containers and remove volumes
docker system prune -f # cleans up dangling images/layers
docker ps
```

making the containers
```bash
docker compose up
```

## Everything in-between...

### Database issues
In GoLand there were connection issues so the port was mapped to then mapped to **localhost:55432** and was able to view the tables again.

### Browser showing issues
Only traefik Dashboard worked, the url were not shown on browser.
Error:
```bash
simone@Simone-2024:/mnt/c/Public/ais/sbd3/Exc_5/skeleton$ docker exec -it order-postgres psql -U docker -d order psql (18.0 (Debian 18.0-1.pgdg13+3)) 
Type "help" for help. 
order=# \dt -- lists tables SELECT * FROM orders LIMIT 5; 
Did not find any tables named "--". 
\dt: extra argument "lists" ignored 
\dt: extra argument "tables" ignored
\dt: extra argument "SELECT" ignored 
\dt: extra argument "*" ignored 
\dt: extra argument "FROM" ignored 
\dt: extra argument "orders" ignored 
\dt: extra argument "LIMIT" ignored 
\dt: extra argument "5;" ignored order=# \q
```

it was available though because
```bash
simone@Simone-2024:/mnt/c/Public/ais/sbd3/Exc_5/skeleton$ ping orders.localhost 
PING orders.localhost (127.0.0.1) 56(84) bytes of data. 
64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.123 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=2 ttl=64 time=0.087 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=3 ttl=64 time=0.094 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=4 ttl=64 time=0.323 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=5 ttl=64 time=0.091 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=6 ttl=64 time=0.049 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=7 ttl=64 time=0.082 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=8 ttl=64 time=0.092 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=9 ttl=64 time=0.092 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=10 ttl=64 time=0.089 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=11 ttl=64 time=0.090 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=12 ttl=64 time=0.095 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=13 ttl=64 time=0.092 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=14 ttl=64 time=0.091 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=15 ttl=64 time=0.094 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=16 ttl=64 time=0.035 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=17 ttl=64 time=0.054 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=18 ttl=64 time=0.039 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=19 ttl=64 time=0.090 ms 
^X64 bytes from localhost (127.0.0.1): icmp_seq=20 ttl=64 time=0.082 ms 
docker compose up -d traefik orderservice sws64 bytes from localhost (127.0.0.1): icmp_seq=21 ttl=64 time=0.089 ms 
64 bytes from localhost (127.0.0.1): icmp_seq=22 ttl=64 time=0.057 ms 64 bytes from localhost (127.0.0.1): icmp_seq=23 ttl=64 time=0.088 ms 
^C 
--- orders.localhost ping statistics --- 23 packets transmitted, 23 received, 0% packet loss, time 23154ms 
rtt min/avg/max/mdev = 0.035/0.092/0.323/0.053 ms
```

#### Trying fix nr 1:
windows + R

search for 'drivers'

open the 'hosts' file in the 'etc' folder as administrator

add 
```bash
# Added for Traefik Exercise AIS BA 2025 SBD
127.0.0.1 localhost
127.0.0.1 orders.localhost
```

this did not quite fix it but was apparently also necessary.

#### Trying fix nr 2:
```bash
netsh interface portproxy add v4tov4 listenport=80 listenaddress=127.0.0.1 connectport=8081 connectaddress=127.0.0.1
netsh interface portproxy show all
```

#### Trying fix nr 3:
working around windows firewall
Temporarily allowed inbound and outbound traffic for `docker-desktop-backend.exe` and `com.docker.backend.exe`. Added a rule for this.

#### Worked?

It worked after all this; I am no longer sure what part of this was importantly necessary.