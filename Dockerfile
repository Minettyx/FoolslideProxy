FROM golang:1.21.0 AS build-stage
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /foolslideproxy cmd/foolslideproxy/foolslideproxy.go 

FROM alpine:latest AS build-release-stage 
WORKDIR /
COPY --from=build-stage /foolslideproxy /foolslideproxy
EXPOSE 3333

ENTRYPOINT ["/foolslideproxy"]
