package mekano

type PaymentData struct {
	Tipo          string
	Prefijo       string
	Numero        int
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

type statistics struct {
	RangoRC     string `json:"rango-rc"`
	Bancolombia string `json:"bancolombia"`
	Davivienda  string `json:"davivienda"`
	Susuerte    string `json:"susuerte"`
	PayU        string `json:"payu"`
	Efectivo    string `json:"efectivo"`
	Total       string `json:"total"`
}
