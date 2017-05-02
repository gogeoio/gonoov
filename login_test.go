package gonoov

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

var (
	tokenTest  = []byte(`{"token":"token1"}`)
	tokenTest2 = []byte(`{"token":"token2"}`)
)

func TestNoov_Authenticate(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	url := getLoginUrl(noov)
	registerAuthenticateResponder(ast, tokenTest, url, "POST", 200)
	err := noov.Authenticate()

	ast.NoError(err)
	ast.Equal(noov.Token, "token1")
}

func TestNoov_AutoUpdateToken(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	url := getLoginUrl(noov)

	registerAuthenticateResponder(ast, tokenTest, url, "POST", 200)
	err := noov.Authenticate()

	ast.NoError(err)
	ast.Equal(noov.Token, "token1")

	noov.TokenTimestamp = noov.TokenTimestamp - 1800000
	registerAuthenticateResponder(ast, tokenTest2, url, "POST", 200)
	err = noov.Authenticate()

	ast.NoError(err)
	ast.Equal(noov.Token, "token2")
}

func TestNoov_ReuseToken(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	url := getLoginUrl(noov)

	registerAuthenticateResponder(ast, tokenTest, url, "POST", 200)
	err := noov.Authenticate()

	ast.NoError(err)
	ast.Equal(noov.Token, "token1")

	ts1 := noov.TokenTimestamp
	err = noov.Authenticate()

	ast.NoError(err)
	ast.Equal(noov.Token, "token1")
	ast.Equal(noov.TokenTimestamp, ts1)
}

func registerAuthenticateResponder(assert *assert.Assertions, token []byte, url, method string, status int) {
	responder := func(req *http.Request) (*http.Response, error) {
		assert.Regexp(noovUrl, "http://"+req.URL.Host)

		m := make(map[string]interface{})
		json.Unmarshal(token, &m)

		assert.Equal(req.Header.Get("Content-Type"), "application/json")

		return httpmock.NewJsonResponse(status, m)
	}
	httpmock.RegisterResponder(method, url, responder)
}

// TODO Add tests to exceptions
