package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"spc-server/api"
	"spc-server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

const SIGN_KEY = "15998569"

var isDocker = false
var loggerWriter io.Writer

func init() {
	envs := os.Environ()
	log.Println(envs)
	isDocker = slices.Contains(envs, "ISDOCKER=true")
}

func initLogger() {
	_l := &lumberjack.Logger{
		MaxSize:    10,
		MaxAge:     60,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	if isDocker {
		_l.Filename = "/data/logs/spc.log"
	} else {
		_l.Filename = "spc.log"
	}
	loggerWriter = io.MultiWriter(_l, os.Stdout)
	log.SetOutput(loggerWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func initDB() {
	dbPath := "spc-db.sqlite"
	if isDocker {
		dbPath = "/data/spc-db.sqlite"
	}
	log.Println("dbpath", dbPath)
	db.Init(dbPath)
}

func findStaticDir() string {
	exec, _ := os.Executable()
	log.Println("exec", exec)

	dir := filepath.Join(exec, "../static/index.html")
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
	initLogger()
	initDB()

	g := gin.New()
	g.Use(cors.Default())
	g.Use(gin.LoggerWithWriter(loggerWriter))
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
	g.Run(":http")
}
