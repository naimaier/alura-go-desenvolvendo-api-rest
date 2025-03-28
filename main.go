package main

import (
	"api/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor REST")
	routes.HandleRequests()
}
