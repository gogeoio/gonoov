package gonoov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

var (
	notFoundError = []byte(`{
		"meta" : {
			"errors" : [ {
				"error" : "NotFoundException",
				"message" : "Nao foi possivel encontrar NFe."
			} ]
		}
	}`)
)

func TestNoov_Get(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	m := make(map[string]interface{})
	fixture, _ := readFixture("fixtures/nfe/fixture-1.json")
	json.Unmarshal(fixture, &m)

	url := getNfeUrl(noov)
	registerNfeGetResponder(ast, url, noov.Token, 200, m)

	params := NfeParams{}
	nfes, err := noov.GetNfe(params)
	ast.NoError(err)
	ast.NotEmpty(nfes.Data)

	ast.Equal(float32(3.1), nfes.Data[0].NfeProc.Version)
}

func TestNoov_GetWithError(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	m := make(map[string]interface{})
	json.Unmarshal(notFoundError, &m)

	url := getNfeUrl(noov)
	registerNfeGetResponder(ast, url, noov.Token, 404, m)

	params := NfeParams{}
	nfes, err := noov.GetNfe(params)
	ast.Error(err)
	ast.Empty(nfes.Data)
}

func TestNoov_GetWithInvalidTime(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	files := []string{"fixture-2", "fixture-3"}

	for _, file := range files {
		m := make(map[string]interface{})
		fixture, _ := readFixture(fmt.Sprintf("fixtures/nfe/%s.json", file))
		json.Unmarshal(fixture, &m)

		url := getNfeUrl(noov)
		registerNfeGetResponder(ast, url, noov.Token, 200, m)

		params := NfeParams{}
		nfes, err := noov.GetNfe(params)
		ast.NoError(err)
		ast.NotEmpty(nfes.Data)

		ast.Equal(float32(3.1), nfes.Data[0].NfeProc.Version)
		ast.False(nfes.Data[0].NfeProc.NFe.InfNfe.Ide.DHEmi.Valid)
	}
}

func TestNoov_GetNfeDet(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	files := []string{"fixture-4", "fixture-5"}

	for _, file := range files {
		m := make(map[string]interface{})
		fixture, _ := readFixture(fmt.Sprintf("fixtures/nfe/%s.json", file))
		json.Unmarshal(fixture, &m)

		url := getNfeUrl(noov)
		registerNfeGetResponder(ast, url, noov.Token, 200, m)

		params := NfeParams{}
		nfes, err := noov.GetNfe(params)
		ast.NoError(err)
		ast.NotEmpty(nfes.Data)

		ast.Equal(float32(3.1), nfes.Data[0].NfeProc.Version)
		ast.NotEmpty(nfes.Data[0].NfeProc.NFe.InfNfe.Det)

		ide := nfes.Data[0].NfeProc.NFe.InfNfe.Ide
		ast.True(ide.DHEmi.Valid)
	}
}

func TestNoov_GetNfeVol(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	files := []string{"fixture-6"}

	for _, file := range files {
		m := make(map[string]interface{})
		fixture, _ := readFixture(fmt.Sprintf("fixtures/nfe/%s.json", file))
		json.Unmarshal(fixture, &m)

		url := getNfeUrl(noov)
		registerNfeGetResponder(ast, url, noov.Token, 200, m)

		params := NfeParams{}
		nfes, err := noov.GetNfe(params)
		ast.NoError(err)
		ast.NotEmpty(nfes.Data)

		ast.Equal(NoovString("608"), nfes.Data[0].NfeProc.NFe.InfNfe.Transp.Vol.Marca)
		ast.Equal(NoovString(""), nfes.Data[0].NfeProc.NFe.InfNfe.Transp.Vol.QVol)
		ast.Equal(json.Number("12000"), nfes.Data[0].NfeProc.NFe.InfNfe.Transp.Vol.PesoL)

		ast.Equal(NoovString("6001"), nfes.Data[1].NfeProc.NFe.InfNfe.Transp.Vol.Marca)
		ast.Equal(NoovString("270"), nfes.Data[1].NfeProc.NFe.InfNfe.Transp.Vol.QVol)
		ast.Equal(json.Number(""), nfes.Data[1].NfeProc.NFe.InfNfe.Transp.Vol.PesoL)
	}
}

func registerNfeGetResponder(ast *assert.Assertions, url, token string, status int, m map[string]interface{}) {
	responder := func(req *http.Request) (*http.Response, error) {
		ast.Regexp(noovUrl, "http://"+req.URL.Host)

		ast.Equal("application/json", req.Header.Get("Content-Type"))
		ast.Equal("Bearer "+token, req.Header.Get("Authorization"))

		return httpmock.NewJsonResponse(status, m)
	}
	httpmock.RegisterResponder(http.MethodPost, url, responder)
}

func readFixture(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
