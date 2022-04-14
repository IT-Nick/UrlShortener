FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/IT-Nick/urlSortener
RUN cd /build && git clone https://github.com/IT-Nick/urlSortener.git

RUN cd /build/urlSortener && go build

EXPOSE 8000
