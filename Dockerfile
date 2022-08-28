# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine AS build
WORKDIR /src
ENV GOMODCACHE /root/.cache/gocache
RUN \
  --mount=target=. \
  --mount=target=/root/.cache,type=cache \
  CGO_ENABLED=0 go build -o /out/playground -ldflags '-s -d -w' ./cmd/playground

FROM scratch
COPY --from=build /out/playground /usr/local/bin/playground
ENTRYPOINT ["playground"]
