package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Host   *url.URL
	Client *http.Client
}

func NewClient(host string) (*Client, error) {
	parsedUrl, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %v", err)
	}
	return &Client{
		parsedUrl,
		&http.Client{Timeout: 0},
	}, nil
}

func (c *Client) ConstructPath(path string) string {
	constructedPath := *c.Host
	constructedPath.Path = path
	return constructedPath.String()
}

func (c *Client) GetRequest(ctx context.Context, path string) ([]byte, int, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ConstructPath(path), nil)
	if err != nil {
		return nil, 0, err
	}
	response, err := c.Client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	answer, err := ioutil.ReadAll(response.Body)
	return answer, response.StatusCode, err
}

func (c *Client) PostRequest(ctx context.Context, path string, value any) ([]byte, int, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, 0, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ConstructPath(path), bytes.NewReader(data))
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.Client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	answer, err := ioutil.ReadAll(response.Body)
	return answer, response.StatusCode, err
}

func (c *Client) RequestVersion(ctx context.Context) ([]byte, int, error) {
	return c.GetRequest(ctx, "/version")
}

func (c *Client) RequestHardOp(ctx context.Context) ([]byte, int, error) {
	return c.GetRequest(ctx, "/hard-op")
}

func (c *Client) RequestDecode(ctx context.Context, string64 string) ([]byte, int, error) {
	value := map[string]string{"inputString": string64}
	return c.PostRequest(ctx, "/decode", value)
}
