package gemini

import (
	"context"
	"errors"
	"log"
	"os"

	"google.golang.org/genai"
)

type GeminiClient interface {
	ExtractLogisticsData(ctx context.Context, systemPrompt string, userMessage string) (string, error)
}

type geminiClient struct {
	client *genai.Client
}

func NewGeminiClient() (GeminiClient, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY belum diset di .env")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, err
	}

	log.Println("🤖 [Gemini] Client AI berhasil diinisialisasi!")
	return &geminiClient{client: client}, nil
}

func (g *geminiClient) ExtractLogisticsData(ctx context.Context, systemPrompt string, userMessage string) (string, error) {
	model := "gemini-2.5-flash"

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: systemPrompt},
			},
		},
		ResponseMIMEType: "application/json",
		Temperature:      genai.Ptr[float32](0.1),
	}

	contents := []*genai.Content{
		{
			Role: "user",
			Parts: []*genai.Part{
				{Text: userMessage},
			},
		},
	}

	resp, err := g.client.Models.GenerateContent(ctx, model, contents, config)
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		jsonOutput := resp.Candidates[0].Content.Parts[0].Text
		if jsonOutput != "" {
			return jsonOutput, nil
		}
	}

	return "", errors.New("gemini mengembalikan respons kosong")
}
