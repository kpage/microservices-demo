The goal of this repo is to demo a development environment running in docker-compose.

NOTE: this is a work in progress and may not yet be in a fully working state.

To use, first install docker and docker-compose.  Then:

    $ cd $GOPATH/src
    $ git clone https://github.com/kpage/microservices-demo
    $ cd microservices-demo
    $ docker-compose up

If docker-compose stops with a message like 'Container "30a7960c4c89" is unhealthy', you can just re-run 'docker-compose up'.  Usually this means that a database migration is taking some time and a container timed out waiting to start.

Note: it's not required to put this project in your GOPATH to run it from the command line, but if you want
to edit any golang from an IDE I recommend you put it there so the go code will compile in the IDE.

## REST API

The REST API is written in golang.

The API is self-documenting using the HAL hypermedia format.  If you request any api url with header "Accept: application/hal+json" you will get a json response, otherwise
you will get an HTML page with documentation about that url.

Third-party dependencies are handled by "vendoring".  This means copying the source of the dependency into the "vendor" folder
and checking it in to git.

If you need to add another third-party dependency, first add it to the "import" section of a .go file.  Example:

```
import (
	"github.com/gorilla/mux"
)
```

Some editors such as Visual Studio Code automatically remove unused imports, so you might need to also use the dependency
in your code:

```
r := mux.NewRouter()
```

Next, run our vendoring script:

```
    $ cd rest-api
    $ ./govendor.sh
    $ sudo chown -R `whoami`:`whoami` vendor
    $ git status
```
TODO: creates files owned by root right now, fix govendor.sh to create files as the current user

You should now see your dependency in the "vendor" directory and your code should compile correctly.  You might have to
restart Visual Studio Code to get it to recognize the new vendor package.

Make sure to add your changes to git:

```
    $ git add rest-api/vendor
    $ git commit
```

Note on REST API documentation: I chose to separate documentation concerns by having the core REST API code just return
responses in a hal+json format that could be read by any compatible API browser (HAL browser, Swagger, etc.).  Currently
there is a small nginx proxy (rest-api-proxy) that sits in front of rest-api to handle showing this documentation.  This
might be better served as a Kong plugin later.  The idea behind this separation of concerns was to avoid having repeated logic in every backend service that has to serve up documentation pages.

#### REST API Tests:

The REST API tests are designed to run in the docker-compose environment against a live server.  If you start with "docker-compose up", the tests will continuously
run when any go code is changed.  Tests are written in go in api_test.go and are run from the rest-api-integration-tester container.

#### REST API authorization flow

TODO: enable https in all environments

Create a new user like this:

```
curl --data 'username=testuser&password=testpassword' 'http://localhost/auth/accounts'
```

Get a token like this:

```
curl --data 'username=testuser&password=testpassword' 'http://localhost/auth/token'
```

```
{"refresh_token":"f7eab65c2d3e43ddb83cfad581bca0bd","token_type":"bearer","access_token":"9b94cf2cad2f47d9a0c641d6fe13a966","expires_in":7200}
```

Pass token when making requests to the REST API:

```
curl -X GET --data 'access_token=9b94cf2cad2f47d9a0c641d6fe13a966' 'http://localhost/api/orders'
```

## Database

Runs mariadb in a container.  If you want to connect directly to the db, it exposes port 3336 on your host.

To drop the database and refresh the schema:

```
    $ docker-compose stop db
    $ docker-compose rm db
    $ docker-compose up -d db
    $ docker-compose up -d db-migrations
```

Database migrations are handled with flyway.  They will run automatically after docker-compose up.  If you add a new migration,
you can run it with:

```
    $ docker-compose up -d db-migrations
```

TODO: for production, db migrations could maybe be handled as Kubernetes Jobs.  For example, when a new migration needs to
be run, we could have a job that snapshots the db and runs flyway in a migration container containing the migration scripts.
This could be deployed separately from the other services to run migrations in preparation for new code changes.

## Web client

TODO: add a frontend web client in HTML & JS to demonstrate using the REST API.  The build is working right now, but the code was
scavenged from an unrelated project that does not actually exercise the API in this project yet.

TODO: consider npm library "traverson" for consuming HAL+JSON instead of follow.js

TODO: node_modules: store base node_modules image that is usable by prod and ensures same modules for all devs.  Right now, node_modules will be created for each dev separately, although
this should not be a problem due to yarn.lock

#### Updating node packages:

- Usually you don't need to edit package.json or yarn.lock directly.  You can use yarn commands to manage packages.
- After running any yarn command, always do the following:

```
    $ docker-compose build web-client
    $ docker-compose restart web-client
```

- When satisfied with changes, check in both package.json and yarn.lock
- TODO: create command to publish a new web-client base image with latest node_modules inside and update 
  Dockerfile to use it
- Example, to upgrade a single existing package to a specified version (example, upgrade webpack to 1.13.3):

```
    $ docker-compose run web-client yarn upgrade webpack@1.13.3
    $ docker-compose build web-client
    $ docker-compose restart web-client
```

- To add a new package:

```
    $ docker-compose run web-client yarn add redux
```

- You can see all the yarn commands here: https://yarnpkg.com/en/docs/cli

## API Gateway (Kong)

This app uses kong for an API gateway (routing URIs to the correct backend service, providing API security, etc.)

More about kong: https://getkong.org/

The config for kong is stored at kong/config.yml

The kong web dashboard runs at:

http://localhost:9999

To play with kong settings, you can change things in the dashboard and then export the current state with:

```
    $ docker-compose run kongfig-dump
```

You can then copy desired settings to kong/config.yml

## Architecture of this app:
- kong: API gateway for routing, authentication
- rest-api: a REST API implemented in golang
- db: a mariadb database
- web-client: an HTML single-page app using react+redux

## Misc TODOs:

- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
-- Right now VS Code running inside docker starts and displays the GUI, but crashes on my system when opening a folder with no error message.  I spent a few hours trying to experiment/debug but was unable to get this working.
- order.price includes currency, maybe would be better to separate currency into separate field?
- Make initial db migration script re-runnable
- Add a kong plugin or service at / to provide a HAL root to the various apis like /api and /auth
