package main_test

import (
	"context"
	"encoding/json"
	brec "github.com/wintbiit/brec-sdk-go"
	"testing"
)

const (
	ENDPOINT = "http://localhost:2356"
	USERNAME = "root"
	PASSWORD = "root"
)

func TestFetchConfig(t *testing.T) {
	c, err := brec.NewBrecClientWithAuth(ENDPOINT, USERNAME, PASSWORD)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetApiConfigDefault(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	var config brec.DefaultConfig
	if err := brec.DecodeResponse(resp, &config); err != nil {
		t.Fatal(err)
	}

	serialize, _ := json.MarshalIndent(config, "", "  ")
	t.Log(string(serialize))
}
