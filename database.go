package main

import (
	"context"
	"fmt"
)

// GetPrompt ...
func GetPrompt(ctx context.Context, parentPhone string) (string, error) {
	var childName, parentName, departureTime, numberPlate, driverName string
	layout := "%s is the the child of %s. %s has left school at %s. %s is in bus with number plate %s with driver %s. %s is using route %s. %s is arriving at home at %s\n%s: When is he coming?\nChat bot:"

	// TODO create sql query
	childName, parentName, departureTime, numberPlate, driverName = "Jackson Juma", "Jane Juma", "01.00pm", "KYC 445L", "Wycliffe"

	prompt := fmt.Sprintf(layout, childName, parentName, departureTime, childName, numberPlate, childName, driverName, parentName)
	return prompt, nil
}
