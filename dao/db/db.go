package db

import (
	"fmt"

	"github.com/changecoder/webservice/common/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Model Model
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup Connect to MYSQL
func Setup() {
	var err error

	dbName := setting.DatabaseSetting.Name
	dbUser := setting.DatabaseSetting.User
	dbPassword := setting.DatabaseSetting.Password
	dbHost := setting.DatabaseSetting.Host
	dbType := setting.DatabaseSetting.Type

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName))

	if err != nil {
		panic(err.Error())
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
}

// Close Close
func Close() {
	defer db.Close()
}
