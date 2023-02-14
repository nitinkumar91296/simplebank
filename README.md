# simplebank

Simple Bank is a service to perform banking operations using Golang, Postgres, REST, gRPC, Docker etc.

INSTALL Golang Migrate Ubuntu:
https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/

SQLc : sudo snap install sqlc
https://docs.sqlc.dev/en/latest/overview/install.html
https://github.com/kyleconroy/sqlc

1. docker pull postgres:12-alpine
2. docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
3. docker network create bank-network
4. docker network connect bank-network postgres12
5. docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank
6. 

## Create migration files command with o/p:

1. nitin@nitin:~/Public/go/src/simplebank$ migrate create -ext sql -dir db/migration -seq add_users
2. /home/nitin/Public/go/src/simplebank/db/migration/000002_add_users.up.sql
3. /home/nitin/Public/go/src/simplebank/db/migration/000002_add_users.down.sql


###### Make start.sh and wait-for.sh files executable:

chmod +x wait-for.sh
chmod +x start.sh
