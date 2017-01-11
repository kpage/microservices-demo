FROM golang:wheezy

# dumb-init makes it easier to use interactive containers, e.g. hit Ctrl+C to kill a container
RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64
RUN chmod +x /usr/local/bin/dumb-init

RUN mkdir -p /go/src
WORKDIR /go/src

RUN go get github.com/cespare/reflex

ADD test.sh /go/src

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]

CMD ["./test.sh"]