# simplebank
Simple Bank is a service to perform banking operations using Golang, Postgres, REST, gRPC, Docker etc.

INSTALL Golang Migrate Ubuntu:
https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/

SQLc : sudo snap install sqlc
https://docs.sqlc.dev/en/latest/overview/install.html
https://github.com/kyleconroy/sqlc

1. docker pull postgres:12-alpine
2. docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


## Create migration files command with o/p:
nitin@nitin:~/Public/go/src/simplebank$ migrate create -ext sql -dir db/migration -seq add_users
/home/nitin/Public/go/src/simplebank/db/migration/000002_add_users.up.sql
/home/nitin/Public/go/src/simplebank/db/migration/000002_add_users.down.sql



