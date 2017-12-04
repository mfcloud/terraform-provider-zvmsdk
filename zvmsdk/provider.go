package zvmsdk

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Global poolMutexKV
var poolMutexKV = mutexkv.NewMutexKV()

// Provider libvirt
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LIBVIRT_DEFAULT_URI", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"zvm_guest":    resourceZVMGuest(),
                        "zvm_vswitch":  resourceZVMVSwitch(),
		},

	}
}
