The goal of this repo is to demo a golang development environment running in containers.

To use:

    $ cd $GOPATH/src
    $ git clone TODO-put-url-of-this-project
    $ cd dockerized-go-app
    $ docker-compose up

The app used to demo this capability is a small "Hello World" web app.

Current capabilities:

- run.sh: start the server using hot reloading with "gin".
-- Starts the gin server on port 3000 inside a docker container.
-- Changes made to app.go are automatically reloaded on the next http request.
- vendor.sh: run every time third party go dependencies change, then check in the changed vendor folder.

So far this has only been tested on linux, would probably need further instructions to get X windows working on Mac/Win.

Refresh database schema:
  $ docker-compose rm db

TODOs:

- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
-- Right now VS Code running inside docker starts and displays the GUI, but crashes on my system when opening a folder with no error message.  I spent a few hours trying to experiment/debug but was unable to get this working.