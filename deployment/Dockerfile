FROM alpine:latest

COPY deployment/entrypoint.sh deployment/server deployment/config.yaml /server/

WORKDIR /server

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]