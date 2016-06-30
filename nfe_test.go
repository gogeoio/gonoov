package gonoov

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	m := make(map[string]interface{})
	fixture, _ := readFixture("fixtures/nfe/fixture-1.json")
	json.Unmarshal(fixture, &m)

	url := getNfeUrl(noov)
	registerNfeGetResponder(assert, url, noov.Token, 200, m)

	params := NfeParams{}
	nfes, err := noov.Get(params)
	assert.NoError(err)
	assert.NotEmpty(nfes)

	assert.Equal(float32(3.1), nfes[0].NfeProc.Version)
}

func TestNoov_GetWithError(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	m := make(map[string]interface{})
	json.Unmarshal(notFoundError, &m)

	url := getNfeUrl(noov)
	registerNfeGetResponder(assert, url, noov.Token, 404, m)

	params := NfeParams{}
	nfes, err := noov.Get(params)
	assert.Error(err)
	assert.Empty(nfes)
}

func registerNfeGetResponder(assert *assert.Assertions, url, token string, status int, m map[string]interface{}) {
	responder := func(req *http.Request) (*http.Response, error) {
		assert.Regexp(noovUrl, "http://"+req.URL.Host)

		assert.Equal("application/json", req.Header.Get("Content-Type"))
		assert.Equal("Bearer "+token, req.Header.Get("Authorization"))

		return httpmock.NewJsonResponse(status, m)
	}
	httpmock.RegisterResponder("POST", url, responder)
}

func readFixture(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
