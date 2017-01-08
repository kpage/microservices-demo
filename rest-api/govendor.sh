#!/bin/sh

docker run --rm -it -v $PWD:/go/src/dockerized-go-app -w /go/src/dockerized-go-app trifs/govendor fetch +missing
docker run --rm -it -v $PWD:/go/src/dockerized-go-app -w /go/src/dockerized-go-app trifs/govendor remove +unused
