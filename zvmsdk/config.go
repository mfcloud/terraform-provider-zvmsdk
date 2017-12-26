package zvmsdk

import (
	"github.com/terraform-provider-zvmsdk/logger"
)

// Config struct for the libvirt-provider
type Config struct {
	URI string
}

// Client libvirt
type Client struct {
	url string
}

// Client libvirt, generate libvirt client given URI
func (c *Config) Client() (*Client, error) {

	client := &Client{
		url: c.URI,
	}

	logger.Log.Printf("[INFO] Created zvmsdk client %s", c.URI)

	return client, nil
}

