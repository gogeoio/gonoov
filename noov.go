package gonoov

import (
	"net/http"
	"time"
)

var (
	noovUrl     = "http://rest.noov.com.br"
	noovVersion = "v1"

	loginUrl = "auth/login"
)

func NewNoov(params LoginParams) *Noov {
	c := http.Client{
		Timeout: 180 * time.Second,
	}

	noov := Noov{
		ApiKey:    params.ApiKey,
		ApiSecret: params.ApiSecret,
		url:       noovUrl,
		version:   noovVersion,
		appname:   params.AppName,
		email:     params.Email,
		Token:     "",
		client:    &c,
	}

	return &noov
}
