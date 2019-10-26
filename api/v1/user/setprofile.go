package user

import (
	client "catdogs-be/client/user"
	"catdogs-be/libs"
	"catdogs-be/logging"
	pb "catdogs-be/pb"

	"github.com/gin-gonic/gin"
)

func SetProfile(c *gin.Context) {
	var profile ProfileParam
	if err := c.ShouldBind(&profile); err != nil {
		logging.Error("Bind profile err", err.Error())
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}

	rsp, err := client.SetProfile(&pb.SetProfileReq{
		Profile: convertToProfile(profile),
	})
	if err != nil {
		logging.Error("Call SetProfile err: ", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -1004,
		})
		return
	}
	libs.Resp(libs.R{
		C:    c,
		Code: int(rsp.Rsp.Code),
		Msg:  rsp.Rsp.Msg,
	})
}

func convertToProfile(profile ProfileParam) *pb.Profile {
	pro := profile.Profile
	return &pb.Profile{
		Name:     pro.Name,
		Age:      pro.Age,
		Gender:   pro.Gender,
		PhoneNum: pro.PhoneNum,
		Email:    pro.Email,
		Birthday: pro.Birthday,
		City:     pro.City,
		Address:  pro.Address,
	}
}
