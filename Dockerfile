ARG PLATFORM=linux/amd64
ARG ARCH=amd64
ARG GOOS=linux
ARG GOARCH=amd64
ARG ALPINE_VERSION=3.21
ARG GO_VERSION=1.24.2
ARG AUTHOR=buyfakett
ARG FRONTEND=learn_modern

FROM node:22 as webui
ARG AUTHOR
ARG FRONTEND
ARG repo_url=https://github.com/${AUTHOR}/${FRONTEND}
ARG branch_name=main
WORKDIR /app
RUN set -eux; \
    git clone --depth 1 --branch "$branch_name" "$repo_url"; \
    cd ${FRONTEND}; \
    npm i -g pnpm; \
    pnpm i; \
    pnpm build; \
    mv dist ../static

FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

ARG GOOS
ARG GOARCH

WORKDIR /app

COPY . .

COPY --from=webui /app/static ./static

RUN apk add --no-cache gcc g++ make libc-dev && \
    go mod download && \
    CGO_ENABLED=1 GOOS=${PLATFORM} go build -ldflags="-s -w" -o /app/hertz_service

WORKDIR /app

FROM --platform=${PLATFORM} alpine:${ALPINE_VERSION} AS final

COPY --from=builder /app/hertz_service /app/hertz_service

EXPOSE 8888
ENTRYPOINT [ "/app/hertz_service" ]
