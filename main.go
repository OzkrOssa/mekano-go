package main

import (
	"fmt"
	"time"

	"github.com/OzkrOssa/mekano-go/database"
	"github.com/OzkrOssa/mekano-go/mekano"
)

func main() {
	start := time.Now()
	db, err := database.Connection()
	if err != nil {
		panic(err)
	}
	//mekano.Payment("PAGOS 18-21 FEBRERO 2023 PARA MEKANO", db)
	mekano.Billing("REPORTE FACTURACION ELECTRONICA FEBRERO 2023", db)
	since := time.Since(start)

	fmt.Println(since)
}
