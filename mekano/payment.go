package mekano

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func Payment(fileName string) {
	currentTime := time.Now().Format("02/01/2006 15:04")
	dirPath := "C:/Users/devre/OneDrive/pagos_mekano/"

	dataSheet := DataSheet{}

	xlsx, err := excelize.OpenFile(filepath.Join(dirPath, fileName+".xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}

	excelRows, err := xlsx.GetRows("hoja1")
	if err != nil {
		fmt.Println(err)
	}

	rowCount := 0
	//TODO: query to db
	currentConsecutive := 1

	for _, row := range excelRows[1:] {
		rowCount++

		dataSheet.Tipo = append(dataSheet.Tipo, "RC")
		dataSheet.Prefijo = append(dataSheet.Prefijo, "_")
		dataSheet.Numero = append(dataSheet.Numero, currentConsecutive+rowCount)
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
		dataSheet.Interface = append(dataSheet.Interface, currentTime)

		dataSheet.Tipo = append(dataSheet.Tipo, "RC")
		dataSheet.Prefijo = append(dataSheet.Prefijo, "_")
		dataSheet.Numero = append(dataSheet.Numero, currentConsecutive+rowCount)
		dataSheet.Secuencia = append(dataSheet.Secuencia, "")
		dataSheet.Fecha = append(dataSheet.Fecha, row[4])
		dataSheet.Cuenta = append(dataSheet.Cuenta, "") //FIXME: add ledger account
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
		dataSheet.Interface = append(dataSheet.Interface, currentTime)
	}

	//TODO: Send last consecutive to database
	fmt.Println(rowCount)
}
