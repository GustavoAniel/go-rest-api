package main

import (
	"fmt"

	"github.com/gustavoaniel/go-api-rest/database"
	"github.com/gustavoaniel/go-api-rest/models"
	"github.com/gustavoaniel/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest em: http://localhost:8000/")
	routes.HandleResquest()
}
