package gonoov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	loginParams = LoginParams{"apikey", "apisecret", "noovTests", "ti@gogeo.io"}
)

func TestNewNoov(t *testing.T) {
	assert := assert.New(t)

	p := loginParams
	noov := NewNoov(p)

	assert.NotNil(noov)
	assert.Equal(noov.url, noovUrl)
	assert.Equal(noov.version, noovVersion)
	assert.Equal(noov.Token, "")
	assert.NotNil(noov.client)
	assert.Equal(noov.ApiKey, p.ApiKey)
	assert.Equal(noov.ApiSecret, p.ApiSecret)
	assert.Equal(noov.appname, p.AppName)
	assert.Equal(noov.email, p.Email)
}
