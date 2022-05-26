package main

// IncomingSMS message from short-code
type IncomingSMS struct {
	ID     string `json:"id,omitempty" form:"id"`
	LinkID string `json:"linkId,omitempty" form:"linkId"`
	Text   string `json:"text,omitempty" form:"text"`
	To     string `json:"to,omitempty" form:"to"`
	From   string `json:"from,omitempty" form:"from"`
}
