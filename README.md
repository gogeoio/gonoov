## Golang Noov Client

### Usage

``` golang
package main

import (
	"encoding/json"
	"fmt"

	"github.com/robertogyn19/gonoov"
)

const (
	appname   = "AppName"
	email     = "Email"
	apikey    = "apikey"
	apisecret = "apisecret"
)

func main() {
	lparams := gonoov.LoginParams{apikey, apisecret, appname, email}
	noov := gonoov.NewNoov(lparams)
	noov.Authenticate()

	params := gonoov.NfeParams{}
	params.Size = 1
	params.Model = []string{"55"}
	params.DCnpj = []string{"01234567890"}
	params.Pagination.PageSize = 1

    p, _ := json.Marshal(params)
	fmt.Println()
	fmt.Println("--> nfe params", string(p))
	fmt.Println()

	nfeResponse, err := noov.GetNfe(params)

	if err != nil {
		fmt.Println("Error in get nfe", err)
	}

	body, _ := json.Marshal(nfeResponse.Data)
	pag, _ := json.Marshal(nfeResponse.Pagination)

	fmt.Println()
	fmt.Println("--> response  ", string(body))
	fmt.Println("--> pagination", string(pag))
	fmt.Println()
}

```