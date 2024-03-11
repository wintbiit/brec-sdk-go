# GO SDK for [BililiveRecorder](https://rec.danmuji.org/)
![Brec Version](https://img.shields.io/badge/version-2.11.0-blue)

## Update API Specification
1. Alter `swagger.json` file to new version.
2. Run command:
```shell
go generate
```
3. Check `client.gen.go` file.

## Usage
```go
package main

import (
    "fmt"
    brec "github.com/wintbiit/brec-sdk-go"
)

func main() {
	client, _ := brec.NewBrecClient("http://localhost:2356");
	// or client, _ := brec.NewBrecClientWithAuth("http://localhost:2356", "username", "password");
}
```