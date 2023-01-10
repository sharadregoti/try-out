FROM golang:1.18-alpine3.14
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
# ENV GO111MODULE=off
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app .

FROM alpine:3.12
WORKDIR /app
COPY --from=0 /build/app .
COPY --from=0 /go/bin/dlv .
# CMD ["./app"]
CMD ["./dlv", "--listen=:4040", "--headless=true", "--api-version=2", "exec", "./app"]