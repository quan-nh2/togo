FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY ./app .
RUN chmod +x ./app

COPY ./config/config.yaml .
ENTRYPOINT ["./app"]