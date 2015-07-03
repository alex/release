FROM convox/alpine:3.1

RUN apk-install docker go git python

ENV GOPATH /go
ENV PATH $PATH:/go/src/github.com/convox/release/bin

WORKDIR /tmp
RUN curl -Ls 'https://api.equinox.io/1/Applications/ap_y4Se864kD0m4rFttBjDpTeahC1/Updates/Asset/equinox.zip?os=linux&arch=386&channel=stable' -o equinox.zip
RUN unzip equinox.zip
RUN cp equinox /usr/bin/equinox

COPY bin/git-restore-mtime /usr/bin/git-restore-mtime

COPY . /go/src/github.com/convox/release
WORKDIR /go/src/github.com/convox/release
RUN go get .

CMD ["/go/bin/release"]
