# simplebank
Simple Bank is a service to perform banking operations using Golang, Postgres, REST, gRPC, Docker etc.

INSTALL Golang Migrate Ubuntu:
https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/

1. docker pull postgres:12-alpine
2. docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
