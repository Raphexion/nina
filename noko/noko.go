package noko

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const (
	BaseURLV2 = "https://api.nokotime.com/v2"
)

type Client struct {
	BaseURL     string
	accessToken string
	HTTPClient  *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:     fetchBaseURL(),
		accessToken: fetchAccessToken(),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-NokoToken", c.accessToken)
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
	req.Header.Set("X-NokoToken", c.accessToken)
	req.Header.Set("User-Agent", "Nina/v0.1")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.New("unable to send")
	}

	return nil
}

func fetchAccessToken() string {
	env_key := viper.GetString("NOKO_ACCESS_TOKEN")
	cnf_key := viper.GetString("access_token")

	if env_key != "" {
		return env_key
	}

	if cnf_key != "" {
		return cnf_key
	}

	log.Fatal(`
	Please set enviromental variable NOKO_ACCESS_TOKEN or create ~/nina.yml with access_token
	`)

	return ""
}

func fetchBaseURL() string {
	env_url := viper.GetString("NOKO_BASE_URL")
	cnf_url := viper.GetString("base_url")

	if env_url != "" {
		return env_url
	}

	if cnf_url != "" {
		return cnf_url
	}

	return BaseURLV2
}
