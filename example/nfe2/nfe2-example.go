package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/robertogyn19/gonoov"

	"gopkg.in/alecthomas/kingpin.v2"
)

const inputDateLayout = "02/01/2006"

var (
	apikey         = kingpin.Flag("apikey", "apikey").Short('a').Required().String()
	apiSecret      = kingpin.Flag("api-secret", "api-secret").Short('s').Required().String()
	appname        = kingpin.Flag("appname", "appname").Short('n').Required().String()
	email          = kingpin.Flag("email", "email").Short('e').Required().String()
	startDate      = kingpin.Flag("start-date", "start-date (Format: DD/MM/YYYY)").Short('d').Required().String()
	endDate        = kingpin.Flag("end-date", "end-date (Format: DD/MM/YYYY)").Short('t').Required().String()
	cnpjs          = kingpin.Flag("cnpj", "cnpj").Short('c').Required().Strings()
	printFirstItem = kingpin.Flag("print", "print first item").Short('p').Bool()
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

	startDateTime, err := time.Parse(inputDateLayout, *startDate)
	endDateTime, err2 := time.Parse(inputDateLayout, *endDate)

	if err != nil || err2 != nil {
		log.Fatalf("could not parse date: %v", err)
	}

	for {
		params := gonoov.NfeParams{
			Model:      []string{"55"},
			EStartDate: startDateTime.Add(2 * time.Hour).Unix() * 1000,
			EEndDate:   startDateTime.Add(2 * time.Hour).AddDate(0, 0, 1).Unix() * 1000,
			ECnpj:      *cnpjs,
			Size:       150,
		}
		pp, _ := json.Marshal(params)
		log.Printf("params: %s", pp)

		nferesp, err := noov.GetNfe(params)

		if err != nil {
			log.Fatalf("could not get nfe: %v %s", err, nferesp.Raw)
		}

		if len(nferesp.Data) > 0 && *printFirstItem {
			log.Printf("First item: %v", nferesp.Data[0].Enrichment)
			ioutil.WriteFile("output.json", nferesp.Raw, 0664)
		}

		log.Printf("%s: %d", startDateTime.Format(inputDateLayout), len(nferesp.Data))

		startDateTime = startDateTime.AddDate(0, 0, 1)

		if startDateTime.After(endDateTime) {
			break
		}
	}
}
