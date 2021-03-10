FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

EXPOSE 8080

USER 1001

CMD ["/app/basketclient"]