package api

import (
	"encoding/json"
	"spc-server/db"
	"spc-server/tools"

	"github.com/gin-gonic/gin"
)

type regParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type changePasswordParam struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type loginResult struct {
	Token string `json:"token"`
}

func Reg(ctx *gin.Context, body string) {
	param := &regParam{}
	json.Unmarshal([]byte(body), param)
	exist := db.GetUser(param.UserName)
	if exist != nil {
		ctx.JSON(200, &Result{
			Code: -1,
			Msg:  "用户名已存在",
		})
		return
	}
	success := db.AddUser(&db.User{
		UserName: param.UserName,
		Password: param.Password,
	})
	result := &Result{}
	if success {
		result.Code = 0
	} else {
		result.Code = -1
		result.Msg = "注册失败"
	}
	ctx.JSON(200, result)
}

func Login(ctx *gin.Context, body string) {
	param := &regParam{}
	json.Unmarshal([]byte(body), param)
	exist := db.GetUser(param.UserName)
	if exist == nil || exist.Password != param.Password {
		ctx.JSON(200, &Result{
			Code: -1,
			Msg:  "用户名密码不正确",
		})
		return
	}
	token := tools.RandomString(32)
	setToken(param.UserName, token)
	ctx.JSON(200, &Result{
		Code: 0,
		Data: &loginResult{
			Token: param.UserName + "." + token,
		},
	})
}

var ChangePassword = WithVerifyUser(func(ctx *gin.Context, body, userName string) {
	param := &changePasswordParam{}
	json.Unmarshal([]byte(body), param)
	user := db.GetUser(userName)
	if user == nil {
		ctx.JSON(200, &Result{
			Code: -1,
			Msg:  "账号不存在",
		})
		return
	}
	if user.Password != param.OldPassword {
		ctx.JSON(200, &Result{
			Code: -1,
			Msg:  "原密码不正确",
		})
		return
	}
	user.Password = param.NewPassword
	flag := db.SaveUser(user)
	rtn := &Result{}
	if flag {
		rtn.Code = 0
	} else {
		rtn.Code = -1
		rtn.Msg = "修改失败"
	}
	ctx.JSON(200, rtn)
})
