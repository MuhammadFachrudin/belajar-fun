package migrations

import (
	"golang-mysql/config"
	"golang-mysql/table"
)

func Migration() {
	db := config.DB()
	db.AutoMigrate(table.User{}, table.Produk{})
}
