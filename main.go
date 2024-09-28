package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io"
)

// SlackCommandRequest はSlackのSlashコマンドリクエストの構造体
type SlackCommandRequest struct {
	Token          string `json:"token"`
	TeamID         string `json:"team_id"`
	TeamDomain     string `json:"team_domain"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TriggerID      string `json:"trigger_id"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディを読み込む
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// URLエンコードされたリクエストボディをパース
	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusInternalServerError)
		return
	}

	// パースしたデータをSlackCommandRequest構造体にマッピング
	slackRequest := SlackCommandRequest{
		// Token:       values.Get("token"),
		// TeamID:      values.Get("team_id"),
		// TeamDomain:  values.Get("team_domain"),
		// ChannelID:   values.Get("channel_id"),
		// ChannelName: values.Get("channel_name"),
		UserID:      values.Get("user_id"),
		// UserName:    values.Get("user_name"),
		// Command:     values.Get("command"),
		// Text:        values.Get("text"),
		// ResponseURL: values.Get("response_url"),
		// TriggerID:   values.Get("trigger_id"),
	}

	// user_idがU01MWGUL5SSの時だけhello worldを返す
	if slackRequest.UserID == "U01MWGUL5SS" {
		w.Write([]byte("Hello, World!"))
	} else {
		w.Write([]byte("Invalid UserID"))
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
