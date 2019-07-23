package zvmsdk

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	zvmsdkgolib "github.com/mfcloud/zvmsdk-go"
)

func resourceZVMVSwitchGrant() *schema.Resource {
	return &schema.Resource{
		Create: resourceZVMVSwitchGrantCreate,
		Delete: resourceZVMVSwitchGrantDelete,
		Read:   resourceZVMVSwitchGrantRead,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vswitch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceZVMVSwitchGrantCreate(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

	var userid, nic, vswitch string
	if item, ok := d.GetOk("userid"); ok {
		userid = item.(string)
	}

	if item, ok := d.GetOk("nic"); ok {
		nic = item.(string)
	}

	if item, ok := d.GetOk("vswitch"); ok {
		vswitch = item.(string)
	}

	Id := fmt.Sprintf("%s-%s", userid, nic)
	d.SetId(Id)

	var body zvmsdkgolib.VswitchGrantCreateBody
	body.Userid = userid
	body.Vswitch = vswitch
	body.Nic = nic

	zvmsdkgolib.VswitchGrant(url, body)

	return nil
}

func resourceZVMVSwitchGrantExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	return true, nil
}

func resourceZVMVSwitchGrantDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceZVMVSwitchGrantRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
