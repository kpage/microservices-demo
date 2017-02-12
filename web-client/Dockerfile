FROM node:6.9.2

# dumb-init makes it easier to use interactive containers, e.g. hit Ctrl+C to kill a container
RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64
RUN chmod +x /usr/local/bin/dumb-init

ENV HOME=/home/node
ARG NODE_ENV
ENV NODE_ENV $NODE_ENV

EXPOSE 4000

RUN curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version 0.19.1
ENV PATH="$HOME/.yarn/bin:$PATH"

# also add .yarnrc, if present (if we want to use our nexus to proxy yarn)
COPY package.json yarn.lock $HOME/app/

# Make sure node has access to his own home dir
RUN chown -R node:node $HOME

# We use the "node" user because this will be the same user that runs the
# container, so any files created by 'node' user during the gulp build will
# have the correct permissions on the host.  This only works if the user
# running docker is uid 1000, which should be true on most single-user dev
# machines
USER node

WORKDIR $HOME/app
RUN yarn install

COPY entrypoint.sh $HOME/app/

ENTRYPOINT ["/usr/local/bin/dumb-init", "--", "sh", "entrypoint.sh"]

CMD ["node", "server.js"]