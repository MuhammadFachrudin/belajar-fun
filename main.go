package main

import (
	"golang-mysql/config"
	"golang-mysql/controller"
	"golang-mysql/migrations"

	"github.com/gin-gonic/gin"
)

func main() {

	migrations.Migration()

	db := config.DB()
	idb := &controller.InDB{DB: db}

	r := gin.Default()

	r.POST("/person", idb.AddUser)
	r.GET("/person", idb.GetAllUser)
	r.GET("/person/:id", idb.GetUser)
	r.PUT("/person", idb.UpdateUser)
	r.DELETE("/person/:id", idb.DeleteUser)
	r.Run()
}
