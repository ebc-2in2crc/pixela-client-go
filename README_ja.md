[English](README.md) | [日本語](README_ja.md)

# pixela-client-go

[![Build Status](https://travis-ci.com/ebc-2in2crc/pixela-client-go.svg?branch=master)](https://travis-ci.com/ebc-2in2crc/pixela-client-go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/pixela-client-go?status.svg)](https://godoc.org/github.com/ebc-2in2crc/pixela-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/pixela-client-go)](https://goreportcard.com/report/github.com/ebc-2in2crc/pixela-client-go)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/pixela-client-go.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/pixela-client-go.svg?label=version)

Go 用の [Pixela](https://pixe.la/) API クライアントです。

[![Cloning count](https://pixe.la/v1/users/ebc-2in2crc/graphs/p-c-g-clone)](https://pixe.la/v1/users/ebc-2in2crc/graphs/p-c-g-clone.html)

## ドキュメント

https://godoc.org/github.com/ebc-2in2crc/pixela-client-go

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

	// 新しいユーザーを作る
	result, err := client.CreateUser(true, true, "")
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しい slack チャンネル を作る
	detail := &pixela.SlackDetail{
		URL:         "https://hooks.slack.com/services/xxxx",
		UserName:    "slack-user-name",
		ChannelName: "slack-channel-name",
	}
	result, err = client.Channel().CreateSlackChannel("channel-id", "channel-name", detail)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しいグラフを作る
	result, err = client.Graph("graph-id").Create(
		"graph-name",
		"commit",
		pixela.TypeInt,
		pixela.ColorShibafu,
		"Asia/Tokyo",
		pixela.SelfSufficientNone,
		false,
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 値をピクセルに記録する
	result, err = client.Pixel("graph-id").Create("20180915", "5", "")
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// ピクセルの値をインクリメントする
	result, err = client.Pixel("graph-id").Increment()
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しい通知ルールを作る
	result, err = client.Notification("graph-id").Create(
		"notification-id",
		"notification-name",
		pixela.TargetQuantity,
		pixela.ConditionGreaterThan,
		"3",
		"channel-id",
	)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しい webhook を作る
	webhook, err := client.Webhook().Create("graph-id", pixela.SelfSufficientIncrement)
	if err != nil {
		log.Fatal(err)
	}
	if webhook.IsSuccess == false {
		log.Fatal(webhook.Message)
	}

	// webhook を呼び出す
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
