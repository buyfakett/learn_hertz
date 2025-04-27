#!/bin/bash

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

# 检查依赖工具
if ! command -v xz &> /dev/null; then
    apt update && apt install -y xz-utils
fi

if ! command -v md5sum &>/dev/null; then
    apt update && apt install -y coreutils
fi

# 下载依赖
go mod download

# 定义多平台编译目标
# platforms=(
#     "linux/amd64"
#     "linux/arm64"
#     "darwin/amd64"
#     "darwin/arm64"
#     "windows/amd64"
#     "windows/arm64"
# )

platforms=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

# 主构建流程
mkdir -p dist/release
for platform in "${platforms[@]}"; do
    # 分割平台信息
    GOOS=${platform%/*}
    GOARCH=${platform#*/}

    # 生成文件名
    BINARY="${SERVER_NAME}_${GOOS}_${GOARCH}"
    [ "$GOOS" = "windows" ] && BINARY="${BINARY}.exe"

    # 目标路径
    OUTPUT_FILE="dist/release/${BINARY}"

    # 编译
    echo "编译中: ${GOOS}-${GOARCH}..."
    env GOOS="$GOOS" GOARCH="$GOARCH" CGO_ENABLED=1 \
        go build -ldflags '-w -s' -o "$OUTPUT_FILE"

    # 压缩，仅包含可执行文件本身
    tar -cJf "${OUTPUT_FILE}.tar.xz" -C "$(dirname "$OUTPUT_FILE")" "$(basename "$OUTPUT_FILE")"
    echo "生成文件: ${OUTPUT_FILE}.tar.xz"
done

# 生成所有 dist/release 下文件的 md5
echo "生成 dist/release 下所有文件的 .md5 文件..."
for file in dist/release/*; do
    [ -f "$file" ] || continue
    md5sum "$file" > "${file}.md5"
done