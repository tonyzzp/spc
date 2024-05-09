package api

import (
	"io"
	"net/http"
	"spc-server/sign"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var _tokens = &sync.Map{}

func setToken(user string, token string) {
	_tokens.Store(user, token)
}

func getToken(user string) string {
	token, ok := _tokens.Load(user)
	if ok {
		return token.(string)
	}
	return ""
}

type HandlerFunc func(ctx *gin.Context, body string)

type HandlerFuncUser func(ctx *gin.Context, body string, userName string)

func WithVerifySign(key string, handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bs, e := io.ReadAll(ctx.Request.Body)
		if e != nil {
			ctx.Abort()
			return
		}
		s := string(bs)
		flag := sign.Verify(sign.VerifyParam{
			Body: s,
			Key:  key,
			Sign: ctx.GetHeader("sign"),
		})
		if !flag {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		handler(ctx, s)
	}
}

func WithVerifyUser(handler HandlerFuncUser) HandlerFunc {
	return func(ctx *gin.Context, body string) {
		_token := ctx.GetHeader("token")
		strs := strings.Split(_token, ".")
		userName := strs[0]
		token := strs[1]
		if getToken(userName) != token {
			ctx.JSON(200, &Result{
				Code: -1,
				Msg:  "token不正确，请重新登录",
			})
			return
		}
		handler(ctx, body, userName)
	}
}
