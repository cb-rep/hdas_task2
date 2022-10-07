package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Database struct {
	//username: root
	//password: cb123456
	//table-name: hdas_task
	//url: 127.0.0.1:3306
	Username string
	Password string
	Table    string
	Url      string
}

func GetConf() Database {
	//var db Database
	db := Database{}
	file, err := ioutil.ReadFile("./conf/db.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(file, &db)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
