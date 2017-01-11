#!/bin/bash

cd /go/src/dockerized-go-app/rest-api
exec reflex -r '\.go$' -- go test api/api_test.go