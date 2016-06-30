package gonoov

import (
	"encoding/json"
	"net/http"
	"time"
)

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
	PageSize  int  `json:"pageSize"`
	PageTotal int  `json:"pageTotalElements"`
	LastPage  bool `json:"lastPage"`
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
	Model        []string `json:"modelo"`
	Number       string   `json:"numero,omitempty"`
	Serie        string   `json:"serie,omitempty"`
	ECnpj        []string `json:"emiCnpj"`
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
	Size         int      `json:"size"`
	NextProtocal int      `json:"nextProtocol,omitempty"`
	AllCnpj      bool     `json:"allCnpj"`
}

type InfProt struct {
	DigVal   string      `json:"digVal"`
	VerAplic string      `json:"verAplic"`
	DhRecbto time.Time   `json:"dhrecbto"`
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
	VICMSUFDest  string  `json:"vICMSUFDest"`
	VFCPUFDest   string  `json:"vFCPUFDest"`
	VBC          string  `json:"vBC"`
	VST          string  `json:"vST"`
	VProd        float64 `json:"vProd"`
	VTotTrib     string  `json:"vTotTrib"`
	VBCST        string  `json:"vBCST"`
	VCOFINS      string  `json:"vCOFINS"`
	VFrete       string  `json:"vFrete"`
	VOutro       string  `json:"vOutro"`
	VICMSDeson   string  `json:"vICMSDeson"`
	VII          string  `json:"vII"`
	VDesc        string  `json:"vDesc"`
	VICMSUFRemet string  `json:"vICMSUFRemet"`
	VIPI         string  `json:"vIPI"`
	VPIS         string  `json:"vPIS"`
	VICMS        string  `json:"vICMS"`
	VSeg         string  `json:"vSeg"`
	VNF          float64 `json:"vNF"`
}

type NfeTotal struct {
	ICMSTotal ICMSTotal `json:"ICMSTot"`
}

type NfeObsCont struct {
	XCampo string  `json:"xCampo"`
	XTexto float64 `json:"xTexto"`
}

type NfeInfAdic struct {
	ObsCont []NfeObsCont `json:"obsCont"`
	InfCpl  string       `json:"infCpl"`
}

type EnderDest struct {
	CEP     json.Number `json:"CEP"`
	Fone    json.Number `json:"fone"`
	Nro     float64     `json:"nro"`
	CMun    float64     `json:"cMun"`
	UF      string      `json:"UF"`
	CPais   float64     `json:"cPais"`
	XMun    string      `json:"xMun"`
	XPais   string      `json:"xPais"`
	XLgr    string      `json:"xLgr"`
	XBairro string      `json:"xBairro"`
}

type NfeDest struct {
	Cnpj      string      `json:"CNPJ"`
	EnderDest EnderDest   `json:"enderDest"`
	IE        json.Number `json:"IE"`
	IndIEDest float64     `json:"indIEDest"`
	Email     string      `json:"email"`
	XNome     string      `json:"xNome"`
}

type NfeVol struct {
	PesoL float64 `json:"pesoL"`
	Esp   string  `json:"esp"`
	QVol  float64 `json:"qVol"`
	PesoB float64 `json:"pesoB"`
}

type NfeTransp struct {
	ModFrete float64 `json:"modFrete"`
	Vol      NfeVol  `json:"vol"`
}

type NfeEnderEmit struct {
	XLgr    string      `json:"xLgr"`
	UF      string      `json:"UF"`
	Nro     float64     `json:"nro"`
	CMun    float64     `json:"cMun"`
	XBairro string      `json:"xBairro"`
	CEP     json.Number `json:"CEP"`
	Fone    json.Number `json:"fone"`
	XPais   string      `json:"xPais"`
	CPais   float64     `json:"cPais"`
	XMun    string      `json:"xMun"`
}

type NfeEmit struct {
	XFant     string       `json:"xFant"`
	CNPJ      string       `json:"CNPJ"`
	EnderEmit NfeEnderEmit `json:"enderEmit"`
	IE        json.Number  `json:"IE"`
	XNome     string       `json:"xNome"`
	CRT       float64      `json:"CRT"`
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
	NItem   float32 `json:"nItem"`
	Prod    Produto `json:"prod"`
}

type NfeIde struct {
	TpEmis   json.Number `json:"tpEmis"`
	TpNF     json.Number `json:"tpNF"`
	CMunFG   json.Number `json:"cMunFG"`
	DhSaiEnt string      `json:"dhSaiEnt"`
	CUF      json.Number `json:"cUF"`
	Mod      json.Number `json:"mod"`
	DhEmi    string      `json:"dhEmi"`
	TpAmb    json.Number `json:"tpAmb"`
	TpImp    json.Number `json:"tpImp"`
	FinNFe   json.Number `json:"finNFe"`
	IndFinal json.Number `json:"indFinal"`
	NatOp    string      `json:"natOp"`
	ProcEmi  json.Number `json:"procEmi"`
	IDDest   json.Number `json:"idDest"`
	NNF      json.Number `json:"nNF"`
	VerProc  string      `json:"verProc"`
	IndPag   json.Number `json:"indPag"`
	IndPres  json.Number `json:"indPres"`
	Serie    json.Number `json:"serie"`
	CNF      json.Number `json:"cNF"`
	CDV      json.Number `json:"cDV"`
}

type InfNfe struct {
	ID      string     `json:"Id"`
	Version float32    `json:"versao"`
	Total   NfeTotal   `json:"total"`
	InfAdic NfeInfAdic `json:"infAdic"`
	Dest    NfeDest    `json:"dest"`
	Transp  NfeTransp  `json:"transp"`
	Emit    NfeEmit    `json:"emit"`
	Det     []NfeDet   `json:"det"`
	Ide     NfeIde     `json:"ide"`
}

type NFe struct {
	Xmlns  string `json:"xmlns"`
	InfNfe InfNfe `json:"infNFe"`
}

type NfeProc struct {
	Xmlns   string `json:"xmlns"`
	ProtNfe ProNfe `json:"protNFe"`

	Xmlns2 string `json:"xmlns:ns2"`
	NFe    NFe    `json:"NFe"`

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
