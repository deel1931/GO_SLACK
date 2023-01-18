package main

import "github.com/slack-go/slack"

func main() {
	// アクセストークンを使用してクライアントを生成する
	tkn := "トークン"
	c := slack.New(tkn)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage("#チャンネル名", slack.MsgOptionText("メッセージ", true))
	if err != nil {
		panic(err)
	}
}
