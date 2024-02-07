package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ConvertedBody represents the structure of desired format
type ConvertedBody struct {
	Event           string                    `json:"event"`
	EventType       string                    `json:"event_type"`
	AppID           string                    `json:"app_id"`
	UserID          string                    `json:"user_id"`
	MessageID       string                    `json:"message_id"`
	PageTitle       string                    `json:"page_title"`
	PageURL         string                    `json:"page_url"`
	BrowserLanguage string                    `json:"browser_language"`
	ScreenSize      string                    `json:"screen_size"`
	Attributes      map[string]AttributeValue `json:"attributes"`
	Traits          map[string]TraitValue     `json:"traits"`
}

// AttributeValue represents the structure of an attribute value
type AttributeValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// TraitValue represents the structure of a trait value
type TraitValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

var webhookURL string

func main() {
	// Create a channel to send requests to the worker
	requestChannel := make(chan map[string]interface{})

	// Start the worker
	go worker(requestChannel)

	// Handle HTTP requests
	http.HandleFunc("/convert-object", func(w http.ResponseWriter, r *http.Request) {
		// Decode the incoming JSON request body
		requestBody := make(map[string]interface{})

		webhookURL = r.FormValue("webhook_url")

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send the request to the worker through the channel

		requestChannel <- requestBody

		// Respond to the client
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Request received and sent to worker successfully")
	})

	// Start the HTTP server
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func worker(requestChannel chan map[string]interface{}) {

	for request := range requestChannel {
		// Convert request to desired format
		formattedRequest := ConvertedBody{
			Event:           getString(request, "ev"),
			EventType:       getString(request, "et"),
			AppID:           getString(request, "id"),
			UserID:          getString(request, "uid"),
			MessageID:       getString(request, "mid"),
			PageTitle:       getString(request, "t"),
			PageURL:         getString(request, "p"),
			BrowserLanguage: getString(request, "l"),
			ScreenSize:      getString(request, "sc"),
			Attributes:      make(map[string]AttributeValue),
			Traits:          make(map[string]TraitValue),
		}

		// Process attributes
		processAttributes(request, &formattedRequest)

		// Process user traits
		processTraits(request, &formattedRequest)

		// Send formatted request to webhook.site
		jsonData, err := json.Marshal(formattedRequest)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}

		fmt.Println(webhookURL)
		sendToWebHook(webhookURL, jsonData)
	}

}

func sendToWebHook(webhookUrl string, data []byte) {

	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error sending request to webhook.site:", err)
	}
}

func getString(data map[string]interface{}, key string) string {
	value, ok := data[key].(string)
	if !ok {
		return ""
	}
	return value
}

func processAttributes(request map[string]interface{}, formattedRequest *ConvertedBody) {
	for i := 1; ; i++ {
		attrKey := fmt.Sprintf("atrk%d", i)
		valueKey := fmt.Sprintf("atrv%d", i)
		typeKey := fmt.Sprintf("atrt%d", i)

		attrValue := getString(request, attrKey)
		if attrValue == "" {
			break
		}

		value := getString(request, valueKey)
		attrType := getString(request, typeKey)

		formattedRequest.Attributes[attrValue] = AttributeValue{Value: value, Type: attrType}
	}
}

func processTraits(request map[string]interface{}, formattedRequest *ConvertedBody) {
	for i := 1; ; i++ {
		traitKey := fmt.Sprintf("uatrk%d", i)
		valueKey := fmt.Sprintf("uatrv%d", i)
		typeKey := fmt.Sprintf("uatrt%d", i)

		traitValue := getString(request, traitKey)
		if traitValue == "" {
			break
		}

		value := getString(request, valueKey)
		traitType := getString(request, typeKey)

		formattedRequest.Traits[traitValue] = TraitValue{Value: value, Type: traitType}
	}
}
