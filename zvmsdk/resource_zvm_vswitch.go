package zvmsdk

import (
        "time"
        "net/http"
	"github.com/terraform-provider-zvmsdk/logger"
        "github.com/hashicorp/terraform/helper/schema"
)


func resourceZVMVSwitch() *schema.Resource {
        return &schema.Resource{
                Create: resourceZVMVSwitchCreate,
                Delete: resourceZVMVSwitchDelete,
                Read:   resourceZVMVSwitchRead,
                Exists: resourceZVMVSwitchExists,
		Update: resourceZVMVSwitchUpdate,
                Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
                Schema: map[string]*schema.Schema{
                        "name": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: false,
                        },
                },
        }
}

func resourceZVMVSwitchCreate(d *schema.ResourceData, meta interface{}) error {
        var vswitch string
	if name, ok := d.GetOk("name"); ok {
		vswitch = name.(string)
	}

        d.SetId(vswitch)
	logger.Log.Printf("create vswitch %s", vswitch) 

	resp, err := http.Get("http://www.baidu.com/")
        if err != nil {
		return nil
	}
	logger.Log.Printf("create vswitch %s", resp)

        return nil
}

func resourceZVMVSwitchExists(d *schema.ResourceData, meta interface{}) (bool, error) {
        var vswitch string
        if name, ok := d.GetOk("name"); ok {
                vswitch = name.(string)
        }
	logger.Log.Printf("check vswitch exist %s", vswitch)       

        return true, nil
}


func resourceZVMVSwitchDelete(d *schema.ResourceData, meta interface{}) error {
        var vswitch string
        if name, ok := d.GetOk("name"); ok {
                vswitch = name.(string)
        }
        logger.Log.Printf("Delete %s", vswitch)

        return nil
}

func resourceZVMVSwitchUpdate(d *schema.ResourceData, meta interface{}) error {
        var vswitch string
        if name, ok := d.GetOk("name"); ok {
                vswitch = name.(string)
        }
	logger.Log.Printf("Update %s", vswitch)
        return nil
}


func resourceZVMVSwitchRead(d *schema.ResourceData, meta interface{}) error {
        var vswitch string
        if name, ok := d.GetOk("name"); ok {
                vswitch = name.(string)
        }
        logger.Log.Printf("Update %s", vswitch)

        return nil
}

