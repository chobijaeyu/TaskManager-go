package models

import (
	"CardTaskGo/conf"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var (
		err                                                     error
		dbType, dbName, user, password, host, tablePrefix, port string
	)
	dbType = conf.DB_TYPE
	dbName = conf.DB_NAME
	user = conf.DB_USER
	password = conf.DB_PASSWORD
	host = conf.DB_HOST
	port = conf.DB_PORT
	tablePrefix = conf.DB_TABLE_PREFIX

	// db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbName,
		password,
	))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableNameHandler string) string {
		return tablePrefix + defaultTableNameHandler
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.Model(&Taskitem{}).AddForeignKey("goods_id", "Bitsu_Goods(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Task{}, &Taskitem{}, &Project{})
}

func CLoseDB() {
	defer db.Close()
}
