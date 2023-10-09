package main

import (
	"fmt"
	"personalidades/models"
	"personalidades/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Algum nome", Historia: "historia"},
		{Id: 2, Nome: "Algum nome 1", Historia: "historia 1"},
	}

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
