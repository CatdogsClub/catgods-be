package user

import (
	"crypto/md5"
	"fmt"
	"strconv"

	configs "catdogs.club/back-end/configs/common"
	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/logging"
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err)
		logging.Info("bind user err", err.Error())
	}

	has := verifyUser(&user)
	if has {
		libs.Resp(libs.R{
			C:    c,
			Code: -1000,
		})
		return
	}
	code := libs.RandString(16)
	url := configs.ActivateUrl + "?v=" + code
	cont := fmt.Sprintf("点击链接激活您的账号: %s", url)
	go saveUser(&user)
	go saveCode(code, &user)
	go libs.SendMail(user.Email, "邮箱激活验证", cont)
	libs.Resp(libs.R{
		C:    c,
		Code: 0,
	})
}

func saveCode(code string, param *User) {
	v := models.VerifyCode{
		Email: param.Email,
		Code:  code,
	}
	err := v.Set()
	if err != nil {
		fmt.Println(err)
	}
}

// 保存用户到数据库
func saveUser(param *User) {
	pwData := md5.Sum([]byte(param.Password))
	pwS := fmt.Sprintf("%x", pwData)
	session := models.Db.NewSession()
	defer session.Close()
	err := session.Begin()
	u := models.User{
		Email:    param.Email,
		Password: pwS,
	}
	err = u.Set()
	if err != nil {
		fmt.Println(err)
		logging.Error("save user error: ", err)
		session.Rollback()
		return
	}
	logging.Info("insert id: ", u.Id, "email: ", u.Email)
	// 生成openid
	strId := strconv.Itoa(u.Id)
	openid := libs.AESEncrypt(strId)
	newu := models.User{Openid: openid}
	_, err = session.Id(u.Id).Update(newu)
	if err != nil {
		logging.Error("update openid error: ", err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		logging.Error("save user commit error: ", err)
		return
	}
}

// 验证用户是否存在
func verifyUser(param *User) bool {
	u := models.User{
		Email: param.Email,
	}
	has, err := u.Get()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return has
}
