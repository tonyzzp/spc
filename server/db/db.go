package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	UserName string `gorm:"primaryKey"`
	Password string
}

type Record struct {
	UserName string `gorm:"primaryKey"`
	Content  string
}

var db *gorm.DB

func Init(file string) {
	var err error
	db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("打开数据库失败:" + file)
	}
	db.AutoMigrate(&User{}, &Record{})
}

func GetUser(userName string) *User {
	rtn := &User{}
	result := db.First(rtn, "user_name=?", userName)
	if result.Error != nil {
		return nil
	}
	return rtn
}

func SaveUser(user *User) bool {
	rtn := db.Save(user)
	return rtn.Error == nil
}

func AddUser(user *User) bool {
	result := db.Create(user)
	return result.Error == nil
}

func GetRecord(user string) *Record {
	rtn := &Record{UserName: user}
	db.First(rtn)
	return rtn
}

func SaveRecord(record *Record) bool {
	result := db.Save(record)
	log.Println("db.Save", record.UserName, record.Content)
	log.Println("", result.Error, result.RowsAffected)
	return result.Error == nil
}
