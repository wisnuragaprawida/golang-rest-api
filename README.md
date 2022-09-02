# Rest API

for test MNC Bank.

## Installation

just type go mod tidy int terminal

```bash
$ go mod tidy
```

## Usage
run and migrate
```bash
ENV=migration API_URL=localhost:8000 DB_HOST=localhost DB_PORT=5432 DB_NAME=db_customer DB_USER=root DB_PASSWORD=root go run .

$
```

and create a user manualy with postgesql CLI

and hit the api POST localhost:8000/user/login  with bodyQuery type JSON email and password user registered before
