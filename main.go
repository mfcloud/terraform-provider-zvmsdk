package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mfcloud/terraform-provider-zvmsdk/zvmsdk"
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
