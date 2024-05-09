package api

import (
	"spc-server/db"

	"github.com/gin-gonic/gin"
)

type loadResult struct {
	Content string `json:"content"`
}

var Save = WithVerifyUser(func(ctx *gin.Context, body string, userName string) {
	flag := db.SaveRecord(&db.Record{UserName: userName, Content: body})
	rtn := &Result{}
	if flag {
		rtn.Code = 0
	} else {
		rtn.Code = -1
		rtn.Msg = "保存失败"
	}
	ctx.JSON(200, rtn)
})

var Load = WithVerifyUser(func(ctx *gin.Context, body string, userName string) {
	record := db.GetRecord(userName)
	content := ""
	if record != nil {
		content = record.Content
	}
	ctx.JSON(200, &Result{
		Code: 0,
		Data: &loadResult{
			Content: content,
		},
	})
})
