package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const (
	modelID = "anthropic.claude-instant-v1"
	region  = "ap-northeast-1"
)

type AnthropicRequest struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"top_p"`
	TopK              int      `json:"top_k"`
	StopSequences     []string `json:"stop_sequences"`
}

type AnthropicResponse struct {
	Completion string `json:"completion"`
}

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Printf("Error loading AWS configuration: %v\n", err)
		return
	}

	client := bedrockruntime.NewFromConfig(cfg)

	fmt.Println("How can I help you?")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}

		response, err := callBedrock(client, input)
		if err != nil {
			fmt.Printf("Error calling Bedrock: %v\n", err)
			continue
		}

		fmt.Println(strings.TrimSpace(response))
		fmt.Println("\nHow can I help you?")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

func callBedrock(client *bedrockruntime.Client, input string) (string, error) {
	prompt := fmt.Sprintf("\n\nHuman: %s\n\nAssistant:", input)

	request := AnthropicRequest{
		Prompt:            prompt,
		MaxTokensToSample: 300,
		Temperature:       0.7,
		TopP:              1,
		TopK:              250,
		StopSequences:     []string{"\n\nHuman:"},
	}

	jsonPayload, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	output, err := client.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
		Body:        jsonPayload,
		ModelId:     aws.String(modelID),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return "", fmt.Errorf("error invoking model: %w", err)
	}

	var response AnthropicResponse
	err = json.Unmarshal(output.Body, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	return response.Completion, nil
}
