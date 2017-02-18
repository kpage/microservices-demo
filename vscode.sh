#!/bin/sh

exec docker run -it \
       -e DISPLAY \
       --net=host \
       -v $(pwd):/home/vscode/go/src/microservices-demo \
       --rm kpage/golang-vscode \
       su - vscode -c "code --verbose --disable-gpu --wait go/src/microservices-demo"

# Here's an alternative approach, instead of loading VSCode directly, this will open a shell in a GUI window.
# VS Code can then be launched with "code".  The shell will stay open as root process, possibly getting around
# the issue I had before where VS Code changes process IDs when opening folders and this somehow causes the container
# to shut down.
# exec docker run -d -d \
#     -v /tmp/.X11-unix:/tmp/.X11-unix:rw \
#     -v ${PWD}:/developer/project \
#     -e DISPLAY=unix${DISPLAY} \
#     -p 5000:5000 \
#     --device /dev/snd \
#     cmiles74/docker-vscode


#       -v /tmp/.X11-unix:/tmp/.X11-unix \
#       -v $HOME/.Xauthority:/home/developer/.Xauthority \
#       --device /dev/snd \

# su - vscode -c "code -w go/src/app"
