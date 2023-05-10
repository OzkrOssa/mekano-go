package mekano

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/OzkrOssa/mekano-go/utils"
	"github.com/mozillazg/go-unidecode"
)

func Billing(fileName string) {
	billingFile, err := xlsxData(utils.BillingFileDirPath, fileName)
	var montoBaseFinal float64
	if err != nil {
		log.Println(err, "billingFile")
	}
	itemsIvaFile, err := xlsxData("C:/Users/devre/OneDrive/facturacion_mekano", "extras")

	var BillingDataSheet []MekanoDataSheet

	if err != nil {
		log.Println(err, "itemsIvaFile")
	}

	for _, bRow := range billingFile[1:] {

		montoBase, err := strconv.ParseFloat(bRow[12], 64)

		if err != nil {
			log.Println(err, "MontoBase")
		}

		montoIva, err := strconv.ParseFloat(strings.TrimSpace(bRow[13]), 64)
		if err != nil {
			log.Println(err, "MontoIva")
		}

		_, decimal := math.Modf(montoBase)

		if decimal >= 0.5 {
			montoBaseFinal = math.Ceil(montoBase)
		} else {
			montoBaseFinal = math.Round(montoBase)
		}

		if !strings.Contains(bRow[21], ",") {
			billingNormal := MekanoDataSheet{
				Tipo:          "FVE",
				Prefijo:       "_",
				Numero:        bRow[8],
				Secuencia:     "",
				Fecha:         bRow[9],
				Cuenta:        utils.Cuentas[bRow[21]],
				Terceros:      bRow[1],
				CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
				Nota:          "FACTURA ELECTRÓNICA DE VENTA",
				Debito:        "0",
				Credito:       fmt.Sprintf("%f", math.Ceil(montoBase)),
				Base:          "0",
				Aplica:        "",
				TipoAnexo:     "",
				PrefijoAnexo:  "",
				NumeroAnexo:   "",
				Usuario:       "SUPERVISOR",
				Signo:         "",
				CuentaCobrar:  "",
				CuentaPagar:   "",
				NombreTercero: bRow[2],
				NombreCentro:  bRow[17],
				Interface:     utils.CurrentTimeMekanoInterface,
			}

			BillingDataSheet = append(BillingDataSheet, billingNormal)

			billingIva := MekanoDataSheet{
				Tipo:          "FVE",
				Prefijo:       "_",
				Numero:        bRow[8],
				Secuencia:     "",
				Fecha:         bRow[9],
				Cuenta:        "24080505",
				Terceros:      bRow[1],
				CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
				Nota:          "FACTURA ELECTRÓNICA DE VENTA",
				Debito:        "0",
				Credito:       fmt.Sprintf("%f", montoIva),
				Base:          fmt.Sprintf("%f", montoBaseFinal),
				Aplica:        "",
				TipoAnexo:     "",
				PrefijoAnexo:  "",
				NumeroAnexo:   "",
				Usuario:       "SUPERVISOR",
				Signo:         "",
				CuentaCobrar:  "",
				CuentaPagar:   "",
				NombreTercero: bRow[2],
				NombreCentro:  bRow[17],
				Interface:     utils.CurrentTimeMekanoInterface,
			}

			BillingDataSheet = append(BillingDataSheet, billingIva)

			billingCxC := MekanoDataSheet{
				Tipo:          "FVE",
				Prefijo:       "_",
				Numero:        bRow[8],
				Secuencia:     "",
				Fecha:         bRow[9],
				Cuenta:        "13050501",
				Terceros:      bRow[1],
				CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
				Nota:          "FACTURA ELECTRÓNICA DE VENTA",
				Debito:        bRow[14],
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
				NombreTercero: bRow[2],
				NombreCentro:  bRow[17],
				Interface:     utils.CurrentTimeMekanoInterface,
			}

			BillingDataSheet = append(BillingDataSheet, billingCxC)
		} else {
			splitBillingItems := strings.Split(bRow[21], ",")
			for _, item := range splitBillingItems {
				for _, itemIva := range itemsIvaFile[1:] {
					if itemIva[1] == strings.TrimSpace(item) && itemIva[0] == bRow[0] {
						itemIvaBase, _ := strconv.ParseFloat(itemIva[2], 64)
						billingNormalPlus := MekanoDataSheet{
							Tipo:          "FVE",
							Prefijo:       "_",
							Numero:        bRow[8],
							Secuencia:     "",
							Fecha:         bRow[9],
							Cuenta:        utils.Cuentas[unidecode.Unidecode(strings.TrimSpace(item))],
							Terceros:      bRow[1],
							CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
							Nota:          "FACTURA ELECTRÓNICA DE VENTA",
							Debito:        "0",
							Credito:       fmt.Sprintf("%f", math.Ceil(itemIvaBase-1)),
							Base:          "0",
							Aplica:        "",
							TipoAnexo:     "",
							PrefijoAnexo:  "",
							NumeroAnexo:   "",
							Usuario:       "SUPERVISOR",
							Signo:         "",
							CuentaCobrar:  "",
							CuentaPagar:   "",
							NombreTercero: bRow[2],
							NombreCentro:  bRow[17],
							Interface:     utils.CurrentTimeMekanoInterface,
						}
						BillingDataSheet = append(BillingDataSheet, billingNormalPlus)
					}
				}
			}
			billingIvaPlus := MekanoDataSheet{
				Tipo:          "FVE",
				Prefijo:       "_",
				Numero:        bRow[8],
				Secuencia:     "",
				Fecha:         bRow[9],
				Cuenta:        "24080505",
				Terceros:      bRow[1],
				CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
				Nota:          "FACTURA ELECTRÓNICA DE VENTA",
				Debito:        "0",
				Credito:       fmt.Sprintf("%f", montoIva),
				Base:          fmt.Sprintf("%f", montoBaseFinal),
				Aplica:        "",
				TipoAnexo:     "",
				PrefijoAnexo:  "",
				NumeroAnexo:   "",
				Usuario:       "SUPERVISOR",
				Signo:         "",
				CuentaCobrar:  "",
				CuentaPagar:   "",
				NombreTercero: bRow[2],
				NombreCentro:  bRow[17],
				Interface:     utils.CurrentTimeMekanoInterface,
			}

			BillingDataSheet = append(BillingDataSheet, billingIvaPlus)

			billingCxCPlus := MekanoDataSheet{
				Tipo:          "FVE",
				Prefijo:       "_",
				Numero:        bRow[8],
				Secuencia:     "",
				Fecha:         bRow[9],
				Cuenta:        "13050501",
				Terceros:      bRow[1],
				CentroCostos:  utils.CentroCostos[unidecode.Unidecode(bRow[17])],
				Nota:          "FACTURA ELECTRÓNICA DE VENTA",
				Debito:        bRow[14],
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
				NombreTercero: bRow[2],
				NombreCentro:  bRow[17],
				Interface:     utils.CurrentTimeMekanoInterface,
			}

			BillingDataSheet = append(BillingDataSheet, billingCxCPlus)
		}
	}

	txtFile, err := os.Create(filepath.Join(utils.MekanoInterfaceDirPath, "CONTABLE.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	csvFile, err := os.Create(filepath.Join(utils.MekanoInterfaceDirPath, "CONTABLE.csv"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer txtFile.Close()

	writer := csv.NewWriter(txtFile)
	w := csv.NewWriter(csvFile)
	writer.Comma = ','
	w.Comma = ','

	for _, data := range BillingDataSheet {
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
		w.Write(row)
	}
	writer.Flush()
	w.Flush()
	BillingStatistics(BillingDataSheet)
}
