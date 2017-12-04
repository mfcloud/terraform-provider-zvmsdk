package main

import (
	"github.com/terraform-provider-zvmsdk/zvmsdk"
	"github.com/hashicorp/terraform/plugin"
	"math/rand"
	"time"
        "os"
        "log"
)

func main() {
        log.SetOutput(os.Stderr)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: zvmsdk.Provider,
	})
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
