FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /tmp/request-service

COPY go.mod .
COPY go.sum .
RUN go clean --modcache
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./out/request-service .

FROM scratch
LABEL maintainer="Julian Klaiber"

COPY --from=build_base /tmp/request-service/out/request-service /usr/bin/request-service

ENTRYPOINT [ "/usr/bin/request-service" ]
