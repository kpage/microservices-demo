The goal of this repo is to demo a golang development environment running in containers.

To use:

    $ cd $GOPATH/src
    $ git clone TODO-put-url-of-this-project
    $ cd dockerized-go-app
    $ docker-compose up

If you have not yet run migrations, then also run:

    $ ./migrate-db.sh

The app used to demo this capability is a small "Hello World" web app.

Current capabilities:

- run.sh: start the server using hot reloading with "gin".
-- Starts the gin server on port 3000 inside a docker container.
-- Changes made to app.go are automatically reloaded on the next http request.
- govendor.sh: run every time third party go dependencies change, then check in the changed vendor folder.

So far this has only been tested on linux, would probably need further instructions to get X windows working on Mac/Win.

Refresh database schema:
    $ docker-compose rm db

Updating node packages:
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


TODOs:

- node_modules: store base node_modules image that is usable by prod and ensures same modules for all devs.  Right now, node_modules will be created for each dev separately, although
  this should not be a problem due to yarn.lock
- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
-- Right now VS Code running inside docker starts and displays the GUI, but crashes on my system when opening a folder with no error message.  I spent a few hours trying to experiment/debug but was unable to get this working.
- node should not proxy to rest-api directly, create a simple nginx service to handle all incoming requests.  Example, / loads static resources from web-client, or anything starting with /api goes to rest-api.  This would reduce need to expose so many ports on host, which could cause annoying conflicts.

Architecture of this app:
- rest-api: a REST API implemented in golang
- db: a mariadb database