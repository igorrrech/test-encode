FROM golang:1.23-alpine as builder
WORKDIR /build
RUN pwd
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main ./cmd/main.go
FROM alpine:3
COPY config.json config.json
COPY --from=builder main /bin/main
ENTRYPOINT [ "/bin/main" ]