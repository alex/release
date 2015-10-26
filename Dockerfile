FROM ubuntu:15.04

RUN apt-get -y update

RUN apt-get -y install docker.io git python

RUN apt-get install -y curl unzip

RUN curl https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz -O
RUN tar -C /usr/local -xzf go1.5.linux-amd64.tar.gz

WORKDIR /tmp
RUN curl -Ls 'https://api.equinox.io/1/Applications/ap_y4Se864kD0m4rFttBjDpTeahC1/Updates/Asset/equinox.zip?os=linux&arch=386&channel=stable' -o equinox.zip
RUN unzip equinox.zip
RUN cp equinox /usr/bin/equinox

RUN apt-get -y install jq make python python-pip zip
RUN pip install awscli

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin:/go/src/github.com/convox/release/bin

RUN go get github.com/ddollar/rerun
RUN go get github.com/jteeuwen/go-bindata/...

COPY bin/git-restore-mtime /usr/bin/git-restore-mtime

EXPOSE 5000

COPY . /go/src/github.com/convox/release
WORKDIR /go/src/github.com/convox/release
RUN go get ./...

CMD ["bin/web"]
