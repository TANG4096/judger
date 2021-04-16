package db

import (
	"fmt"
	"judger/util"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	dbConfigMap := util.GetDbConfigMap("test")
	log.Println(dbConfigMap)
	user := dbConfigMap["userName"]
	passworld := dbConfigMap["passworld"]
	ip := dbConfigMap["ip"]
	port := dbConfigMap["port"]
	arg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, passworld, ip, port, "test")
	log.Println(arg)
	var err error
	db, err = gorm.Open("mysql", arg)
	//log.Printf("%v", db)
	if err != nil {
		log.Fatalf("%s", err.Error())
		panic(err)
	}
}

func GetDB() *gorm.DB {
	err := db.DB().Ping()
	if err != nil {
		log.Panicln(err.Error())
		db.Close()
	}
	return db
}
