package main

import (
	"github.com/terraform-provider-zvmsdk/zvmsdk"
	"github.com/hashicorp/terraform/plugin"
	"math/rand"
	"time"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: zvmsdk.Provider,
	})
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
