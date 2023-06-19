# API Cotação

## About the project

- An api for quote.
App use Golang, Postgres

## How to run the api?

**Step 1: Clone the project, run the following commands:**

**Step 2: Create file config.json**
{
  "database": {
    "url": "://postgres:postgres@172.21.0.2:5432/postgres"
  },
  "server": {
    "port": "8081"
  },
  "secret_key": "HsZwxbcwuK/yZBbNeUjTAA=="
}

To get the database ip, use...
docker ps -> check container id postgres, then...
docker inspect container_id
Exemple; "IPAddress": "172.21.0.2"

**Step 3: Run docker(at the root of the project)**
- docker-composer up -d

**Step 4: go mod tidy(at the root of the project)**
- go mod tidy

**Step 6: run api**
- go run adapter/http/main.go

**Step 7: run consumer**
- go run adapter/cmd/main.go

# Link Doc(Swagger)
http://localhost:3001/swagger/index.html
To generate docs, run; swag init -d adapter/http --parseDependency --parseInternal --parseDepth 2 -o adapter/http/docs
