#!/bin/sh

docker run -v $(pwd):/go/src/app kpage/golang-gin
