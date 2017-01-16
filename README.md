The goal of this repo is to demo a development environment running in docker-compose.

To use, first install docker and docker-compose.  Then:

    $ cd $GOPATH/src
    $ git clone https://github.com/kpage/microservices-demo
    $ cd microservices-demo
    $ docker-compose up

It's not required to put this project in your GOPATH to run it from the command line, but if you want
to edit any golang from an IDE I recommend you put it there so the go code will compile in the IDE.

Current capabilities:

- Changes made to .go files will reload the server automatically
- govendor.sh: run every time third party go dependencies change, then check in the changed vendor folder.

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
    $ git status
```

You should now see your dependency in the "vendor" directory and your code should compile correctly.  You might have to
restart Visual Studio Code to get it to recognize the new vendor package.

Make sure to add your changes to git:

```
    $ git add rest-api/vendor
    $ git commit
```
#### REST API Tests:

The REST API tests are designed to run in the docker-compose environment against a live server.  If you start with "docker-compose up", the tests will continuously
run when any go code is changed.  Tests are written in go in api_test.go and are run from the rest-api-integration-tester container.

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

## Architecture of this app:
- router & load balancer: nginx
- rest-api: a REST API implemented in golang
- db: a mariadb database
- web-client: an HTML single-page app using react+redux

## Misc TODOs:

- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
-- Right now VS Code running inside docker starts and displays the GUI, but crashes on my system when opening a folder with no error message.  I spent a few hours trying to experiment/debug but was unable to get this working.
- order.price includes currency, maybe would be better to separate currency into separate field?
- Make initial db migration script re-runnable