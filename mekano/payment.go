package mekano

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/OzkrOssa/mekano-go/database"
	"github.com/OzkrOssa/mekano-go/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func Payment(fileName string, db *gorm.DB) {

	currentTimeInterface := time.Now().Format("02/01/2006 15:04")
	currentTimeMySQL := time.Now().Format("2006-01-02")
	dirPath := "C:/Users/devre/OneDrive/pagos_mekano/"
	rowCount := 0
	var lastConsecutive int = 0
	dataSheet := dataSheet{}

	xlsx, err := excelize.OpenFile(filepath.Join(dirPath, fileName+".xlsx"))
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

		dataSheet.Tipo = append(dataSheet.Tipo, "RC")
		dataSheet.Prefijo = append(dataSheet.Prefijo, "_")
		dataSheet.Numero = append(dataSheet.Numero, currentConsecutive.Consecutive+rowCount) //
		dataSheet.Secuencia = append(dataSheet.Secuencia, "")
		dataSheet.Fecha = append(dataSheet.Fecha, row[4])
		dataSheet.Cuenta = append(dataSheet.Cuenta, "13050501")
		dataSheet.Terceros = append(dataSheet.Terceros, row[1])
		dataSheet.CentroCostos = append(dataSheet.CentroCostos, "C1")
		dataSheet.Nota = append(dataSheet.Nota, "RECAUDO POR VENTA SERVICIOS")
		dataSheet.Debito = append(dataSheet.Debito, "0")
		dataSheet.Credito = append(dataSheet.Credito, row[5])
		dataSheet.Base = append(dataSheet.Base, "0")
		dataSheet.Aplica = append(dataSheet.Aplica, "")
		dataSheet.TipoAnexo = append(dataSheet.TipoAnexo, "")
		dataSheet.PrefijoAnexo = append(dataSheet.PrefijoAnexo, "")
		dataSheet.NumeroAnexo = append(dataSheet.NumeroAnexo, "")
		dataSheet.Usuario = append(dataSheet.Usuario, "SUPERVISOR")
		dataSheet.Signo = append(dataSheet.Signo, "")
		dataSheet.CuentaCobrar = append(dataSheet.CuentaCobrar, "")
		dataSheet.CuentaPagar = append(dataSheet.CuentaPagar, "")
		dataSheet.NombreTercero = append(dataSheet.NombreTercero, row[2])
		dataSheet.NombreCentro = append(dataSheet.NombreCentro, "CENTRO DE COSTOS GENERAL")
		dataSheet.Interface = append(dataSheet.Interface, currentTimeInterface)

		dataSheet.Tipo = append(dataSheet.Tipo, "RC")
		dataSheet.Prefijo = append(dataSheet.Prefijo, "_")
		dataSheet.Numero = append(dataSheet.Numero, currentConsecutive.Consecutive+rowCount)
		dataSheet.Secuencia = append(dataSheet.Secuencia, "")
		dataSheet.Fecha = append(dataSheet.Fecha, row[4])
		dataSheet.Cuenta = append(dataSheet.Cuenta, utils.Caja[row[9]])
		dataSheet.Terceros = append(dataSheet.Terceros, row[1])
		dataSheet.CentroCostos = append(dataSheet.CentroCostos, "C1")
		dataSheet.Nota = append(dataSheet.Nota, "RECAUDO POR VENTA SERVICIOS")
		dataSheet.Debito = append(dataSheet.Debito, row[5])
		dataSheet.Credito = append(dataSheet.Credito, "0")
		dataSheet.Base = append(dataSheet.Base, "0")
		dataSheet.Aplica = append(dataSheet.Aplica, "")
		dataSheet.TipoAnexo = append(dataSheet.TipoAnexo, "")
		dataSheet.PrefijoAnexo = append(dataSheet.PrefijoAnexo, "")
		dataSheet.NumeroAnexo = append(dataSheet.NumeroAnexo, "")
		dataSheet.Usuario = append(dataSheet.Usuario, "SUPERVISOR")
		dataSheet.Signo = append(dataSheet.Signo, "")
		dataSheet.CuentaCobrar = append(dataSheet.CuentaCobrar, "")
		dataSheet.CuentaPagar = append(dataSheet.CuentaPagar, "")
		dataSheet.NombreTercero = append(dataSheet.NombreTercero, row[2])
		dataSheet.NombreCentro = append(dataSheet.NombreCentro, "CENTRO DE COSTOS GENERAL")
		dataSheet.Interface = append(dataSheet.Interface, currentTimeInterface)

		lastConsecutive = currentConsecutive.Consecutive + rowCount
	}

	db.Create(&database.MekanoPayments{Consecutive: lastConsecutive, CreateAt: currentTimeMySQL})

	txtFile, err := os.Create("C:/APOLOSOFT/MEKANO_REMOTO/INTERFACES/CONTABLE.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer txtFile.Close()

	writer := csv.NewWriter(txtFile)
	writer.Comma = ','

	for i := range dataSheet.Tipo {
		row := []string{
			dataSheet.Tipo[i],
			dataSheet.Prefijo[i],
			strconv.Itoa(dataSheet.Numero[i]),
			dataSheet.Secuencia[i],
			dataSheet.Fecha[i],
			dataSheet.Cuenta[i],
			dataSheet.Terceros[i],
			dataSheet.CentroCostos[i],
			dataSheet.Nota[i],
			dataSheet.Debito[i],
			dataSheet.Credito[i],
			dataSheet.Base[i],
			dataSheet.Aplica[i],
			dataSheet.TipoAnexo[i],
			dataSheet.PrefijoAnexo[i],
			dataSheet.NumeroAnexo[i],
			dataSheet.Usuario[i],
			dataSheet.Signo[i],
			dataSheet.CuentaCobrar[i],
			dataSheet.CuentaPagar[i],
			dataSheet.NombreTercero[i],
			dataSheet.NombreCentro[i],
			dataSheet.Interface[i],
		}
		writer.Write(row)
	}
	writer.Flush()
}
