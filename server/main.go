package main

import (
	"log"
	"os"
	"path/filepath"
	"spc-server/api"
	"spc-server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const SIGN_KEY = "15998569"

func findStaticDir() string {
	exec, _ := os.Executable()
	log.Println("exec", exec)

	dir := filepath.Join(exec, "static/index.html")
	if fi, _ := os.Stat(dir); fi != nil {
		dir = filepath.Dir(dir)
		dir, _ = filepath.Abs(dir)
		return dir
	}

	dir = filepath.Join(exec, "../client/dist/index.html")
	if fi, _ := os.Stat(dir); fi != nil {
		dir = filepath.Dir(dir)
		dir, _ = filepath.Abs(dir)
		return dir
	}

	dir = "static/index.html"
	if fi, _ := os.Stat(dir); fi != nil {
		dir, _ = filepath.Abs("static")
		return dir
	}

	dir = "../client/dist/index.html"
	if fi, _ := os.Stat(dir); fi != nil {
		dir = filepath.Dir(dir)
		dir, _ = filepath.Abs(dir)
		return dir
	}

	return ""
}

func main() {
	db.Init("spc-db.sqlite")
	g := gin.New()
	g.Use(cors.Default())
	g.Use(gin.Logger())
	staticDir := findStaticDir()
	log.Println("staticDir", staticDir)
	if staticDir != "" {
		g.Static("/", staticDir)
	}
	g.POST("/api/reg", api.WithVerifySign(SIGN_KEY, api.Reg))
	g.POST("/api/login", api.WithVerifySign(SIGN_KEY, api.Login))
	g.POST("/api/changePassword", api.WithVerifySign(SIGN_KEY, api.ChangePassword))
	g.POST("/api/save", api.WithVerifySign(SIGN_KEY, api.Save))
	g.POST("/api/load", api.WithVerifySign(SIGN_KEY, api.Load))
	g.Run(":8080")
}
