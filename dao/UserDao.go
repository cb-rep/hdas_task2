package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hdas_task/conf"
	"hdas_task/model"
	"log"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User
	GetUser(username string) model.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	conf := conf.GetConf()
	dsn := conf.Username + ":" + conf.Password + "@tcp(" + conf.Url + ")/" + conf.Table + "?charset=utf8mb4&parseTime=true"
	//dsn:="root:cb123456@tcp(127.0.0.1:3306)/hdas_task?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username =?", username).First(&user)
	return user
}

func (mgr *manager) GetUser(username string) model.User {
	var user model.User
	mgr.db.Where("username =?", username).First(&user)
	return user
}
