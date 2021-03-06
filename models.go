package main

// IncomingSMS message from short-code
type IncomingSMS struct {
	ID     string `json:"id,omitempty" form:"id"`
	LinkID string `json:"linkId,omitempty" form:"linkId"`
	Text   string `json:"text,omitempty" form:"text"`
	To     string `json:"to,omitempty" form:"to"`
	From   string `json:"from,omitempty" form:"from"`
}

// SMS output
type SMS struct {
	Mobile    string `json:"mobile,omitempty"`
	Text      string `json:"text,omitempty"`
	ShortCode string `json:"shortCode,omitempty"`
}

// Choice ...
type Choice struct {
	Text string `json:"text,omitempty"`
}

// OpenAIResponse ...
type OpenAIResponse struct {
	ID      string   `json:"id,omitempty"`
	Choices []Choice `json:"choices,omitempty"`
}
