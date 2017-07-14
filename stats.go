package gonoov

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	statsUrl = "app/nfe/stats/emissao"
)

type StatsParams struct {
	EnvironmentType int        `json:"tpAmb"`
	ECnpj           []string   `json:"emiDoc,omitempty"`
	EStartDate      string     `json:"emData,omitempty"`
	AllCnpj         bool       `json:"allCnpj,omitempty"`
	Size            int        `json:"pageSize,omitempty"`
	Page            int        `json:"page,omitempty"`
	NextProtocol    NoovString `json:"nextProtocol,omitempty"`
}

type Stats struct {
	MetaResponse
	Emitentes  []string `json:"emitentes"`
	EStartDate string   `json:"emData,omitempty"`
	Total      int      `json:"totalNfesEmitidas,omitempty"`
	Raw        []byte   `json:"-"`
}

func (noov *Noov) Stats(params StatsParams) (Stats, error) {
	stats := Stats{
		Total: 1,
	}

	b, err := json.Marshal(params)
	if err != nil {
		log.Printf("could not marshal params: %v", err)
		return stats, err
	}

	if params.EnvironmentType == 0 {
		// Tipo de ambiente padrão
		params.EnvironmentType = 1
	}

	url := getStatsUrl(noov)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))

	noov.Authenticate()
	setRequestHeaders(req, noov.Token)
	resp, err := noov.client.Do(req)

	if err != nil {
		return stats, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("could not read response body: %v", err)
		return stats, err
	}

	if resp.StatusCode == http.StatusGatewayTimeout {
		return stats, fmt.Errorf("%s", "Gatetway timeout")
	}

	if resp.StatusCode != http.StatusOK {
		json.Unmarshal(body, &stats)
		return stats, fmt.Errorf("%s", body)
	}

	m := make(map[string]Stats)
	json.Unmarshal(body, &m)

	stats = m["data"]
	stats.Raw = body

	return stats, nil
}

func getStatsUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, statsUrl)
}
