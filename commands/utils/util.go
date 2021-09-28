package utils

import (
	"context"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func NewContext() context.Context {
	return context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_CLIENT_APP_KEY"),
			},
		},
	)
}
