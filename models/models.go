package models

import (
	"fmt"
	"log"

	"github.com/LiuHuanQing/go_server_base/pkg/setting"
	"github.com/jinzhu/gorm"

	//postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//GetEngine ...
func GetEngine() *gorm.DB {
	return db
}

//Setup ...
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Password))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
}
