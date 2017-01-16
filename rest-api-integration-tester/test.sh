#!/bin/bash

cd /go/src/microservices-demo/rest-api
exec reflex -r '\.go$' -- go test api/api_test.go

# TODO: this should also re-run integration tests when new SQL migrations are added.  Could mount the migrations folder
# and watch for changes to .sql files.  Probably this would have to pause for a few seconds then wait for the migration 
# container to stop, to make sure the tests only run after the migration.