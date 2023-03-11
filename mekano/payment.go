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
	var rowCount, lastConsecutive int = 0, 0
	paymentDatasheet := dataSheet{}

	xlsx, err := excelize.OpenFile(filepath.Join(utils.LocalCloudDirPath, fileName+".xlsx"))
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

		paymentDatasheet.Tipo = append(paymentDatasheet.Tipo, "RC")
		paymentDatasheet.Prefijo = append(paymentDatasheet.Prefijo, "_")
		paymentDatasheet.Numero = append(paymentDatasheet.Numero, currentConsecutive.Consecutive+rowCount) //
		paymentDatasheet.Secuencia = append(paymentDatasheet.Secuencia, "")
		paymentDatasheet.Fecha = append(paymentDatasheet.Fecha, row[4])
		paymentDatasheet.Cuenta = append(paymentDatasheet.Cuenta, "13050501")
		paymentDatasheet.Terceros = append(paymentDatasheet.Terceros, row[1])
		paymentDatasheet.CentroCostos = append(paymentDatasheet.CentroCostos, "C1")
		paymentDatasheet.Nota = append(paymentDatasheet.Nota, "RECAUDO POR VENTA SERVICIOS")
		paymentDatasheet.Debito = append(paymentDatasheet.Debito, "0")
		paymentDatasheet.Credito = append(paymentDatasheet.Credito, row[5])
		paymentDatasheet.Base = append(paymentDatasheet.Base, "0")
		paymentDatasheet.Aplica = append(paymentDatasheet.Aplica, "")
		paymentDatasheet.TipoAnexo = append(paymentDatasheet.TipoAnexo, "")
		paymentDatasheet.PrefijoAnexo = append(paymentDatasheet.PrefijoAnexo, "")
		paymentDatasheet.NumeroAnexo = append(paymentDatasheet.NumeroAnexo, "")
		paymentDatasheet.Usuario = append(paymentDatasheet.Usuario, "SUPERVISOR")
		paymentDatasheet.Signo = append(paymentDatasheet.Signo, "")
		paymentDatasheet.CuentaCobrar = append(paymentDatasheet.CuentaCobrar, "")
		paymentDatasheet.CuentaPagar = append(paymentDatasheet.CuentaPagar, "")
		paymentDatasheet.NombreTercero = append(paymentDatasheet.NombreTercero, row[2])
		paymentDatasheet.NombreCentro = append(paymentDatasheet.NombreCentro, "CENTRO DE COSTOS GENERAL")
		paymentDatasheet.Interface = append(paymentDatasheet.Interface, utils.CurrentTimeMekanoInterface)

		paymentDatasheet.Tipo = append(paymentDatasheet.Tipo, "RC")
		paymentDatasheet.Prefijo = append(paymentDatasheet.Prefijo, "_")
		paymentDatasheet.Numero = append(paymentDatasheet.Numero, currentConsecutive.Consecutive+rowCount)
		paymentDatasheet.Secuencia = append(paymentDatasheet.Secuencia, "")
		paymentDatasheet.Fecha = append(paymentDatasheet.Fecha, row[4])
		paymentDatasheet.Cuenta = append(paymentDatasheet.Cuenta, utils.Caja[row[9]])
		paymentDatasheet.Terceros = append(paymentDatasheet.Terceros, row[1])
		paymentDatasheet.CentroCostos = append(paymentDatasheet.CentroCostos, "C1")
		paymentDatasheet.Nota = append(paymentDatasheet.Nota, "RECAUDO POR VENTA SERVICIOS")
		paymentDatasheet.Debito = append(paymentDatasheet.Debito, row[5])
		paymentDatasheet.Credito = append(paymentDatasheet.Credito, "0")
		paymentDatasheet.Base = append(paymentDatasheet.Base, "0")
		paymentDatasheet.Aplica = append(paymentDatasheet.Aplica, "")
		paymentDatasheet.TipoAnexo = append(paymentDatasheet.TipoAnexo, "")
		paymentDatasheet.PrefijoAnexo = append(paymentDatasheet.PrefijoAnexo, "")
		paymentDatasheet.NumeroAnexo = append(paymentDatasheet.NumeroAnexo, "")
		paymentDatasheet.Usuario = append(paymentDatasheet.Usuario, "SUPERVISOR")
		paymentDatasheet.Signo = append(paymentDatasheet.Signo, "")
		paymentDatasheet.CuentaCobrar = append(paymentDatasheet.CuentaCobrar, "")
		paymentDatasheet.CuentaPagar = append(paymentDatasheet.CuentaPagar, "")
		paymentDatasheet.NombreTercero = append(paymentDatasheet.NombreTercero, row[2])
		paymentDatasheet.NombreCentro = append(paymentDatasheet.NombreCentro, "CENTRO DE COSTOS GENERAL")
		paymentDatasheet.Interface = append(paymentDatasheet.Interface, utils.CurrentTimeMekanoInterface)

		lastConsecutive = currentConsecutive.Consecutive + rowCount
	}

	db.Create(&database.MekanoPayments{Consecutive: lastConsecutive, CreateAt: utils.CurrentTimeToMySQL})

	txtFile, err := os.Create(filepath.Join(utils.MekanoInterfaceDirPath, "CONTABLE.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer txtFile.Close()

	writer := csv.NewWriter(txtFile)
	writer.Comma = ','

	for i := range paymentDatasheet.Tipo {
		row := []string{
			paymentDatasheet.Tipo[i],
			paymentDatasheet.Prefijo[i],
			strconv.Itoa(paymentDatasheet.Numero[i]),
			paymentDatasheet.Secuencia[i],
			paymentDatasheet.Fecha[i],
			paymentDatasheet.Cuenta[i],
			paymentDatasheet.Terceros[i],
			paymentDatasheet.CentroCostos[i],
			paymentDatasheet.Nota[i],
			paymentDatasheet.Debito[i],
			paymentDatasheet.Credito[i],
			paymentDatasheet.Base[i],
			paymentDatasheet.Aplica[i],
			paymentDatasheet.TipoAnexo[i],
			paymentDatasheet.PrefijoAnexo[i],
			paymentDatasheet.NumeroAnexo[i],
			paymentDatasheet.Usuario[i],
			paymentDatasheet.Signo[i],
			paymentDatasheet.CuentaCobrar[i],
			paymentDatasheet.CuentaPagar[i],
			paymentDatasheet.NombreTercero[i],
			paymentDatasheet.NombreCentro[i],
			paymentDatasheet.Interface[i],
		}
		writer.Write(row)
	}
	writer.Flush()
}
