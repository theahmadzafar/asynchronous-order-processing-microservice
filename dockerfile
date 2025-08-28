FROM golang:1.24.5-alpine3.22 AS build

ARG TAG
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod download

RUN set -x; apk add --no-cache && CGO_ENABLED=0 go build -ldflags="-s -w -X asynchronous-order-processing-microservice/internal/transport/http/handlers.Tag=$TAG" -o ./bin/app cmd/main.go

FROM alpine:3.22 as run

WORKDIR /app

COPY --from=build /app/bin .
COPY --from=build /app/config.yaml .

RUN chmod +x ./app

ENTRYPOINT ["./app"]