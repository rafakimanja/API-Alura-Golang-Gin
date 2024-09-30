package main

import (
	"API/GIN/database"
	"API/GIN/routes"
)


func main() {
	database.ConectaComDB()
	routes.HandleRequest()
}

