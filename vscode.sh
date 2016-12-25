#!/bin/sh

# vscode user environment variables need to be set inline because
# we are changing user...
docker run -it \
       -e DISPLAY \
       -v /tmp/.X11-unix:/tmp/.X11-unix \
       -v $HOME/.Xauthority:/home/developer/.Xauthority \
       --net=host \
       -v $(pwd):/home/vscode/go/src/app \
       --rm ctaggart/golang-vscode \
       su - vscode -c "export GOPATH=/home/vscode/go ; code -w go/src/app"
#       -e GOPATH=/home/vscode/app \
