golang-managedb

A example code to use ent/ent module for Go.

## Installation

```Shell
go mod tidy
```

## Usage
First, create and start database container.

Testing connect to database.
```Shell
go run ./start/connect-db/main.go
```

## Using Docker

This example uses Docker for testing because it uses a database.

Create and start containers:
```shell
cd ./docker/golang-managedb
docker-compose up -d
```

Execute an interactive bash shell on the container:
```shell
docker exec -it golang-managedb-mariadb bash
```

View docker log
```shell
docker logs golang-managedb-mariadb
```

### Browsing database 
Visit `http://localhost:8080/` and browse database with "adminer".

The password is set in compose.yml.

This is the easiest way to verify that the database is up and running.