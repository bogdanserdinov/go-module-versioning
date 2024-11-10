package sdk_versioning

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

type GetPriceResponse struct {
	Bitcoin struct {
		Usd int `json:"usd"`
	} `json:"bitcoin"`
}

const vsCurrency = "usd"

func (c *Client) GetPrice(ctx context.Context, id string) (*GetPriceResponse, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", id, vsCurrency)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := new(GetPriceResponse)

	err = json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
