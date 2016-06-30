package gonoov

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNoov_Authenticate(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	url := getLoginUrl(noov)
	registerAuthenticateResponder(assert, url, "POST", 200)
	err := noov.Authenticate()

	assert.NoError(err)
	assert.Equal(noov.Token, "tokenTest")
}

func registerAuthenticateResponder(assert *assert.Assertions, url, method string, status int) {
	responder := func(req *http.Request) (*http.Response, error) {
		assert.Regexp(noovUrl, "http://"+req.URL.Host)

		m := make(map[string]interface{})
		json.Unmarshal(tokenTest, &m)

		assert.Equal(req.Header.Get("Content-Type"), "application/json")

		return httpmock.NewJsonResponse(status, m)
	}
	httpmock.RegisterResponder(method, url, responder)
}

// TODO Add tests to exceptions
