package main

import (
	"api/model"
	"api/routes"
	"fmt"
)

func main() {
	model.Personalidades = []model.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
	}
	fmt.Println("Iniciando o servidor REST")
	routes.HandleRequests()
}
