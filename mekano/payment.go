package mekano

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/OzkrOssa/mekano-go/database"
	"github.com/OzkrOssa/mekano-go/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func Payment(fileName string, db *gorm.DB) {
	var paymentDataSlice []MekanoDataSheet
	var rowCount, consecutive int = 0, 0

	xlsx, err := excelize.OpenFile(filepath.Join(utils.PaymentFileDirPath, fileName+".xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}

	excelRows, err := xlsx.GetRows("hoja1")
	if err != nil {
		fmt.Println(err)
	}

	var currentConsecutive database.MekanoPayments
	db.Table("mekanopayments").Order("consecutive DESC").Limit(1).First(&currentConsecutive)

	for _, row := range excelRows[1:] {
		rowCount++
		consecutive = currentConsecutive.Consecutive + rowCount

		paymentData := MekanoDataSheet{
			Tipo:          "RC",
			Prefijo:       "_",
			Numero:        strconv.Itoa(consecutive),
			Secuencia:     "",
			Fecha:         row[4],
			Cuenta:        "13050501",
			Terceros:      row[1],
			CentroCostos:  "C1",
			Nota:          "RECAUDO POR VENTA SERVICIOS",
			Debito:        "0",
			Credito:       row[5],
			Base:          "0",
			Aplica:        "",
			TipoAnexo:     "",
			PrefijoAnexo:  "",
			NumeroAnexo:   "",
			Usuario:       "SUPERVISOR",
			Signo:         "",
			CuentaCobrar:  "",
			CuentaPagar:   "",
			NombreTercero: row[2],
			NombreCentro:  "CENTRO DE COSTOS GENERAL",
			Interface:     utils.CurrentTimeMekanoInterface,
		}
		paymentDataSlice = append(paymentDataSlice, paymentData)

		paymentData2 := MekanoDataSheet{
			Tipo:          "RC",
			Prefijo:       "_",
			Numero:        strconv.Itoa(consecutive),
			Secuencia:     "",
			Fecha:         row[4],
			Cuenta:        utils.Caja[row[9]],
			Terceros:      row[1],
			CentroCostos:  "C1",
			Nota:          "RECAUDO POR VENTA SERVICIOS",
			Debito:        row[5],
			Credito:       "0",
			Base:          "0",
			Aplica:        "",
			TipoAnexo:     "",
			PrefijoAnexo:  "",
			NumeroAnexo:   "",
			Usuario:       "SUPERVISOR",
			Signo:         "",
			CuentaCobrar:  "",
			CuentaPagar:   "",
			NombreTercero: row[2],
			NombreCentro:  "CENTRO DE COSTOS GENERAL",
			Interface:     utils.CurrentTimeMekanoInterface,
		}
		paymentDataSlice = append(paymentDataSlice, paymentData2)
	}

	//Save in database last consecutive generated by iter excel rows
	db.Create(&database.MekanoPayments{Consecutive: consecutive, CreateAt: utils.CurrentTimeToMySQL})

	txtFile, err := os.Create(filepath.Join(utils.MekanoInterfaceDirPath, "CONTABLE.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer txtFile.Close()

	writer := csv.NewWriter(txtFile)
	writer.Comma = ','

	for _, data := range paymentDataSlice {
		row := []string{
			data.Tipo,
			data.Prefijo,
			data.Numero,
			data.Secuencia,
			data.Fecha,
			data.Cuenta,
			data.Terceros,
			data.CentroCostos,
			data.Nota,
			data.Debito,
			data.Credito,
			data.Base,
			data.Aplica,
			data.TipoAnexo,
			data.PrefijoAnexo,
			data.NumeroAnexo,
			data.Usuario,
			data.Signo,
			data.CuentaCobrar,
			data.CuentaPagar,
			data.NombreTercero,
			data.NombreCentro,
			data.Interface,
		}
		writer.Write(row)
	}
	writer.Flush()
}
