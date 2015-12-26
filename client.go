package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GDClient struct {
	Url string
}

func (this *GDClient) CheckAuthorization() (bool, error) {
	resp, err := http.Post(this.Url+"/api/v1/auth/check", "application/json", strings.NewReader(""))
	if err != nil {
		return false, err
	} else {
		cnt, err := ioutil.ReadAll(resp.Body)
		log.Printf("respond:%v\n", string(cnt))
		return true, err
	}
}
