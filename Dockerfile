FROM golang:1.26rc2

WORKDIR /app

ADD password-manager /app/
EXPOSE 7001-8001

ENTRYPOINT ["./password-manager"]