#FROM golang:latest
#
#WORKDIR $GOPATH/src/ginServer
#COPY . $GOPATH/src/ginServer
#RUN go build .
#
#EXPOSE 3000
#
#ENTRYPOINT ["./ginserver"]

FROM scratch

WORKDIR $GOPATH/src/ginServer
COPY . $GOPATH/src/ginServer

EXPOSE 3000
CMD ["./ginserver"]