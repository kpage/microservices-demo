#!/bin/sh

rm -rf ./vendor/*
docker run --rm -v $PWD:/app -w /app treeder/go vendor
mv ./vendor/src/* ./vendor
rm -rf ./vendor/src
rm -rf ./vendor/pkg
# By default this vendoring will add git submodules, we remove them
# and just check in the files directly to our repo
find ./vendor -type d -name .git -print0  | xargs -0 rm -rf
