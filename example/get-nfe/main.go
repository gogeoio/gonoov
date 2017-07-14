package main

import (
	"io/ioutil"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/robertogyn19/gonoov"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apikey    = kingpin.Flag("apikey", "apikey").Short('a').Required().String()
	apiSecret = kingpin.Flag("api-secret", "api-secret").Short('s').Required().String()
	appname   = kingpin.Flag("appname", "appname").Short('n').Required().String()
	email     = kingpin.Flag("email", "email").Short('e').Required().String()
	startDate = kingpin.Flag("start-date", "start-date (Format: DD/MM/YYYY)").Short('d').Required().String()
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

	startDateTime, err := time.Parse("01/02/2006", *startDate)

	if err != nil {
		log.Fatalf("could not parse date: %v", err)
	}

	params := gonoov.NfeParams{
		Model:      []string{"55"},
		EStartDate: startDateTime.Unix(),
		ECnpj:      *cnpjs,
		Size:       1,
	}

	nferesp, err := noov.GetNfe(params)

	if err != nil {
		log.Fatalf("could not get nfe: %v", err)
	}

	if len(nferesp.Data) > 0 {
		log.Printf("First item: %v", nferesp.Data[0].Enrichment)
	}

	ioutil.WriteFile("output.json", nferesp.Raw, 0664)
}
