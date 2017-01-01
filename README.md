The goal of this repo is to demo a golang development environment running in containers.

The app used to demo this capability is a small "Hello World" web app.

Current capabilities:

- run.sh: start the server using hot reloading with "gin".
-- Starts the gin server on port 3000 inside a docker container.
-- Changes made to app.go are automatically reloaded on the next http request.
- vscode.sh: start preconfigured Visual Studio Code editor in this folder.
-- Uses X forwarding to display the GUI
- vendor.sh: run every time third party go dependencies change, then check in the changed vendor folder.

So far this has only been tested on linux, would probably need further instructions to get X windows working on Mac/Win.