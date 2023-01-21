package main

type Client struct {
	Name string
}

func NewClient(name string) *Client {
	c := &Client{
		Name: name,
	}
	return c
}
