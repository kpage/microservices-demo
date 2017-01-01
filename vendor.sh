#!/bin/sh

docker run --rm -v $PWD:/app -w /app treeder/go vendor
mv ./vendor/src/* ./vendor
rm -rf ./vendor/src
rm -rf ./vendor/pkg
