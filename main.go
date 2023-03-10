package main

import (
	"github.com/OzkrOssa/mekano-go/database"
	"github.com/OzkrOssa/mekano-go/mekano"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		panic(err)
	}
	mekano.Payment("PAGOS 18-21 FEBRERO 2023 PARA MEKANO", db)
}
