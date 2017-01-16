#!/bin/sh

docker run --rm -it -v $PWD:/go/src/microservices-demo -w /go/src/microservices-demo trifs/govendor fetch +missing
docker run --rm -it -v $PWD:/go/src/microservices-demo -w /go/src/microservices-demo trifs/govendor remove +unused
