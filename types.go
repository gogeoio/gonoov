package gonoov

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NoovString string

func (v *NoovString) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)

	if err != nil {
		s = fmt.Sprintf("%s", data)
	}

	*v = NoovString(s)

	return nil
}

type NoovTime struct {
	Time  time.Time
	Valid bool
}

func (t *NoovTime) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	var tt time.Time

	err = json.Unmarshal(data, &s)

	if err != nil {
		*t = NoovTime{Time: time.Now(), Valid: false}
		return nil
	}

	layouts := []string{"2006-01-02T15:04:05", "2006-01-02T15:04:05-07:00"}

	for _, layout := range layouts {
		tt, err = time.Parse(layout, s)

		if err == nil {
			break
		}
	}

	if err != nil {
		*t = NoovTime{Time: time.Now(), Valid: false}
		return nil
	}

	*t = NoovTime{Time: tt, Valid: true}
	return err
}

type Noov struct {
	ApiKey    string
	ApiSecret string
	url       string
	version   string
	appname   string
	email     string
	Token     string
	client    *http.Client
}

type LoginParams struct {
	ApiKey    string
	ApiSecret string
	AppName   string
	Email     string
}

type Pagination struct {
	PageSize  int  `json:"pageSize,omitempty"`
	PageTotal int  `json:"pageTotalElements,omitempty"`
	LastPage  bool `json:"lastPage,omitempty"`
}

