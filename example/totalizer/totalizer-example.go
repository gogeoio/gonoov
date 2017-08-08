package main

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/robertogyn19/gonoov"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	dateLayout = "2006-01-02"
)

var (
	apikey         = kingpin.Flag("apikey", "apikey").Short('a').Required().String()
	apiSecret      = kingpin.Flag("api-secret", "api-secret").Short('s').Required().String()
	appname        = kingpin.Flag("appname", "appname").Short('n').Required().String()
	email          = kingpin.Flag("email", "email").Short('e').Required().String()
	startDate      = kingpin.Flag("start-date", "start-date (Format: YYYY-MM-DD)").Required().Short('d').String()
	endDate        = kingpin.Flag("end-date", "end-date (Format: YYYY-MM-DD)").Required().Short('t').String()
	cnpjs          = kingpin.Flag("cnpjs", "cnpjs").Short('c').Required().Strings()
	nfeKey         = kingpin.Flag("nfe-key", "nfe-key").Short('k').String()
	printTotalizer = kingpin.Flag("print", "print totalizers").Short('p').Bool()
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

	currentDate := *startDate
	startDateTime, err := time.Parse(dateLayout, currentDate)
	endDateTime, err2 := time.Parse(dateLayout, *endDate)

	if startDateTime.IsZero() || err != nil || err2 != nil {
		log.Fatalf("invalid date: %s, error: %v", currentDate, err)
	}

	for {
		totalizerParams := gonoov.TotalizerParams{
			Day:   currentDate,
			Emits: *cnpjs,
		}

		if nfeKey != nil {
			totalizerParams.NFEKey = *nfeKey
		}

		totalizer, err := noov.Totalizer(totalizerParams)

		if err != nil {
			log.Fatalf("could not get totalizer from noov: %v", err)
		}

		if *printTotalizer {
			d, _ := json.MarshalIndent(totalizer.Totalizers, "", "  ")
			log.Printf("%s", d)
		}
		log.Printf("%s: %d", currentDate, len(totalizer.Totalizers))

		startDateTime = startDateTime.AddDate(0, 0, 1)
		currentDate = startDateTime.Format(dateLayout)

		if startDateTime.After(endDateTime) {
			break
		}
	}

	//log.Printf("raw: %s", totalizer.Raw)
}
