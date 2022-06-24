FROM alpine:latest

RUN apk add --no-cache --upgrade ca-certificates tzdata curl

WORKDIR /app

COPY ./cmd/build/. ./

HEALTHCHECK --start-period=4s --interval=5s --timeout=2s --retries=2 CMD curl -f http://localhost/healthcheck || false

CMD ["./svc"]
