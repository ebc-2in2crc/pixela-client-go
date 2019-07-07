[English](README.md) | [日本語](README_ja.md)

# pixela-client-go

[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/pixela-client-go?status.svg)](https://godoc.org/github.com/ebc-2in2crc/pixela-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/pixela-client-go)](https://goreportcard.com/report/github.com/ebc-2in2crc/pixela-client-go)

[Pixela](https://pixe.la/) API client for Go.

## Documentation

https://godoc.org/github.com/ebc-2in2crc/pixela-client-go

## Installation

```
$ go get -u github.com/ebc-2in2crc/pixela-client-go
```

## Usage

```go
package main

import (
	"log"
	
	"github.com/ebc-2in2crc/pixela-client-go"
)

func main() {
	client := pixela.NewClient("YOUR_NAME", "YOUR_TOKEN")

	// create new user
	result, err := client.CreateUser(true, true, "")
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// create new graph
	result, err = client.Graph("graph-id").Create(
		"graph-name",
		"commit",
		pixela.TypeInt,
		pixela.ColorShibafu,
		"Asia/Tokyo",
		pixela.SelfSufficientNone,
		true,
	)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// register value
	result, err = client.Pixel("graph-id").Create("20180915", "5")
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// increment value
	result, err = client.Pixel("graph-id").Increment()
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// create new webhook
	webhook, err := client.Webhook().Create("graph-id", pixela.SelfSufficientIncrement)
	if err != nil {
		log.Fatal(err)
	}
	if webhook.IsSuccess == false {
		log.Fatal(webhook.Message)
	}

	// invoke webhook
	result, err = client.Webhook().Invoke(webhook.WebhookHash)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}
}
```

## Contribution

1. Fork this repository
2. Create your issue branch (`git checkout -b issue/:id`)
3. Change codes
4. Run test suite with the `make test` command and confirm that it passes
5. Run `make fmt`
6. Commit your changes (`git commit -am 'Add some feature'`)
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/pixela-client-go/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
