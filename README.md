# Test Gin Go

```go run migrate/migrate.go```
Migrate the database.

```go mod download```
To install all the packages.

```CompileDaemon -command="./test_gin_go"```
To run the api and auto refresh for dev.

```go run .```
To run the api.

```go build```
To create an executable of the api.

```md
# Server
PORT=3000

# Database
DB_USER=root
DB_PASS=
DB_ADDRESS=127.0.0.1
DB_PORT=3306
DB_NAME=test_gin_go

```