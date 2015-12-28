package client

import (
	"testing"
)

func TestCheck(t *testing.T) {
	cl := &GDClient{Url: "http://localhost:9000"}
	_, err := cl.CheckAuthorization("igor", "/home/igor/git/testAuth.git", "refs/heads/master", "commit")
	if err != nil {
		t.Error(err)
	}
}
