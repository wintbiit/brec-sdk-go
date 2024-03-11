//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package main -generate types,client -o client.gen.go swagger.json

package main

import (
	"encoding/json"
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

func DecodeResponse(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
