package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gitDashboard/client/v1/misc"
	"github.com/gitDashboard/client/v1/request"
	"github.com/gitDashboard/client/v1/response"
	"io"
	"io/ioutil"
	"net/http"
)

type GDClient struct {
	Url string
}

func parseResponse(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	respCnt, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respCnt, v)
}

func (this *GDClient) CheckAuthorization(username, repoPath, refName, operation string) (bool, bool, error) {
	req := &request.AuthorizationRequest{Username: username, RepositoryPath: repoPath, RefName: refName, Operation: operation}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return false, false, err
	}
	httpResp, err := http.Post(this.Url+"/api/v1/auth/check", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return false, false, err
	} else {
		resp := new(response.AuthorizationResponse)
		err := parseResponse(httpResp.Body, resp)
		if err != nil {
			return false, false, err
		}
		return resp.Authorized, resp.Locked, err
	}
}

/**
return eventId
*/
func (this *GDClient) StartEvent(repoPath, eventType, user, reference, eventDescription string, level misc.EventLevel) (uint, error) {
	req := request.RepoEventRequest{
		RepositoryPath: repoPath,
		Type:           eventType,
		Level:          level,
		User:           user, Reference: reference,
		Description: eventDescription}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}
	httpResp, err := http.Post(this.Url+"/api/v1/event/start", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return 0, err
	} else {
		resp := new(response.RepoEventResponse)
		err := parseResponse(httpResp.Body, resp)
		if err != nil {
			return 0, err
		}
		if !resp.Success {
			return 0, errors.New(resp.Error.Message)
		}
		return resp.EventID, nil
	}
}

func (this *GDClient) FinishEvent(eventId uint) error {
	urlToCall := fmt.Sprintf("%s/api/v1/event/%d/finish", this.Url, eventId)
	httpResp, err := http.Get(urlToCall)
	if err != nil {
		return err
	} else {
		resp := new(response.BasicResponse)
		err := parseResponse(httpResp.Body, resp)
		if err != nil {
			return err
		}
		if !resp.Success {
			return errors.New(resp.Error.Message)
		}
		return nil
	}
}

func (this *GDClient) AddEvent(repoPath, eventType, user, reference, eventDescription string, level misc.EventLevel) (uint, error) {
	req := request.RepoEventRequest{
		RepositoryPath: repoPath,
		Type:           eventType,
		Level:          level,
		User:           user,
		Reference:      reference,
		Description:    eventDescription}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}
	httpResp, err := http.Post(this.Url+"/api/v1/event/add", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return 0, err
	} else {
		resp := new(response.RepoEventResponse)
		err := parseResponse(httpResp.Body, resp)
		if err != nil {
			return 0, err
		}
		if !resp.Success {
			return 0, errors.New(resp.Error.Message)
		}
		return resp.EventID, nil
	}
}
