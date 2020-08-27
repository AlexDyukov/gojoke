############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# install git to fetch dependencies and gcc for runtime/cgo
RUN apk update && apk add --no-cache git build-base

WORKDIR /root
COPY . .

# Run unit tests
RUN go test -v

# Build
RUN go build -o /bin/gojoke

############################
# STEP 2 build a small image
############################
FROM alpine

# ssl support require ca-certificates
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# default env vars
ENV PG_URL postgresql://postgres:postgres@postgres:5432/postgres?pool_max_conns=5

COPY --from=builder /bin/gojoke /gojoke

# Run app
ENTRYPOINT ["/gojoke"]
