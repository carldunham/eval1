package openai

import "context"

// ClientInterface defines the interface for OpenAI client operations
type ClientInterface interface {
	ExtractData(ctx context.Context, transcript string) (*ExtractedData, error)
}