type StaticPagination struct {
	Pagination
	Number        int   `json:"number,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
}

type DynamicPagination struct {
	Pagination
	NextProtocol int64 `json:"nextProtocol,omitempty"`
}

type NoovError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type Meta struct {
	Errors []NoovError `json:"errors"`
}

type GenericResponse struct {
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
}

/* ------------------------------------------------------------------------- */
/*                                                                           */
/*                                 NFe types                                 */
/*                                                                           */
/* ------------------------------------------------------------------------- */

type NfeRawResponse struct {
	GenericResponse
	Data []NfeResponse
}

type NfeParams struct {
	DynamicPagination
	Model        []string `json:"modelo"`
	Number       string   `json:"numero,omitempty"`
	Serie        string   `json:"serie,omitempty"`
	ECnpj        []string `json:"emiCnpj,omitempty"`
	DCnpj        []string `json:"destCnpj,omitempty"`
	Cancelled    bool     `json:"cancelados"`
	Cean         []string `json:"cean,omitempty"`
	Key          string   `json:"chave,omitempty"`
	EEndDate     int64    `json:"emDataFinal,omitempty"`
	EStartDate   int64    `json:"emDataInicial,omitempty"`
	EDate        int64    `json:"emData,omitempty"`
	REndDate     int64    `json:"recDataFinal,omitempty"`
	RStartDate   int64    `json:"recDataInicial,omitempty"`
	RDate        int64    `json:"recData,omitempty"`
	ECity        string   `json:"emitCidade,omitempty"`
	EState       string   `json:"emitUF,omitempty"`
	Size         int      `json:"size,omitempty"`
	NextProtocol int      `json:"nextProtocol,omitempty"`
	AllCnpj      bool     `json:"allCnpj"`
}

type InfProt struct {
	DigVal   string      `json:"digVal"`
	VerAplic string      `json:"verAplic"`
	DhRecbto NoovTime    `json:"dhrecbto"`
	ChNfe    string      `json:"chNFe"`
	XMotivo  string      `json:"xMotivo"`
	TpAmb    float32     `json:"tpAmb"`
	CStat    float32     `json:"cStat"`
	NProt    json.Number `json:"nProt"`
}

type ProNfe struct {
	Version float32 `json:"version"`
	InfProt InfProt `json:"infProt"`
}

type ICMSTotal struct {
	VICMSUFDest  json.Number `json:"vICMSUFDest"`
	VFCPUFDest   json.Number `json:"vFCPUFDest"`
	VBC          json.Number `json:"vBC"`
	VST          json.Number `json:"vST"`
	VProd        json.Number `json:"vProd"`
	VTotTrib     json.Number `json:"vTotTrib"`
	VBCST        json.Number `json:"vBCST"`
	VCOFINS      json.Number `json:"vCOFINS"`
	VFrete       json.Number `json:"vFrete"`
	VOutro       json.Number `json:"vOutro"`
	VICMSDeson   json.Number `json:"vICMSDeson"`
	VII          json.Number `json:"vII"`
	VDesc        json.Number `json:"vDesc"`
	VICMSUFRemet json.Number `json:"vICMSUFRemet"`
	VIPI         json.Number `json:"vIPI"`
	VPIS         json.Number `json:"vPIS"`
	VICMS        json.Number `json:"vICMS"`
	VSeg         json.Number `json:"vSeg"`
	VNF          json.Number `json:"vNF"`
}

type NfeTotal struct {
	ICMSTotal ICMSTotal `json:"ICMSTot"`
}

type NfeObsCont struct {
	XCampo string `json:"xCampo"`
	XTexto string `json:"xTexto"`
}

type NfeInfAdic struct {
	// TODO Criar um tipo para fazer parser de objeto para array
	// Enviar email para Oobj para tentar alterar esse formato
	//ObsCont []NfeObsCont `json:"obsCont"`
	InfCpl string `json:"infCpl"`
}

type NfeAddress struct {
	CEP     string     `json:"CEP"`
	CMun    NoovString `json:"cMun"`
	CPais   NoovString `json:"cPais"`
	Fone    string     `json:"fone"`
	Nro     NoovString `json:"nro"`
	UF      string     `json:"UF"`
	XMun    string     `json:"xMun"`
	XPais   string     `json:"xPais"`
	XBairro string     `json:"xBairro"`
	XLgr    NoovString `json:"xLgr"`
	XCpl    NoovString `json:"xCpl"`
}

type NfeDest struct {
	Cnpj      string     `json:"CNPJ"`
	Email     string     `json:"email"`
	EnderDest NfeAddress `json:"enderDest"`
	IE        NoovString `json:"IE"`
	IndIEDest int        `json:"indIEDest"`
	XNome     string     `json:"xNome"`
}

type NfeVol struct {
	PesoL json.Number `json:"pesoL"`
	Esp   string      `json:"esp"`
	QVol  int         `json:"qVol"`
	PesoB json.Number `json:"pesoB"`
}

type NfeTransp struct {
	ModFrete int    `json:"modFrete"`
	Vol      NfeVol `json:"vol"`
}

type NfeEmit struct {
	XFant     string     `json:"xFant"`
	CNPJ      string     `json:"CNPJ"`
	EnderEmit NfeAddress `json:"enderEmit"`
	XNome     string     `json:"xNome"`
	CRT       int        `json:"CRT"`
	IE        NoovString `json:"IE"`
}

type ICMS struct {
	ICMS60 struct {
		CST        float64 `json:"CST"`
		VBCSTRet   string  `json:"vBCSTRet"`
		VICMSSTRet string  `json:"vICMSSTRet"`
		Orig       float64 `json:"orig"`
	} `json:"ICMS60"`
}

type COFINS struct {
	COFINSAliq struct {
		VCOFINS string `json:"vCOFINS"`
		CST     string `json:"CST"`
		VBC     string `json:"vBC"`
		PCOFINS string `json:"pCOFINS"`
	} `json:"COFINSAliq"`
}

type PIS struct {
	PISAliq struct {
		VPIS string `json:"vPIS"`
		CST  string `json:"CST"`
		VBC  string `json:"vBC"`
		PPIS string `json:"pPIS"`
	} `json:"PISAliq"`
}

type Imposto struct {
	ICMS     ICMS   `json:"ICMS"`
	COFINS   COFINS `json:"COFINS"`
	PIS      PIS    `json:"PIS"`
	VTotTrib string `json:"vTotTrib"`
}

type Produto struct {
	XProd    string      `json:"xProd"`
	CFOP     json.Number `json:"CFOP"`
	UCom     string      `json:"uCom"`
	UTrib    string      `json:"uTrib"`
	CEANTrib json.Number `json:"cEANTrib"`
	VUnCom   json.Number `json:"vUnCom,number"`
	CProd    json.Number `json:"cProd"`
	CEST     string      `json:"CEST"`
	VUnTrib  json.Number `json:"vUnTrib"`
	IndTot   json.Number `json:"indTot"`
	QCom     json.Number `json:"qCom"`
	VProd    json.Number `json:"vProd"`
	QTrib    json.Number `json:"qTrib"`
	NCM      json.Number `json:"NCM"`
	CEAN     json.Number `json:"cEAN"`
}

type NfeDet struct {
	Imposto Imposto `json:"imposto"`
	NItem   int     `json:"nItem"`
	Prod    Produto `json:"prod"`
}

type NfeIde struct {
	DhSaiEnt string      `json:"dhSaiEnt"`
	TpEmis   json.Number `json:"tpEmis"`
	TpNF     json.Number `json:"tpNF"`
	CMunFG   json.Number `json:"cMunFG"`
	CUF      json.Number `json:"cUF"`
	Mod      json.Number `json:"mod"`
	TpAmb    json.Number `json:"tpAmb"`
	TpImp    json.Number `json:"tpImp"`
	FinNFe   json.Number `json:"finNFe"`
	IndFinal json.Number `json:"indFinal"`
	ProcEmi  json.Number `json:"procEmi"`
	IDDest   json.Number `json:"idDest"`
	NNF      json.Number `json:"nNF"`
	IndPag   json.Number `json:"indPag"`
	IndPres  json.Number `json:"indPres"`
	Serie    json.Number `json:"serie"`
	CDV      json.Number `json:"cDV"`
	CNF      NoovString  `json:"cNF"`
	NatOp    NoovString  `json:"natOp"`
	VerProc  NoovString  `json:"verProc"`
	DhEmi    NoovTime    `json:"dhEmi"`
}

type InfNfe struct {
	ID      string     `json:"Id"`
	Version float32    `json:"versao"`
	Total   NfeTotal   `json:"total"`
	InfAdic NfeInfAdic `json:"infAdic"`
	Transp  NfeTransp  `json:"transp"`
	Ide     NfeIde     `json:"ide"`
	Dest    NfeDest    `json:"dest"`
	Emit    NfeEmit    `json:"emit"`

	// TODO Tratar quando for array ou objeto
	//Det     []NfeDet   `json:"det"`
}

type NFe struct {
	Xmlns  string `json:"xmlns"`
	InfNfe InfNfe `json:"infNFe"`
}

type NfeProc struct {
	Xmlns   string  `json:"xmlns"`
	ProtNfe ProNfe  `json:"protNFe"`
	Xmlns2  string  `json:"xmlns:ns2"`
	NFe     NFe     `json:"NFe"`
	Version float32 `json:"versao"`
}

type NfeResponse struct {
	ID      string  `json:"id"`
	NfeProc NfeProc `json:"nfeProc"`
}

/* ------------------------------------------------------------------------- */
/*                                                                           */
/*                             Unexported types                              */
/*                                                                           */
/* ------------------------------------------------------------------------- */

type noovLoginParams struct {
	ApiKey    string `json:"apiKey"`
	Timestamp int64  `json:"timestamp"`
	Secret    string `json:"secret"`
}

type token struct {
	Token string `json:"token"`
}
