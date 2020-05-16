FROM golang:alpine as builder
LABEL maintainer="Lakkini Music <lakkinzimusic@gmail.com>"

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o horse-maze


ENV PORT 8050
EXPOSE $PORT

# ADD . /app/ 
# WORKDIR /app 

# RUN go build -o main . 
# CMD ["/app/main"]
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/horse-maze .

CMD ["./horse-maze"]