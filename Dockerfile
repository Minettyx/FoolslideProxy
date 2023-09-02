FROM golang:1.21.0
WORKDIR /usr/src/app
COPY . .
RUN go build cmd/foolslideproxy/foolslideproxy.go
CMD ["./foolslideproxy"]
EXPOSE 3333
