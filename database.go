package identity

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/kthomas/go-db-config"
)

func DatabaseConnection() *gorm.DB {
	return dbconf.DatabaseConnection()
}

func MigrateSchema() {
	db := dbconf.DatabaseConnection()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Token{})
}

func ResetData() {
	db := dbconf.DatabaseConnection()
	db.Delete(&User{})
	db.Delete(&Token{})
}
