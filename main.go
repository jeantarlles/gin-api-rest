package main

import (
	"github.com/jeantarlles/api-go-gin/database"
	"github.com/jeantarlles/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
