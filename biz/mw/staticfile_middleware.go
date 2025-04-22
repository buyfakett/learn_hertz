package mw

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func StaticFileMiddleware(staticDir string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		filePath := string(ctx.Path())

		// 跳过 /api 开头的路径
		if strings.HasPrefix(filePath, "/api") {
			ctx.Next(c)
			return
		}

		if filePath == "" || filePath == "/" {
			filePath = "/index.html"
		}

		fullPath := filepath.Join(staticDir, filePath)
		indexPath := filepath.Join(staticDir, "index.html")

		// fullPath 是否存在
		if _, err := os.Stat(fullPath); err == nil {
			ctx.File(fullPath)
			ctx.Abort()
			return
		}

		// index.html 是否存在
		if _, err := os.Stat(indexPath); err == nil {
			hlog.Debugf("文件 %s 不存在，使用 index.html 代替", fullPath)
			ctx.File(indexPath)
			ctx.Abort()
			return
		}

		// 全部找不到
		hlog.Infof("文件 %s 和 index.html 都不存在，返回 404", fullPath)
		ctx.String(http.StatusNotFound, "404 not found")
		ctx.Abort()
	}
}
