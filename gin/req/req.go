package req

// 登录
type LoginForm struct {
	Username string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
