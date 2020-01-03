FROM alpine:latest

WORKDIR /

ADD main  /

EXPOSE 3100

ENTRYPOINT ["./main"]
