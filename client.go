package gotesla

import ()

const (
	BASE_URI    = "https://owner-api.teslamotors.com"
	SSO_URI     = "https://auth.tesla.com"
	AUTH_CN_URI = "https://auth.tesla.cn" // refresh access token for China token
)

type Client struct {
}

func NewClient() (*Client, error) {
	c := &Client{}
	return c, nil
}
