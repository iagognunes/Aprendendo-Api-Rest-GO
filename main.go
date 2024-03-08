package main

import (
	"api-rest-go/models"
	"api-rest-go/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go.")
	routes.HandleRequest()
}
