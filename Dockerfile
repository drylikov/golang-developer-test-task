# Image page: <https://hub.docker.com/_/golang>
FROM golang:1.13-alpine as builder
COPY . /src
WORKDIR /src
RUN go build -ldflags="-s -w" -o /tmp/app .

FROM alpine:latest
COPY --from=builder /tmp/app /bin/app
ENTRYPOINT ["/bin/app"]
