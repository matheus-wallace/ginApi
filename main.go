package main

import (
	"ginApi/database"
	"ginApi/routes"
)

func main() {
	database.ConectaComOBancoDeDados()

	routes.HandleRequests()
}
