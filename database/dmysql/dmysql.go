//@program: work
//@description: 
//@author: edte
//@create: 2020-05-30 18:08
package dmysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var G_db = &gorm.DB{}

type Game struct {
	gorm.Model
	CnName string
	JnName string
	Info   string
	Vote   int
}

func Start() {
	db, err := gorm.Open("mysql", "root:mima@/game?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	if db.HasTable(&Game{}) {
		db.AutoMigrate(&Game{})
	} else {
		db.CreateTable(&Game{})
	}
	G_db = db
}
