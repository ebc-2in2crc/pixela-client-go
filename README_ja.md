[English](README.md) | [日本語](README_ja.md)

# Pixela

Go 用の [Pixela](https://pixe.la/) API クライアントです。

## インストール

```
$ go get -u github.com/ebc-2in2crc/pixela-client-go
```

## 使い方

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

## コントリビューション

1. このリポジトリをフォークします
2. issue ブランチを作成します (`git checkout -b issue/:id`)
3. コードを変更します
4. `make test` でテストを実行し, パスすることを確認します
5. `make fmt` でコードをフォーマットします
6. 変更をコミットします (`git commit -am 'Add some feature'`)
7. 新しいプルリクエストを作成します

## ライセンス

[MIT](https://github.com/ebc-2in2crc/pixela-client-go/blob/master/LICENSE)

## 作者

[ebc-2in2crc](https://github.com/ebc-2in2crc)
