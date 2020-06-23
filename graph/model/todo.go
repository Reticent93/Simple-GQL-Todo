package model

type Todo struct {
	ID string `json:"id""`
	Text string `json:"text"`
	Done bool `json:"done"`
	User string `json:"user"`
}
