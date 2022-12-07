FROM golang:1.17-alpine3.14
WORKDIR /build
COPY . .
#RUN apk --no-cache add build-base
# RUN GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o app .
# RUN go build -gcflags "all=-N -l" -o app .
# RUN go install github.com/go-delve/delve/cmd/dlv@v1.9.1
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
ENV GO111MODULE=off
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app .

FROM alpine:3.12
WORKDIR /app
COPY --from=0 /build/app .
COPY --from=0 /go/bin/dlv .
# CMD ["./app"]
CMD ["./dlv", "--listen=:4040", "--headless=true", "--api-version=2", "exec", "./app"]