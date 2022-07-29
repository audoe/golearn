FROM golang:alpine as builder

WORKDIR $GOPATH/src/github.com/audoe/golearn
COPY . $GOPATH/src/github.com/audoe/golearn

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-server .

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
EXPOSE 8090
CMD ["./http-server"]


#