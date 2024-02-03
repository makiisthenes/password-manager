FROM golang:1.22rc2-bookworm

WORKDIR /app

ADD password-manager /app/
EXPOSE 7001-8001

ENTRYPOINT ["./password-manager"]