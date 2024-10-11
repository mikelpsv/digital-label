FROM golang:1.23-alpine
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_CPPFLAGS="-I/usr/include" \
    UID=0 GID=0 \
    CGO_CFLAGS="-I/usr/include" \
    CGO_LDFLAGS="-L/usr/lib -lpthread -lrt -lstdc++ -lm -lc -lgcc -lz " \
    PKG_CONFIG_PATH="/usr/lib/pkgconfig"

WORKDIR /app
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./templates ./templates
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.env ./.env

RUN go mod vendor

RUN go test ./...

# set version from git
RUN go build -v \
    -o /app/service \
    ./cmd/service

CMD ["/app/service"]