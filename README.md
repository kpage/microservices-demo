The goal of this repo is to demo a development environment running in docker-compose.

To use, first install docker and docker-compose.  Then:

    $ cd $GOPATH/src
    $ git clone https://github.com/kpage/dockerized-go-app
    $ cd dockerized-go-app
    $ docker-compose up

It's not required to put this project in your GOPATH to run it from the command line, but if you want
to edit any golang from an IDE I recommend you put it there so the go code will compile in the IDE.

Current capabilities:

- Changes made to .go files will reload the server automatically
- govendor.sh: run every time third party go dependencies change, then check in the changed vendor folder.

## REST API

The REST API is written in golang.

#### REST API Tests:

The REST API tests are designed to run in the docker-compose environment against a live server.  If you start with "docker-compose up", the tests will continuously
run when any go code is changed.  Tests are written in go in api_test.go and are run from the rest-api-integration-tester container.

## Database

Runs mariadb in a container.  If you want to connect directly to the db, it exposes port 3336 on your host.

To drop the database and refresh the schema:

    $ docker-compose stop db
    $ docker-compose rm db
    $ docker-compose up -d db
    $ docker-compose up -d db-migrations

## Web client

TODO: add a frontend web client in HTML & JS to demonstrate using the REST API.  The build is working right now, but the code was
scavenged from an unrelated project that does not actually exercise the API in this project yet.

TODO: node_modules: store base node_modules image that is usable by prod and ensures same modules for all devs.  Right now, node_modules will be created for each dev separately, although
this should not be a problem due to yarn.lock

#### Updating node packages:

- Usually you don't need to edit package.json or yarn.lock directly.  You can use yarn commands to manage packages.
- After running any yarn command, always do the following:

    $ docker-compose build web-client
    $ docker-compose restart web-client

- When satisfied with changes, check in both package.json and yarn.lock
- TODO: create command to publish a new web-client base image with latest node_modules inside and update 
  Dockerfile to use it
- Example, to upgrade a single existing package to a specified version (example, upgrade webpack to 1.13.3):

    $ docker-compose run web-client yarn upgrade webpack@1.13.3
    $ docker-compose build web-client
    $ docker-compose restart web-client

- To add a new package:

    $ docker-compose run web-client yarn add redux
    
- You can see all the yarn commands here: https://yarnpkg.com/en/docs/cli

## Architecture of this app:
- rest-api: a REST API implemented in golang
- db: a mariadb database

## Misc TODOs:

- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
-- Right now VS Code running inside docker starts and displays the GUI, but crashes on my system when opening a folder with no error message.  I spent a few hours trying to experiment/debug but was unable to get this working.
- node should not proxy to rest-api directly, create a simple nginx service to handle all incoming requests.  Example, / loads static resources from web-client, or anything starting with /api goes to rest-api.  This would reduce need to expose so many ports on host, which could cause annoying conflicts.