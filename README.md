# URLShortener

URLShortener is a microservice built with Go. It uses gRPC for communication and Postgres as a db.

## How to run

1. `git clone github.com/bhankey/URLShortener`
2. `cd URLShortener`
3. `docker-compose up`
4. connect from any gRPC client (e.g evans `evans ./backend/api/URLShortener.proto -p 8080`)

## gRPC URLShortener server
Building the gRPC server is done with the Dockerfile in the backend folder. Proto file is in backend/api

### Configuration 
configuration of server parameters is carried out using the .yaml file. Examples of configuration file /backend/config.yml

### Usage
  -config-path string
    	path to configuration file (default "config/config.yml")
      
