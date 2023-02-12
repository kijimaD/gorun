###########
# builder #
###########

FROM golang:1.19-buster AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .

WORKDIR /build/cmd
RUN GO111MODULE=on CGO_ENABLED=0 go build \
      -ldflags='-w -s -extldflags "-static"' \
      -o /build/bin/gorun \
 && upx-ucl --best --ultra-brute /build/bin/gorun

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
