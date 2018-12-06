FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/pbgrpc

ADD . /go/src/github.com/emurmotol/pbgrpc

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/emurmotol/pbgrpc/bugs/cmd  -watch github.com/emurmotol/pbgrpc/bugs
