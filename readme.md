# golang-managedb
[![](https://img.shields.io/github/go-mod/go-version/golang/go/release-branch.go1.18?filename=src%2Fgo.mod&label=GO%20VERSION&style=for-the-badge&logo=appveyor)](https://github.com/golang/go/releases/tag/go1.18)

A example code to use [ent/ent](https://github.com/ent/ent) entity framework for Go.

## Installation

Create go.sum
```Shell
go mod tidy
```

## Usage
First, create and start database container.

### 1. Create "testdb" database.

### 2. Testing connect to database
```Shell
go run ./example/connect_db/main.go
```

### 3. Testing migration tool

#### The structure of the table to create
- user_group
  - id : bigint auto increment (Primary Key)
  - name :  varchar(30) not null
  - created_at : datetime
- user
  - id : bigint auto increment (Primary Key)
  - first name : varchar(30) not null
  - last name : varchar(30) not null
  - mail : varchar(50) not null
  - usergroup_id : bigint not null
  - reated_at : datetime

```Shell
go generate ./ent
```

```Shell
go run ./example/exec_migration/main.go
```

### 4. Testing CURD

Create (Insert):
```Shell
go run ./example/curd/create/main.go
```

Update:
```Shell
go run ./example/curd/update/main.go
```

Read (Select):
```Shell
go run ./example/curd/read/main.go
```

Delete:
```Shell
go run ./example/curd/delete/main.go
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

## ToDo
- Clean up example codes
