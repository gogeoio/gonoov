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
	statsUrl = "app/stats/emissao"
)

type StatsParams struct {
	EnvironmentType int        `json:"tpAmb"`
	ECnpj           []string   `json:"emiDoc,omitempty"`
	EStartDate      int64      `json:"emDataInicial,omitempty"`
	EEndDate        int64      `json:"emDataFinal,omitempty"`
	AllCnpj         bool       `json:"allCnpj,omitempty"`
	Size            int        `json:"pageSize,omitempty"`
	Page            int        `json:"page,omitempty"`
	NextProtocol    NoovString `json:"nextProtocol,omitempty"`
}

type Stats struct {
	Emitentes  []string `json:"emitentes"`
	EStartDate int64    `json:"emDataInicial,omitempty"`
	EEndDate   int64    `json:"emDataFinal,omitempty"`
	Total      int      `json:"totalNfesEmitidas,omitempty"`
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

	m := make(map[string]Stats)
	json.Unmarshal(body, &m)

	return m["data"], nil
}

func getStatsUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, statsUrl)
}
