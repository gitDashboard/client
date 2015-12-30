package client

import (
	"bytes"
	"encoding/json"
	"github.com/gitDashboard/client/v1/request"
	"github.com/gitDashboard/client/v1/response"
	"io/ioutil"
	"net/http"
)

type GDClient struct {
	Url string
}

func (this *GDClient) CheckAuthorization(username, repoPath, refName, operation string) (bool, error) {
	req := &request.AuthorizationRequest{Username: username, RepositoryPath: repoPath, RefName: refName, Operation: operation}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return false, err
	}
	resp, err := http.Post(this.Url+"/api/v1/auth/check", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return false, err
	} else {
		respCnt, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		resp := new(response.AuthorizationResponse)
		err = json.Unmarshal(respCnt, resp)
		if err != nil {
			return false, err
		}
		return resp.Authorized, err
	}
}
