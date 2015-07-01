FROM convox/alpine:3.1

RUN apk-install docker git

COPY bin/release /usr/local/bin/release

CMD ["/usr/local/bin/release"]
