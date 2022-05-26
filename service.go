package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	gohttp "github.com/ochom/go-http"
)

// SendSMS ....
func SendSMS(ctx context.Context, sms SMS) (*string, error) {

	apiKey, err := MustGetEnv("AT_SANDBOX_KEY")
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept":       "application/json",
		"apiKey":       apiKey,
		"Content-Type": "application/x-www-form-urlencoded",
	}

	baseURL, err := MustGetEnv("AT_SANDBOX_URL")
	if err != nil {
		return nil, err
	}

	username, err := MustGetEnv("AT_USERNAME")
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Add("username", username)
	data.Add("to", sms.Mobile)
	data.Add("message", sms.Text)
	data.Add("from", sms.ShortCode)

	encoded := data.Encode()

	payload := []byte(encoded)

	client := gohttp.New(time.Minute * 2)
	status, res, err := client.Post(ctx, baseURL, headers, payload)
	if err != nil {
		return nil, err
	}

	log.Println("request status: ", status)

	resp := string(res)

	return &resp, nil
}

// QueryOpenAI ...
func QueryOpenAI(ctx context.Context, prompt string) (*OpenAIResponse, error) {
	url := "https://api.openai.com/v1/engines/text-davinci-002/completions"

	apiKey, err := MustGetEnv("OPEN_AI_KEY")
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", apiKey),
	}

	d := map[string]any{
		"prompt":            prompt,
		"temperature":       0.7,
		"max_tokens":        256,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	payload, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	client := gohttp.New(time.Minute * 2)
	_, res, err := client.Post(ctx, url, headers, payload)
	if err != nil {
		return nil, err
	}

	var data OpenAIResponse
	if err = json.Unmarshal(res, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
