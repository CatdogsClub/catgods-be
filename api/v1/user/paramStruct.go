package user

type User struct {
	Email    string `json:"email" binding:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ProfileParam struct {
	Profile *Profile `form:"profile" json:"profile" validate:"required"`
}

type Profile struct {
	Name     string `form:"name" json:"name"`
	Gender   string `form:"gender" json:"gender"`
	Age      int32  `form:"age" json:"age"`
	PhoneNum string `form:"phonenum" json:"phonenum"`
	Email    string `form:"email" json:"email"`
	Birthday string `form:"birthday" json:"birthday"`
	City     string `form:"city" json:"city"`
	Address  string `form:"address" json:"address"`
}
