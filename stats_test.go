package gonoov

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestNoov_Stats(t *testing.T) {
	ast := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	noov := NewNoov(loginParams)
	noov.Token = "token-test"

	m := make(map[string]interface{})
	fixture, _ := readFixture("fixtures/stats/stats-1.json")
	json.Unmarshal(fixture, &m)

	url := getStatsUrl(noov)
	registerStatsResponder(ast, url, noov.Token, http.StatusOK, m)

	params := StatsParams{}
	stats, err := noov.Stats(params)
	ast.NoError(err)
	ast.Equal(239, stats.Total)
}

func registerStatsResponder(ast *assert.Assertions, url string, token string, status int, m map[string]interface{}) {
	responder := func(req *http.Request) (*http.Response, error) {
		ast.Regexp(noovUrl, "http://"+req.URL.Host)

		ast.Equal("application/json", req.Header.Get("Content-Type"))
		ast.Equal("Bearer "+token, req.Header.Get("Authorization"))

		return httpmock.NewJsonResponse(status, m)
	}

	httpmock.RegisterResponder(http.MethodPost, url, responder)
}
