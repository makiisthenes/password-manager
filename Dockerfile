FROM golang:1.22.4

WORKDIR /app

ADD password-manager /app/
EXPOSE 7001-8001

ENTRYPOINT ["./password-manager"]