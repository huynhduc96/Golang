# Golang

# Init Docker

```bash
docker-compose up -d
```

make database reset

```bash
 sh database/db_reset.sh
```

make database dump

```bash
 sh database/db_dump.sh
```

destroy

```bash
docker-compose down -v
```

# Run app

```bash

go run Assignment/http-client/main.go

```

## check port 8080 exists

```bash
lsof -i :8080
```
