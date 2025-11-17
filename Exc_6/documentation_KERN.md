# Final commands
```bash
docker compose down -v
docker compose up --build -d
```

# Final Links

http://localhost:8085/

http://orders.localhost:8085/openapi/index.html

http://localhost:42143/browser/orders/order_1.md

# AIS Order System – Debugging & Fix Report

This document summarizes all problems encountered during setup of the AIS Order System, why they occurred, and how they were resolved.  
It serves as a **technical post-mortem** and **reference guide**.

---

# 1. Background

The project uses:

- **Traefik** (reverse proxy)
- **static-web-server** (frontend)
- **orderservice** (Go backend)
- **MinIO** (object storage)
- **PostgreSQL**
- **Docker Compose**
- **Hostnames** using `orders.localhost`

Originally the system was designed so that Traefik listened on port **80**, meaning:

- `http://orders.localhost/`
- `http://orders.localhost/api/...`
- `http://orders.localhost/openapi/...`

would all be routed correctly.

---

# 2. The Initial Problem: Two PostgreSQL Instances

I had:

- one local PostgreSQL server running on port **5432**
- one Docker PostgreSQL container also originally using **5432**

This created a port conflict.

### Fix
The dockerized PostgreSQL server was moved to port **5555**:

```yaml
ports:
  - "5555:5555"
environment:
  - PGPORT=5555

# 3. Windows port issues

My windows laptop already maps 3 other processes to the :80 port so that had some conflicts when displaying the websites. 
Therefore I changed the port in the yaml of traefik to 8085:80. Then I was able to view the websites.


