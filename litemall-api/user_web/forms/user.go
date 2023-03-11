package forms

//表单校验

type PassWordLoginForm struct {
	//Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //validator:3 自定义验证器：ValidateMobile
	PassWord  string `form:"password" json:"password" binding:"min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"min=5,max=50"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
