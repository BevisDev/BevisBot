package service

import (
	"context"

	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAIService interface {
	Reply(ctx context.Context, msg string) (string, error)
}

type openAIService struct {
	cfg    config.OpenAI
	client openai.Client
}

func NewOpenAIService(
	cfg config.OpenAI,
) (OpenAIService, error) {
	client := openai.NewClient(
		option.WithAPIKey(cfg.APIKey),
	)
	return &openAIService{
		cfg:    cfg,
		client: client,
	}, nil
}

func (o *openAIService) Reply(ctx context.Context, msg string) (string, error) {
	resp, err := o.client.Responses.New(
		ctx,
		responses.ResponseNewParams{
			Model: "gpt-4.1-mini",
			Input: responses.ResponseNewParamsInputUnion{
				OfString: openai.String(msg),
			},
		},
	)
	return resp.OutputText(), err
}
