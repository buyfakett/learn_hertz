ARG PLATFORM=linux/amd64
ARG ARCH=amd64
ARG ALPINE_VERSION=3.21
ARG GO_VERSION=1.24.2

FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

WORKDIR /app

COPY . .
RUN apk add --no-cache gcc g++ make libc-dev && \
    go mod tidy && \
    mkdir -p static/ && echo "<h1>hertz service</h1>" > static/index.html && \
    CGO_ENABLED=1 GOOS=${PLATFORM} go build -ldflags="-s -w" -o /app/hertz_service

WORKDIR /app

FROM --platform=${PLATFORM} alpine:${ALPINE_VERSION} AS final

COPY --from=builder /app/hertz_service /app/hertz_service

EXPOSE 8888
ENTRYPOINT [ "/app/hertz_service" ]
