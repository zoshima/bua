package main

type Client struct {
	host  string
	token string
}

func NewClient(token string) *Client {
	return &Client{
		"https://api.bua.io/api/bua/a/v1",
		token,
	}
}
