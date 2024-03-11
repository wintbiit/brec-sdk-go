package main

import (
	"net/http"
)

type BrecClient struct {
	*Client
}

func NewBrecClient(endpoint string) (*BrecClient, error) {
	c, err := NewClient(endpoint)
	if err != nil {
		return nil, err
	}

	return &BrecClient{
		Client: c,
	}, nil
}

func NewBrecClientWithAuth(endpoint, username, password string) (*BrecClient, error) {
	doer := &httpDoerWithAuth{
		Client:   http.DefaultClient,
		username: username,
		password: password,
	}

	c, err := NewClient(endpoint, WithHTTPClient(doer))
	if err != nil {
		return nil, err
	}

	return &BrecClient{
		Client: c,
	}, nil
}

type httpDoerWithAuth struct {
	*http.Client
	username string
	password string
}

func (d *httpDoerWithAuth) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(d.username, d.password)
	return d.Client.Do(req)
}
