package zvmsdk

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Global poolMutexKV
var poolMutexKV = mutexkv.NewMutexKV()

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("http://localhost:8080", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"zvm_guest":   resourceZVMGuest(),
			"zvm_vswitch": resourceZVMVSwitch(),
			"zvm_image":   resourceZVMImage(),
			"zvm_interface": resourceZVMInterface(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URI: d.Get("uri").(string),
	}

	return config.Client()
}
