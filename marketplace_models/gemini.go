package marketplace_models

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func generateContentWithImage(apiKey, imagePath, text string) (string, error) {
	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %v", err)
	}

	// Encode the image to base64
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Create the JSON payload
	jsonPayload := fmt.Sprintf(`{
        "contents": [
            {
                "parts": [
                    {"text": "%s"},
                    {
                        "inlineData": {
                            "mimeType": "image/png",
                            "data": "%s"
                        }
                    }
                ]
            }
        ]
    }`, text, encodedImage)

	// Create the request
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=%s", apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonPayload)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

func GenerateText(apiKey, text string) (string, error) {
	// Create the JSON payload
	jsonPayload := fmt.Sprintf(`{
        "contents": [
            {
                "parts": [
                    {"text": "%s"},
                ]
            }
        ]
    }`, text)

	// Create the request
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=%s", apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonPayload)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

//func main() {
//	startTime := time.Now()
//	// Replace with your API key and image path
//	apiKey := "AIzaSyBHENaVP_KEfM7Bm0fuLAfxllJ8MGECpms"
//	imagePath := "img.png"
//	text := "Write about the image"
//
//	// Create a channel to receive the response
//	responseChan := make(chan struct {
//		response string
//		err      error
//	})
//
//	// Call generateText asynchronously
//	go func() {
//		response, err := generateContentWithImage(apiKey, imagePath, text)
//		responseChan <- struct {
//			response string
//			err      error
//		}{response, err}
//	}()
//
//	// Wait for the response
//	result := <-responseChan
//	if result.err != nil {
//		fmt.Println("Error:", result.err)
//		os.Exit(1)
//	}
//
//	endTime := time.Now()
//
//	fmt.Println(result.response)
//	fmt.Println("Time taken:", endTime.Sub(startTime))
//}
