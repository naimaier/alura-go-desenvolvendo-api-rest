package main

import (
	"api/model"
	"api/routes"
	"fmt"
)

func main() {
	model.Personalidades = []model.Personalidade{
		{Nome: "Nome1", Historia: "Historia1"},
		{Nome: "Nome2", Historia: "Historia2"},
	}
	fmt.Println("Iniciando o servidor REST")
	routes.HandleRequests()
}
