package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// ExampleFetch demonstrates Resty usage
func ExampleFetch() {
	client := resty.New()

	resp, err := client.R().
		Get("https://httpbin.org/get")

	if err != nil {
		log.Error().Err(err).Msg("HTTP request failed")
		return
	}

	log.Info().
		Int("status", resp.StatusCode()).
		Str("body", string(resp.Body()[:min(len(resp.Body()), 50)])).
		Msg("HTTP request succeeded")
}
