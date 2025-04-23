#!/bin/bash
RUN_NAME=hertz_service
mkdir -p static/
if [ ! -f static/index.html ]; then
    echo "<h1>hertz service</h1>" > static/index.html
fi
go mod tidy
go build -ldflags '-w -s' -o ${RUN_NAME}