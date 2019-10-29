package user

import (
	client "catdogs-be/client/user"
	"catdogs-be/libs"
	"catdogs-be/logging"
	pb "catdogs-be/pb"

	"github.com/gin-gonic/gin"
)

// @Tags 用户
// @Summary 登录接口
// @Description 登录接口
// @Accept  json
// @Produce  json
// @Param params body user.User true "Login Params"
// @Success 0 {object} libs.R
// @Failure -999 {string} string "服务器出问题"
// @Failure -1002 {string} string "用户不存在"
// @Failure -1003 {string} string "密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		logging.Info("bind user err", err.Error())
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}

	rsp, err := client.Login(&pb.LoginReq{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		logging.Error("Call Login err: ", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -999,
		})
		return
	}
	libs.Resp(libs.R{
		C:    c,
		Code: int(rsp.Rsp.Code),
		Msg:  rsp.Rsp.Msg,
		Data: gin.H{
			"token": rsp.Token,
		},
	})
}
