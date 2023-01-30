// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/errno"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		handler.SendResponse(c, errno.ErrBind)
		return
	}

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
	}

	handler.SendResponse(c, resp)
}

// Login .
// @router /douyin/user/login [POST]
// 使用 hertz-jwt 时需要替换掉这个 LoginHandler
//func Login(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinUserRegisterRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinUserRegisterResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}

// GetUserById .
// @router /douyin/user [GET]
func GetUserById(ctx context.Context, c *app.RequestContext) {
	// 如果是需要授权访问的接口，则在进入时已经被中间件从 body 中获取 token 解析完成了，这里无需额外解析
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind)
		return
	}

	resp, err := rpc.GetUserById(ctx, &user.DouyinUserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	handler.SendResponse(c, resp)
}
