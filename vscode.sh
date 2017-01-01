#!/bin/sh

# vscode user environment variables need to be set inline because
# we are changing user...
docker run -it \
       -e DISPLAY \
       -v /tmp/.X11-unix:/tmp/.X11-unix \
       -v $HOME/.Xauthority:/home/developer/.Xauthority \
       --net=host \
       -v $(pwd):/home/vscode/go/src/app \
       --rm kpage/golang-vscode \
       su - vscode -c "code -w go/src/app"

# su - vscode -c "code -w go/src/app"

#       su - vscode -c "code"
       #       su - vscode -c "export GOPATH=/home/vscode/go ; code -w go/src/app"

