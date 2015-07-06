FROM ubuntu:14.04

RUN apt-get -y update

RUN apt-get -y install docker.io git python

RUN apt-get -y install golang \
  golang-go-darwin-386 golang-go-darwin-amd64 \
  golang-go-linux-386 golang-go-linux-amd64 golang-go-linux-arm \
  golang-go-windows-386 golang-go-windows-amd64

ENV GOPATH /go
ENV PATH $PATH:/go/src/github.com/convox/release/bin

RUN apt-get install -y curl unzip

WORKDIR /tmp
RUN curl -Ls 'https://api.equinox.io/1/Applications/ap_y4Se864kD0m4rFttBjDpTeahC1/Updates/Asset/equinox.zip?os=linux&arch=386&channel=stable' -o equinox.zip
RUN unzip equinox.zip
RUN cp equinox /usr/bin/equinox

RUN go get github.com/jteeuwen/go-bindata/...

COPY bin/git-restore-mtime /usr/bin/git-restore-mtime

EXPOSE 5000

COPY . /go/src/github.com/convox/release
WORKDIR /go/src/github.com/convox/release
RUN go get .

CMD ["/go/bin/release"]
