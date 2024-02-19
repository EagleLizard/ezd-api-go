package services

import (
	"context"
)

type Client struct {
}

func CreateClient(ctx context.Context) *Client {
	client := &Client{}
	return client
}
