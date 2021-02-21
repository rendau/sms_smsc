FROM golang:1.15 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make



FROM alpine:latest

RUN apk --no-cache update && apk --no-cache upgrade && apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/cmd/build/* ./

CMD ["./svc"]
