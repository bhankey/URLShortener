# URLShortener

URLShortener microservice built with Go. Use gRPC for communication and Postgres as a db.

## How to run

1. `git clone github.com/bhankey/URLShortener`
2. `cd URLShortener`
3. `docker-compose up`
4. connect from any gRPC clien (e.g evans `evans backend/api/URLShortener.proto -p 8080`)
