FROM golang:1.8.0-alpine

MAINTAINER Eric Shi <shibingli@realclouds.org>
ADD . /go/
EXPOSE 8080

CMD ["/go/realclouds"]