###########
# builder #
###########

FROM golang:1.19-buster AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/gorun \
    -ldflags='-w -s -extldflags "-static"' \
    ./cmd \
 && upx-ucl --best --ultra-brute ./bin/gorun

###########
# release #
###########

FROM golang:1.19-buster AS release
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    git

COPY --from=builder /build/bin/gorun /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/gorun"]
