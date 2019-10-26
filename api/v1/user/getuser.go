package user

import (
	"catdogs-be/libs"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	libs.Resp(libs.R{
		C:    c,
		Msg:  "success",
		Code: 0,
	})
}
