package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	entity2 "interview/pkg/core/entity"
	"os"
)

func GetDatabase(memory bool, dsn string) *gorm.DB {

	// MySQL connection string

	var db *gorm.DB
	var err error
	if memory {
		// Update the username, password, host, port, and database name accordingly
		fullDirectory := os.TempDir() + "/" + "ice_db" + ".db"
		db, err = gorm.Open(sqlite.Open(fullDirectory), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		// Open the connection to the database
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 gormlogger.Default.LogMode(gormlogger.Silent),
		})
		if err != nil {
			panic("failed to connect database")
		}

	}

	err = db.AutoMigrate(&entity2.Cart{}, &entity2.CartItem{})
	if err != nil {
		panic(err)
	}

	return db
}
