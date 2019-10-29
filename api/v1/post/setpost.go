package post

import (
	client "catdogs-be/client/post"
	"catdogs-be/libs"
	"catdogs-be/logging"
	pb "catdogs-be/pb"

	"github.com/gin-gonic/gin"
)

// @Tags 文章
// @Summary 发布文章
// @Description 发布文章
// @Accept  json
// @Produce  json
// @Param params body post.SetPost true "SetPost Params"
// @Success 0 {object} libs.R
// @Failure -999 {string} string "服务器出问题"
// @Failure -3000 {string} string "参数错误"
// @Router /setpost [post]
func SetPostHandler(c *gin.Context) {
	var setpost SetPost
	if err := c.ShouldBind(&setpost); err != nil {
		logging.Error("bind setpost err: ", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}
	rsp, err := client.SetPost(&pb.SetPostReq{
		Title:   setpost.Title,
		Content: setpost.Content,
		Author:  setpost.Author,
	})
	if err != nil {
		logging.Error("Call SetPost err: ", err)
	}
	libs.Resp(libs.R{
		C:    c,
		Code: int(rsp.Rsp.Code),
		Msg:  rsp.Rsp.Msg,
	})
}
