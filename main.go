package main

import (
	"todo_list/routes"
)

func main() {
	router := routes.SetupRoutes()
	router.Run(":8080")
}
