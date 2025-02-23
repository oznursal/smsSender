package model

type SmsRequest struct {
	To      string `json:"to"`
	Content string `json:"content"`
}
