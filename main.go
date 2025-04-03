/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/swagger"
	_ "github.com/hertz-contrib/swagger/example/basic/docs"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	// The default listening port is 8888.
	// You can modify it with server.WithHostPorts().
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithMaxRequestBodySize(20<<20),
		server.WithTransport(standard.NewTransporter),
	)

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})
	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}
