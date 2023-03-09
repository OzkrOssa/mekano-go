package main

import (
	"fmt"

	"github.com/OzkrOssa/mekano-go/database"
)

func main() {
	//mekano.Payment("PAGOS 18-21 FEBRERO 2023 PARA MEKANO")
	db, err := database.Connection()
	if err != nil {
		panic(err)
	}

	var mekanoPayment database.MekanoPayments
	db.First(&mekanoPayment, 31)

	fmt.Println(mekanoPayment)
}
