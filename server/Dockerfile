FROM golang:1.10

WORKDIR $GOPATH/src/github.com/alyyousuf7/kube-demo/server
ADD . .
RUN make build
ENTRYPOINT ./server

