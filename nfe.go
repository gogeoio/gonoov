package gonoov

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	setRequestHeaders(req, noov.Token)
	resp, err := noov.client.Do(req)

	if err != nil {
		return rresp, err
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

func getNfeUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, nfeUrl)
}

func setRequestHeaders(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")

	auth := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", auth)
}
