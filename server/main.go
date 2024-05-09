package main

import (
	"spc-server/api"
	"spc-server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const SIGN_KEY = "15998569"

func main() {
	db.Init("spc-db.sqlite")
	g := gin.New()
	g.Use(cors.Default())
	g.Use(gin.Logger())
	g.POST("/api/reg", api.WithVerifySign(SIGN_KEY, api.Reg))
	g.POST("/api/login", api.WithVerifySign(SIGN_KEY, api.Login))
	g.POST("/api/changePassword", api.WithVerifySign(SIGN_KEY, api.ChangePassword))
	g.POST("/api/save", api.WithVerifySign(SIGN_KEY, api.Save))
	g.POST("/api/load", api.WithVerifySign(SIGN_KEY, api.Load))
	g.Run(":8080")
}
