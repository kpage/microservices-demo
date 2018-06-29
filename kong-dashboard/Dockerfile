FROM pgbi/kong-dashboard:v3.3.0

# Parent image is based on alpine, which does not support wget to https without openssl
RUN apk --no-cache add openssl
# dumb-init makes it easier to use interactive containers, e.g. hit Ctrl+C to kill a container
RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64
RUN chmod +x /usr/local/bin/dumb-init

ENTRYPOINT ["/usr/local/bin/dumb-init", "--", "./docker/entrypoint.sh"]