# vinyl-store

## Tools:
 - Swagger Editor: [`editor.swagger.io/`](https://editor.swagger.io/) 
 - Postman: [`www.postman.com/downloads/`](https://www.postman.com/downloads/)

## API Documentation
Get at: [`openapi.yaml`](api/openapi.yaml) 

Preview via copying into: [`editor.swagger.io/`](https://editor.swagger.io/) 

## Requirements
 - Go: `brew install go`
 - Entgo: `go install entgo.io/ent/cmd/ent@latest`
 - PostgreSQL: [`www.postgresql.org/`](https://www.postgresql.org/)

## Prepare Database
1. Create database via PostgreSQL named `vinyl-store`
2. Create tables using queries from [`create-table.sql`](database/create-table.sql)
3. Insert data in tables using [`insert-into.sql`](database/insert-into.sql)

## Clone project
```
git clone https://github.com/sashawins/vinyl-store.git
```
```
cd vinyl-store
```

## Create environment file
In project's folder run:
```
cp .env.example .env
```
```
vim .env
```
 — inside `.env` file replace `yourpassword` with your PostgreSQL password and `your_secret_key` with your JWT secret key.


## Install server dependencies
In project's folder run:
```
go mod tidy
```

## Run server
In project's folder run:
```
go run .
```

## Test API requests:
In Postman:
 - File -> Import... -> Choose [`openapi.yaml`](api/openapi.yaml)
