FROM golang:alpine as builder
LABEL maintainer="Lakkini Music <lakkinzimusic@gmail.com>"

RUN mkdir /app 

ENV PORT 8050
EXPOSE $PORT

ADD . /app/ 
WORKDIR /app 

RUN go build -o main . 
CMD ["/app/main"]

