package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/evidently"
)

type EvidentlyClient struct {
	client *evidently.Client
}

type NewClientInput struct {
	Context context.Context
	Region  string
}

func NewEvidentlyClient(in *NewClientInput) (*EvidentlyClient, error) {
	cfg, err := config.LoadDefaultConfig(in.Context,
		config.WithRegion(in.Region),
	)
	if err != nil {
		return nil, err
	}

	c := evidently.NewFromConfig(cfg)
	return &EvidentlyClient{client: c}, nil
}

func (e *EvidentlyClient) ListFeatures(ctx context.Context, project string) (*evidently.ListFeaturesOutput, error) {
	in := &evidently.ListFeaturesInput{
		Project: aws.String(project),
	}
	return e.client.ListFeatures(ctx, in)
}
