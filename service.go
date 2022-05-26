package main

import (
	"context"
	"fmt"
	"net/http"
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

	client := gohttp.New(time.Minute * 2)
	baseURL, err := MustGetEnv("AT_SANDBOX_URL")
	if err != nil {
		return nil, err
	}

	username, err := MustGetEnv("")
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

	status, res, err := client.Post(ctx, baseURL, headers, payload)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %v", status)
	}

	resp := string(res)

	return &resp, nil

}
