package gonoov

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	nfeUrl = "app/nfe"
)

func (noov *Noov) GetNfe(params NfeParams) (NfeRawResponse, error) {
	rresp := NfeRawResponse{}

	b, err := json.Marshal(params)

	if err != nil {
		return rresp, err
	}

	url := getNfeUrl(noov)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))

	// EXPLAIN: Tenta autenticar novamente. Este só será chamado quando o timestamp já tiver com a "idade" de pelo menos 25 minutos
	noov.Authenticate()
	setRequestHeaders(req, noov.Token)
	resp, err := noov.client.Do(req)

	if err != nil || resp.StatusCode == 503 {
		if retryError(resp, err) {
			fmt.Println("Timeout error. Trying again...")
			count := 0
			for !retryError(resp, err) {
				resp, err = noov.client.Do(req)
				if err != nil {
					if retryError(resp, err) {
						time.Sleep(5 * time.Second)
						count++
					} else {
						fmt.Println("Erro desconhecido:", err)
						body, _ := ioutil.ReadAll(resp.Body)
						fmt.Println("Dado bruto", string(body))
						return rresp, err
					}
				}

				if count > 10 {
					if resp != nil {
						fmt.Println("Status code", resp.StatusCode)
					}

					fmt.Printf("Atingido limite máximo de retries: %d\n", count)
					break
				}
			}
		} else {
			if resp != nil {
				fmt.Println("Status code", resp.StatusCode)
			}
			return rresp, err
		}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return rresp, err
	}

	rresp.Raw = body
	err = json.Unmarshal(body, &rresp)

	if err != nil {
		return rresp, err
	}

	if len(rresp.Meta.Errors) > 0 {
		err := fmt.Sprintf("%v", rresp.Meta.Errors)
		return rresp, errors.New(err)
	}

	return rresp, err
}

func retryError(resp *http.Response, err error) bool {
	return err == http.ErrHandlerTimeout || (resp != nil && resp.StatusCode == http.StatusServiceUnavailable)
}

func getNfeUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, nfeUrl)
}

func setRequestHeaders(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")

	auth := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", auth)
}
