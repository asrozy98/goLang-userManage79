package main

import (
	"goLang-userManage79/config"
	"goLang-userManage79/model"
	"goLang-userManage79/router"
)

func main() {
	db := config.DatabaseConfig()
	db.AutoMigrate(&model.Users{})

	router.Router(db)
}
