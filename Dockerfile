FROM convox/alpine:3.1

RUN apk-install docker go git

ENV GOPATH /go
ENV PATH $PATH:/go/src/github.com/convox/release/bin

COPY . /go/src/github.com/convox/release
WORKDIR /go/src/github.com/convox/release
RUN go get .

CMD ["/go/bin/release"]
