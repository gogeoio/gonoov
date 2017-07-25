package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/robertogyn19/gonoov"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apikey    = kingpin.Flag("apikey", "apikey").Short('a').Required().String()
	apiSecret = kingpin.Flag("api-secret", "api-secret").Short('s').Required().String()
	appname   = kingpin.Flag("appname", "appname").Short('n').Required().String()
	email     = kingpin.Flag("email", "email").Short('e').Required().String()
	startDate = kingpin.Flag("start-date", "start-date (Format: YYYY-MM-DD)").Required().Short('d').String()
	cnpjs     = kingpin.Flag("cnpjs", "cnpjs").Short('c').Required().Strings()
)

func main() {
	kingpin.Parse()

	loginParams := gonoov.LoginParams{
		ApiKey:    *apikey,
		ApiSecret: *apiSecret,
		AppName:   *appname,
		Email:     *email,
	}

	noov := gonoov.NewNoov(loginParams)
	err := noov.Authenticate()

	if err != nil {
		log.Fatalf("could not authenticate to noov: %v", err)
	}

	statsParams := gonoov.StatsParams{
		ECnpj:      *cnpjs,
		EStartDate: *startDate,
	}

	stats, err := noov.Stats(statsParams)

	if err != nil {
		log.Fatalf("could not get stats from noov: %v", err)
	}

	d, _ := json.MarshalIndent(stats, "", "  ")
	log.Printf("%s", d)
}
