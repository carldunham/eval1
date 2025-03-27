package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type ExtractedData struct {
	VitalSigns struct {
		BloodPressure    *string  `json:"bloodPressure"`
		HeartRate        *int     `json:"heartRate"`
		Temperature      *float64 `json:"temperature"`
		RespiratoryRate  *int     `json:"respiratoryRate"`
		OxygenSaturation *int     `json:"oxygenSaturation"`
		BloodSugar       *int     `json:"bloodSugar"`
	} `json:"vitalSigns"`
	OasisElements struct {
		M0069 *string `json:"m0069"` // Patient's Living Situation
		M0102 *string `json:"m0102"` // Patient's Primary Language
		M0110 *string `json:"m0110"` // Patient's Ethnicity
		M0140 *string `json:"m0140"` // Patient's Race
		M0150 *string `json:"m0150"` // Patient's Current Health Status
		M1030 *string `json:"m1030"` // Risk for Hospitalization
		M1033 *string `json:"m1033"` // Risk for Death
		M1034 *string `json:"m1034"` // Risk for Pressure Ulcer/Injury
		M1036 *string `json:"m1036"` // Risk for Falls
		M1040 *string `json:"m1040"` // Risk for Depression
		M1046 *string `json:"m1046"` // Risk for Weight Loss
		M1051 *string `json:"m1051"` // Risk for Pressure Ulcer/Injury
		M1056 *string `json:"m1056"` // Risk for Falls
		M1058 *string `json:"m1058"` // Risk for Depression
		M1060 *string `json:"m1060"` // Risk for Weight Loss
	} `json:"oasisElements"`
	VisitType *string `json:"visitType"`
	Summary   *string `json:"summary"`
}

func NewClient() (*Client, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) ExtractData(ctx context.Context, transcript string) (*ExtractedData, error) {
	prompt := `Analyze the following home health visit transcript and extract vital signs, OASIS elements, and visit type.
Format the response as a JSON object with the following structure:
{
  "vitalSigns": {
    "bloodPressure": "string or null",
    "heartRate": "number or null",
    "temperature": "number or null",
    "respiratoryRate": "number or null",
    "oxygenSaturation": "number or null",
    "bloodSugar": "number or null"
  },
  "oasisElements": {
    "m0069": "string or null (Patient's Living Situation)",
    "m0102": "string or null (Patient's Primary Language)",
    "m0110": "string or null (Patient's Ethnicity)",
    "m0140": "string or null (Patient's Race)",
    "m0150": "string or null (Patient's Current Health Status)",
    "m1030": "string or null (Risk for Hospitalization)",
    "m1033": "string or null (Risk for Death)",
    "m1034": "string or null (Risk for Pressure Ulcer/Injury)",
    "m1036": "string or null (Risk for Falls)",
    "m1040": "string or null (Risk for Depression)",
    "m1046": "string or null (Risk for Weight Loss)",
    "m1051": "string or null (Risk for Pressure Ulcer/Injury)",
    "m1056": "string or null (Risk for Falls)",
    "m1058": "string or null (Risk for Depression)",
    "m1060": "string or null (Risk for Weight Loss)"
  },
  "visitType": "string or null (one of: 'SOC', 'Follow-up', 'Recertification', 'Discharge')",
  "summary": "string (A concise summary of the visit, including key findings, concerns, and recommendations based on the vital signs and OASIS elements)"
}

Transcript:
` + transcript

	reqBody := Request{
		Model: "gpt-4-0125-preview",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a medical data extraction assistant. Extract vital signs and OASIS elements from the transcript. Return only valid JSON.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var apiResp Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding API response: %w", err)
	}

	if len(apiResp.Choices) == 0 {
		return nil, fmt.Errorf("no response from API")
	}

	var extractedData ExtractedData
	content := apiResp.Choices[0].Message.Content

	// Strip markdown code blocks if present
	if len(content) >= 7 && content[:7] == "```json" {
		content = content[7:]
	}
	if len(content) >= 3 && content[len(content)-3:] == "```" {
		content = content[:len(content)-3]
	}
	content = strings.TrimSpace(content)

	if err := json.Unmarshal([]byte(content), &extractedData); err != nil {
		return nil, fmt.Errorf("error parsing extracted data: %w\nContent received: %s", err, content)
	}

	return &extractedData, nil
}
