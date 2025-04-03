// Code generated by hertz generator.

package test

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	test "hertz_demo/biz/model/api/test"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req test.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(test.HelloResp)

	resp.RespBody = "hello," + req.Name

	c.JSON(consts.StatusOK, resp)
}
