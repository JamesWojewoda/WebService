# Go Web Service

This Web Service is set up to post strings and their sha256 encrypted values into a map. The map's data is ephemeral in nature
and will last as long as the docker container runs. Also, the web service is set up to get string values based on sha256 values
received via the respective endpoint. If nothing is returned a 404 will be served.

## To compile and run golang app locally

Reminder: Make sure your GOPATH is set

`go get github.com/gorilla/mux`

`go build -a -installsuffix cgo -o main .`

`./main'

## To rebuild and rerun docker container

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`

`docker-compose build`

`docker-compose up`

