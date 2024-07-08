FROM --platform=$BUILDPLATFORM golang:1.23rc1 as builder
RUN mkdir /storage

WORKDIR /go/src/github.com/morinoparty/suspension_bridge

COPY ./go.* ./

RUN --mount=type=cache,target=/go/pkg/mod go mod download

ENV GOCACHE=/tmp/go/cache
ENV CGO_ENABLED=0

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/tmp/go/cache \
    go build -o /suspension_bridge

FROM gcr.io/distroless/static-debian11
WORKDIR /app
EXPOSE 25565

COPY --from=builder /storage/ /app/storage/
VOLUME /app/storage

COPY --from=builder /suspension_bridge ./
CMD ["./suspension_bridge"]
