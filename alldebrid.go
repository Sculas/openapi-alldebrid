package openapi_alldebrid

import (
	"fmt"

	"github.com/sculas/openapi-alldebrid/client"
)

// Alldebrid is a client for the Alldebrid API.
type Alldebrid struct {
	c *client.APIClient
}

// New creates a new Alldebrid client. The API key is optional.
func New(apiKey string) *Alldebrid {
	cfg := client.NewConfiguration()
	cfg.UserAgent = "github.com/sculas/openapi-alldebrid"
	if apiKey != "" {
		cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}

	return &Alldebrid{c: client.NewAPIClient(cfg)}
}

// Service returns the default API service.
func (a *Alldebrid) Service() *client.DefaultAPIService {
	return a.c.DefaultAPI
}
