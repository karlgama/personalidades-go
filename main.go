package main

import (
	"fmt"
	"personalidades/database"
	"personalidades/models"
	"personalidades/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
