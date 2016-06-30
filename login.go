package gonoov

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (noov *Noov) Authenticate() error {
	hash, ts := createHash(noov)
	data, err := noov.login(hash, ts)

	if err != nil {
		return err
	}

	token := token{}
	err = json.Unmarshal(data, &token)

	noov.Token = string(token.Token)
	return err
}

func (noov *Noov) login(secret string, ts int64) ([]byte, error) {
	p := noovLoginParams{noov.ApiKey, ts, secret}
	d, _ := json.Marshal(p)

	url := getLoginUrl(noov)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(d))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return []byte{}, err
	}

	resp, err := noov.client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}

func getLoginUrl(noov *Noov) string {
	return fmt.Sprintf("%s/%s/%s", noov.url, noov.version, loginUrl)
}

func computeHmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func createHash(noov *Noov) (string, int64) {
	ts := time.Now().UnixNano() / int64(time.Millisecond)
	str := fmt.Sprintf("%s%s%d", noov.appname, noov.email, ts)
	hash := computeHmac256(str, noov.ApiSecret)

	return hash, ts
}
