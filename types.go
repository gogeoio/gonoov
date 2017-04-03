package gonoov

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a8m/djson"
)

// Custom type to transform any type to string
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

func (v NoovString) String() string {
	return string(v)
}

type NoovTime struct {
	Time  time.Time
	Valid bool
}

func NewNoovTime(s string) (NoovTime, error) {
	layouts := []string{"2006-01-02T15:04:05", "2006-01-02T15:04:05-07:00", "2006-01-02"}

	var t NoovTime
	var err error
	var tt time.Time

	for _, layout := range layouts {
		tt, err = time.Parse(layout, s)

		if err == nil {
			break
		}
	}

	if err != nil {
		t = NoovTime{Valid: false}
		return t, err
	}

	t = NoovTime{Time: tt, Valid: true}

	return t, err
}

func (t *NoovTime) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Replace(s, "\"", "", -1)
	tt, err := NewNoovTime(s)

	*t = tt

	return err
}

type Noov struct {
	ApiKey         string
	ApiSecret      string
	url            string
	version        string
	appname        string
	email          string
	Token          string
	TokenTimestamp int64 // token timestamp
	client         *http.Client
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
	NextProtocol NoovString `json:"nextProtocol,omitempty"`
}

type NoovError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type Meta struct {
	Errors []NoovError `json:"errors"`
}

type MetaResponse struct {
	Meta Meta `json:"meta"`
}

/* ------------------------------------------------------------------------- */
/*                                                                           */
/*                                 NFe types                                 */
/*                                                                           */
/* ------------------------------------------------------------------------- */

type NfeParams struct {
	Model        []string   `json:"modelo"`
	Number       string     `json:"numero,omitempty"`
	Serie        string     `json:"serie,omitempty"`
	ECnpj        []string   `json:"emiDoc,omitempty"`
	DCnpj        []string   `json:"destDoc,omitempty"`
	Cancelled    bool       `json:"cancelados"`
	Cean         []string   `json:"cean,omitempty"`
	Key          string     `json:"chave,omitempty"`
	EEndDate     int64      `json:"emDataFinal,omitempty"`
	EStartDate   int64      `json:"emDataInicial,omitempty"`
	EDate        int64      `json:"emData,omitempty"`
	REndDate     int64      `json:"recDataFinal,omitempty"`
	RStartDate   int64      `json:"recDataInicial,omitempty"`
	RDate        int64      `json:"recData,omitempty"`
	ECity        string     `json:"emitCidade,omitempty"`
	EState       string     `json:"emitUF,omitempty"`
	Size         int        `json:"pageSize,omitempty"`
	NextProtocol NoovString `json:"nextProtocol,omitempty"`
	AllCnpj      bool       `json:"allCnpj"`
}

type InfProt struct {
	DigVal   NoovString  `json:"digVal"`
	VerAplic NoovString  `json:"verAplic"`
	DhRecbto NoovTime    `json:"dhrecbto"`
	ChNfe    NoovString  `json:"chNFe"`
	XMotivo  NoovString  `json:"xMotivo"`
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
	XCampo NoovString `json:"xCampo"`
	XTexto NoovString `json:"xTexto"`
}

type NfeInfAdic struct {
	ObsCont NfeObsContArray `json:"obsCont"`
	InfCpl  NoovString      `json:"infCpl"`
}

func (nia *NfeInfAdic) UnmarshalJSON(data []byte) error {
	type Alias NfeInfAdic

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(nia),
	}

	err := json.Unmarshal(data, &aux)

	if err != nil {
		//log.Println("----------------------> data", string(data))
	}

	return nil
}

type NfeObsContArray []NfeObsCont

func (nv *NfeObsContArray) UnmarshalJSON(data []byte) error {
	dj := djson.NewDecoder(data)
	array, err := dj.DecodeArray()

	result := []NfeObsCont{}

	if err != nil {
		err = nil

		obj, err := dj.DecodeObject()

		if err != nil {
			return err
		}

		item := NfeObsCont{}
		item.XCampo = NoovString(interfaceToString(obj["xCampo"]))
		item.XTexto = NoovString(interfaceToString(obj["xTexto"]))
		result = append(result, item)

	} else {
		// Transform []interface{} --> []map[string]interface{}
		m := make([]map[string]interface{}, len(array))
		for i, item := range array {
			m[i] = item.(map[string]interface{})
		}

		for _, mm := range m {
			item := NfeObsCont{}
			item.XCampo = NoovString(interfaceToString(mm["xCampo"]))
			item.XTexto = NoovString(interfaceToString(mm["xTexto"]))
			result = append(result, item)
		}
	}

	*nv = NfeObsContArray(result)

	return err
}

func interfaceToString(v interface{}) string {
	vv, ok := v.(string)

	if !ok {
		vv = fmt.Sprintf("%v", v)
	}

	return vv
}

