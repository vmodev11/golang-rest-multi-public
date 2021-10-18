package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func InitFromSQLLite(db_path_str string) (func(), error) {

	db, closeFunc, err := NewGormDB(db_path_str)
	if err != nil {
		panic(err)
	}
	err = migrateTable(db)
	migrateData()
	return closeFunc, err
}

func migrateTable(db *gorm.DB) error {
	return db.AutoMigrate(
		new(User),
	).Error
}
func NewGormDB(db_path_str string) (*gorm.DB, func(), error) {
	db, err = gorm.Open("sqlite3", db_path_str)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	cleanFunc := func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Gorm db close error: %s", err)
		}
	}
	return db, cleanFunc, err
}

func migrateData() {
	(&User{}).MigrateData()
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
func TransactionDB(db *gorm.DB, txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit().Error // err is nil; if Commit returns error update err
		}
	}()
	err = txFunc(tx)
	return err
}
