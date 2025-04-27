#!/bin/bash
set -e

WORKDIR=$(pwd)
SERVER_NAME=hertz_service

# 创建静态文件目录及默认页面
# mkdir -p static/
# if [ ! -f static/index.html ]; then
#     echo "<h1>hertz service</h1>" > static/index.html
# fi

AUTHOR=buyfakett
FRONTEND=learn_modern
repo_url=https://github.com/${AUTHOR}/${FRONTEND}
branch_name=main
git clone --depth 1 --branch "$branch_name" "$repo_url"
cd ${WORKDIR}/${FRONTEND}/
npm i -g pnpm
pnpm i
pnpm build
mv dist ../static
cd ${WORKDIR}/
rm -rf ${WORKDIR}/${FRONTEND}/

# 检查依赖工具，只在 Linux 上执行
if [ "$(uname)" = "Linux" ]; then
  # 使用 sudo 以避免权限问题
  if ! command -v xz &> /dev/null; then
      sudo apt update && sudo apt install -y xz-utils
  fi
  if ! command -v md5sum &>/dev/null; then
      sudo apt update && sudo apt install -y coreutils
  fi
fi

# 下载依赖
go mod download

# —— 定义多平台编译目标 ——
# 接受外部 GOOS/GOARCH，否则默认四平台
if [ -n "$GOOS" ] && [ -n "$GOARCH" ]; then
    platforms=("$GOOS/$GOARCH")
else
    platforms=(
        "linux/amd64"
        "linux/arm64"
        "darwin/amd64"
        "darwin/arm64"
    )
fi
# —— 编译流程 ——
mkdir -p dist/release
for platform in "${platforms[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}

    # 使用纯Go的SQLite实现，禁用CGO
    CGO_ENABLED=0

    # 生成文件名
    BINARY="${SERVER_NAME}_${GOOS}_${GOARCH}"
    [ "$GOOS" = "windows" ] && BINARY="${BINARY}.exe"
    OUTPUT="dist/release/${BINARY}"

    echo "编译：${GOOS}-${GOARCH} (CGO_ENABLED=$CGO_ENABLED)..."
    env GOOS="$GOOS" GOARCH="$GOARCH" CGO_ENABLED="$CGO_ENABLED" \
        go build -ldflags '-w -s' -o "$OUTPUT"

    if [ -f "$OUTPUT" ]; then
        echo "打包：${BINARY}.tar.xz"
        tar -cJf "${OUTPUT}.tar.xz" -C "$(dirname "$OUTPUT")" "$(basename "$OUTPUT")"
    else
        echo "错误：${OUTPUT} 未生成" >&2
        exit 1
    fi

done

# 生成 MD5 校验文件：macOS 与 Linux 区分处理
if [ "$(uname)" = "Darwin" ]; then
    for f in dist/release/*; do
        [ -f "$f" ] || continue
        md5 -r "$f" > "$f.md5"
    done
else
    for f in dist/release/*; do
        [ -f "$f" ] || continue
        md5sum "$f" > "$f.md5"
    done
fi