type NfeAddress struct {
	CEP     NoovString `json:"CEP"`
	CMun    NoovString `json:"cMun"`
	CPais   NoovString `json:"cPais"`
	Fone    NoovString `json:"fone"`
	Nro     NoovString `json:"nro"`
	UF      NoovString `json:"UF"`
	XMun    NoovString `json:"xMun"`
	XPais   NoovString `json:"xPais"`
	XBairro NoovString `json:"xBairro"`
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
	Marca NoovString  `json:"marca"`
	PesoL json.Number `json:"pesoL"`
	Esp   NoovString  `json:"esp"`
	QVol  NoovString  `json:"qVol"`
	PesoB json.Number `json:"pesoB"`
}

func (nv *NfeVol) UnmarshalJSON(data []byte) error {
	type Alias NfeVol

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(nv),
	}

	err := json.Unmarshal(data, &aux)

	if err != nil {
		err = nil

		dj := djson.NewDecoder(data)
		array, err := dj.DecodeArray()

		if err != nil {
			return err
		}

		// Transform []interface{} --> []map[string]interface{}
		m := make([]map[string]interface{}, len(array))
		for i, item := range array {
			m[i] = item.(map[string]interface{})
		}

		fm := map[string]interface{}{}

		for _, mmap := range m {
			for k, v := range mmap {
				fm[k] = v
			}
		}

		d, err := json.Marshal(fm)

		if err != nil {
			return err
		}

		err = json.Unmarshal(d, &aux)
	}

	return err
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
		CST        NoovString `json:"CST"`
		VBCSTRet   NoovString `json:"vBCSTRet"`
		VICMSSTRet NoovString `json:"vICMSSTRet"`
		Orig       NoovString `json:"orig"`
	} `json:"ICMS60"`
}

type COFINS struct {
	COFINSAliq struct {
		VCOFINS NoovString `json:"vCOFINS"`
		CST     NoovString `json:"CST"`
		VBC     NoovString `json:"vBC"`
		PCOFINS NoovString `json:"pCOFINS"`
	} `json:"COFINSAliq"`
}

type PIS struct {
	PISAliq struct {
		VPIS NoovString `json:"vPIS"`
		CST  NoovString `json:"CST"`
		VBC  NoovString `json:"vBC"`
		PPIS NoovString `json:"pPIS"`
	} `json:"PISAliq"`
}

type Imposto struct {
	ICMS     ICMS       `json:"ICMS"`
	COFINS   COFINS     `json:"COFINS"`
	PIS      PIS        `json:"PIS"`
	VTotTrib NoovString `json:"vTotTrib"`
}

type Produto struct {
	CEAN     NoovString  `json:"cEAN"`
	CProd    NoovString  `json:"cProd"`
	CEANTrib NoovString  `json:"cEANTrib"`
	CEST     NoovString  `json:"CEST"`
	CFOP     json.Number `json:"CFOP"`
	IndTot   json.Number `json:"indTot"`
	QCom     json.Number `json:"qCom"`
	QTrib    json.Number `json:"qTrib"`
	UCom     NoovString  `json:"uCom"`
	UTrib    NoovString  `json:"uTrib"`
	VProd    json.Number `json:"vProd"`
	VFrete   json.Number `json:"vFrete"`
	VUnCom   NoovString  `json:"vUnCom,number"`
	VUnTrib  NoovString  `json:"vUnTrib"`
	XProd    string      `json:"xProd"`

	// TODO Criar marshal para não deixar como notação científica
	//NCM json.Number `json:"NCM"`
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
	DHEmi    NoovTime    `json:"dhEmi"`
}

func (n *NfeIde) UnmarshalJSON(data []byte) error {
	type Alias NfeIde
	alias := Alias{}
	json.Unmarshal(data, &alias)

	if !alias.DHEmi.Valid {
		type dhEmi struct {
			DEmi NoovTime `json:"dEmi"`
		}

		temp := dhEmi{}
		json.Unmarshal(data, &temp)

		alias.DHEmi = temp.DEmi
	}

	nn := NfeIde(alias)
	*n = nn

	return nil
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
	Det     NoovNfeDet `json:"det"`
}

type NoovNfeDet []NfeDet

func (det *NoovNfeDet) UnmarshalJSON(data []byte) error {
	array := []NfeDet{}

	err := json.Unmarshal(data, &array)

	if err != nil {
		// TODO handle it
		var item NfeDet

		err = json.Unmarshal(data, &item)

		if err != nil {
			return err
		}

		array = append(array, item)
	}

	*det = NoovNfeDet(array)

	return err
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
	ID         string     `json:"id"`
	NfeProc    NfeProc    `json:"nfeProc"`
	Enrichment Enrichment `json:"enrichment"`
}

type Enrichment struct {
	CodVendedor   string `json:"codvendedor"`
	NomeVendedor  string `json:"nomevendedor"`
	CodSupervisor string `json:"codsupervisor"`
	CodTipologia  string `json:"tipologia"`
	NomeTipologia string `json:"nometipologia"`
}

type NfeRawResponse struct {
	MetaResponse
	Data       []NfeResponse
	Pagination DynamicPagination `json:"pagination"`
	Raw        []byte            `json:"-"`
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
