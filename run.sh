#!/bin/sh

docker run -v $(pwd):/go/src/app -p 3000:3000 kpage/golang-gin
