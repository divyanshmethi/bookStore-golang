package main

import (
	"bookStore/config"
	"bookStore/controllers"
	"bookStore/db"
	"bookStore/handlers"
)

func main() {
	config.InitGlobalConfig()

	//init our DB
	database := db.NewPostgresDBMSImpl()

	//init our controller
	controller := controllers.NewBookController(database)

	// init handlers
	service := handlers.NewService(controller)
	service.RegisterRoutesAndStart()
}
