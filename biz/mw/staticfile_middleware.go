package mw

import (
	"context"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// StaticFileMiddleware 处理静态文件请求，并支持 fallback 到 index.html（用于 SPA）
func StaticFileMiddleware(staticFS fs.FS) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		filePath := string(ctx.Path())

		// 跳过 API 路由
		if strings.HasPrefix(filePath, "/api") {
			ctx.Next(c)
			return
		}

		if filePath == "" || filePath == "/" {
			filePath = "/index.html"
		}

		fullPath := filepath.Join("static", filePath)
		indexPath := filepath.Join("static", "index.html")

		// 返回指定路径的文件
		if serveFileFromFS(ctx, staticFS, fullPath) {
			ctx.Abort()
			return
		}

		// fallback 到 index.html
		if serveFileFromFS(ctx, staticFS, indexPath) {
			hlog.Debugf("文件 %s 不存在，使用 index.html 代替", fullPath)
			ctx.Abort()
			return
		}

		hlog.Infof("文件 %s 和 index.html 都不存在，返回 404", fullPath)
		ctx.String(http.StatusNotFound, "404 not found")
		ctx.Abort()
	}
}

func serveFileFromFS(ctx *app.RequestContext, filesystem fs.FS, path string) bool {
	file, err := filesystem.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return false
	}

	// 设置 content-type
	contentType := mime.TypeByExtension(filepath.Ext(path))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	ctx.Response.Header.Set("Content-Type", contentType)
	ctx.Data(http.StatusOK, contentType, data)
	return true
}
