FROM alpine:latest

COPY entrypoint.sh server config.yaml /server/

WORKDIR /server

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]