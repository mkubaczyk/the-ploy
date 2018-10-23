# ThePloy

## Development

### docker-compose

#### Dependencies

[Docker](https://www.docker.com)

[docker-compose](https://docs.docker.com/compose/install/)

#### Run

`docker-compose up -d --force-recreate --build`

It runs one API server and two workers for the queue.

API is available at `http://127.0.0.1`.

Test it with `http://127.0.0.1/api/rmq` or `http://127.0.0.1/api/deployments/1`

`docker-compose logs -f` for real-time logs.

### Locally

#### Dependencies

`brew install dep` for golang package manager

Place this project in `$GOPATH/src/github.com/mkubaczyk/` directory

Run `dep ensure` to install packages in `vendor` directory

It still requires docker-compose to be built and run because of Redis and MySQL services dependency. Or do it on your own.

#### Run

`go run main.go` runs API at `http://127.0.0.1:8080`

`go run consumer.go` runs one queue worker
