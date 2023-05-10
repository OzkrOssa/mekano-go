package mekano

import (
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

type MekanoDataSheet struct {
	Tipo          string
	Prefijo       string
	Numero        string
	Secuencia     string
	Fecha         string
	Cuenta        string
	Terceros      string
	CentroCostos  string
	Nota          string
	Debito        string
	Credito       string
	Base          string
	Aplica        string
	TipoAnexo     string
	PrefijoAnexo  string
	NumeroAnexo   string
	Usuario       string
	Signo         string
	CuentaCobrar  string
	CuentaPagar   string
	NombreTercero string
	NombreCentro  string
	Interface     string
}

type paymentStatistics struct {
	RangoRC     string `json:"rango-rc"`
	Bancolombia int    `json:"bancolombia"`
	Davivienda  int    `json:"davivienda"`
	Susuerte    int    `json:"susuerte"`
	PayU        int    `json:"payu"`
	Efectivo    int    `json:"efectivo"`
	Total       int    `json:"total"`
}
type billingStatistics struct {
	Debito  float64 `json:"debito"`
	Credito float64 `json:"credito"`
	Base    float64 `json:"base"`
}

func xlsxData(filePath string, fileName string) ([][]string, error) {
	xlsx, err := excelize.OpenFile(filepath.Join(filePath, fileName+".xlsx"))
	if err != nil {
		return nil, err
	}

	excelRows, err := xlsx.GetRows(xlsx.GetSheetName(0))
	if err != nil {
		return nil, err
	}

	return excelRows, nil
}
