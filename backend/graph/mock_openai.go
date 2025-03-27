package graph

import (
	"context"

	"github.com/carldunham/nestmed/eval1/backend/openai"
)

type TestOpenAIClient struct {
	extractDataFunc func(ctx context.Context, transcript string) (*openai.ExtractedData, error)
}

func NewTestOpenAIClient() *TestOpenAIClient {
	return &TestOpenAIClient{}
}

func (c *TestOpenAIClient) SetExtractDataFunc(f func(ctx context.Context, transcript string) (*openai.ExtractedData, error)) {
	c.extractDataFunc = f
}

func (c *TestOpenAIClient) ExtractData(ctx context.Context, transcript string) (*openai.ExtractedData, error) {
	if c.extractDataFunc != nil {
		return c.extractDataFunc(ctx, transcript)
	}
	return &openai.ExtractedData{}, nil
}
