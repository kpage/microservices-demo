#!/bin/sh

exec docker run -it \
       -e DISPLAY \
       --net=host \
       -v $(pwd):/home/vscode/go/src/app \
       --rm kpage/golang-vscode \
       su - vscode -c "code --verbose --disable-gpu --wait ."

#       -v /tmp/.X11-unix:/tmp/.X11-unix \
#       -v $HOME/.Xauthority:/home/developer/.Xauthority \
#       --device /dev/snd \

# su - vscode -c "code -w go/src/app"
