ARG PLATFORM=linux/amd64
ARG ALPINE_VERSION=3.21
ARG GO_VERSION=1.24.2
ARG AUTHOR=buyfakett
ARG FRONTEND=learn_modern
ARG SERVER_NAME=hertz_service

# 支持多平台构建
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

# 前端
FROM --platform=${PLATFORM} node:22-alpine as webui
ARG AUTHOR
ARG FRONTEND
ARG repo_url=https://github.com/${AUTHOR}/${FRONTEND}
ARG branch_name=main
WORKDIR /app
RUN set -eux; \
    apk add --no-cache git; \
    git clone --depth 1 --branch "$branch_name" "$repo_url"; \
    cd ${FRONTEND}; \
    npm i -g pnpm; \
    pnpm i; \
    pnpm build; \
    mv dist ../static

# 后端
FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

ARG PLATFORM
ARG ALPINE_VERSION
ARG GO_VERSION
ARG SERVER_NAME

WORKDIR /app

COPY . .
COPY --from=webui /app/static ./static

# 根据平台推导出 GOOS 和 GOARCH
RUN set -eux; \
    apk add --no-cache gcc g++ make libc-dev; \
    TARGETOS=${TARGETOS:-linux}; \
    TARGETARCH=${TARGETARCH:-amd64}; \
    echo "Building for TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH}"; \
    CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/${SERVER_NAME}

# 最小编译
FROM --platform=${PLATFORM} alpine:${ALPINE_VERSION} AS final

ARG SERVER_NAME

COPY --from=builder /app/${SERVER_NAME} /app/${SERVER_NAME}

EXPOSE 8888
ENTRYPOINT ["/app/hertz_service"]
