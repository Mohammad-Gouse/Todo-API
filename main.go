package main

import (
	"os"
	"todo-api/controllers"
	"todo-api/routes"
	"todo-api/utils"
)

func main() {
	utils.ConnectDB()
	controllers.InitController(utils.GetDB(), os.Getenv("DB_NAME"))
	router := routes.SetupRouter()
	router.Run(":8080")
}
