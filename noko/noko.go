package noko

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const (
	BaseURLV2 = "https://api.nokotime.com/v2"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient() *Client {
	apiKey := fetchApiKey()

	return &Client{
		BaseURL: BaseURLV2,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-NokoToken", c.apiKey)
	req.Header.Set("User-Agent", "Nina/v0.1")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.New("unable to sendRequest")
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) send(ctx context.Context, req *http.Request) error {
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-NokoToken", c.apiKey)
	req.Header.Set("User-Agent", "Nina/v0.1")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.New("unable to send")
	}

	return nil
}

func fetchApiKey() string {
	env_key := viper.GetString("NOKO_API_KEY")
	cnf_key := viper.GetString("api_key")

	if env_key != "" {
		return env_key
	}

	if cnf_key != "" {
		return cnf_key
	}

	log.Fatal(`
	Please set enviromental variable NOKO_API_KEY or create ~/nina.yml with api_key
	`)

	return ""
}
