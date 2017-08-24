package gonoov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-querystring/query"
)

const (
	totalizerUrl = "app/nfe/totalizer"
)

type TotalizerParams struct {
	NFEKey string   `json:"chave"    url:"chave,omitempty"`
	Emits  []string `json:"emits"    url:"emits,omitempty"`
	CFOPs  []int    `json:"cfops"    url:"cfops,omitempty"`
	Day    string   `json:"dia"      url:"dia,omitempty"`
	Seller string   `json:"vendedor" url:"vendedor,omitempty"`
}

type Totalizer struct {
	Day        string   `json:"dia"`
	Emit       string   `json:"emitente"`
	NFEKey     string   `json:"chNF"`
	SalesmanID string   `json:"vendedor"`
	NFEValue   float64  `json:"valorNF"`
	CFOPs      []string `json:"cfops"`
	NFEType    float64  `json:"tipoNF"`
	Raw        []byte   `json:"-"`
}

type TotalizerResponse struct {
	Totalizers []Totalizer `json:"totalizers"`
	Raw        []byte      `json:"-"`
}

func (noov *Noov) Totalizer(params TotalizerParams) (TotalizerResponse, error) {
	totalizer := TotalizerResponse{}

	urlParams, err := query.Values(params)
	if err != nil {
		log.Printf("could not parse totalizer params: %v", err)
		return totalizer, err
	}

	// TODO Testar formato da data 2017-03-12

	url := getTotalizerUrl(noov)
	url = fmt.Sprintf("%s?%s", url, urlParams.Encode())

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	noov.Authenticate()
	setRequestHeaders(req, noov.Token)
	resp, err := noov.client.Do(req)

	if err != nil {
		log.Printf("could not execute totalizer request: %v", err)
		return totalizer, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("could not read totalizer response body: %v", err)
		return totalizer, err
	}

	if resp.StatusCode == http.StatusGatewayTimeout {
		return totalizer, fmt.Errorf("%s", "Gateway timeout")
	}

	// not found data
	if resp.StatusCode == http.StatusNotFound {
		return totalizer, nil
	}

	if resp.StatusCode != http.StatusOK {
		json.Unmarshal(body, &totalizer)
		return totalizer, fmt.Errorf("%s", body)
	}

	m := make(map[string][]Totalizer)
	// TODO Tratar erro de unmarshal
	json.Unmarshal(body, &m)

	totalizers := m["data"]
	totalizer.Totalizers = totalizers
	totalizer.Raw = body

	return totalizer, nil
}

func getTotalizerUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, totalizerUrl)
}
