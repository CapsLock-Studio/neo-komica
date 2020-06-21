package model

import (
	"github.com/jinzhu/gorm"
	// Import dialects for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/CapsLock-Studio/neo-komica/config"
	"github.com/golang/glog"
)

// SharedDB - Shared gorm connection instance
var SharedDB *gorm.DB

// InitDB - Initializing SharedDB instance
func InitDB() *gorm.DB {
	c := config.NewGorm()
	db, err := gorm.Open(c.Dialect, c.ConnectURI())

	if err != nil {
		glog.Error("Initilizing DB Failed - ", err)
		panic(err)
	}
	db.LogMode(c.Log)

	SharedDB = db

	return SharedDB
}
