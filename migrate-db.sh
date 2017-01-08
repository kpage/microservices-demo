#!/bin/sh

docker-compose up db-wait
docker-compose up db-migrations