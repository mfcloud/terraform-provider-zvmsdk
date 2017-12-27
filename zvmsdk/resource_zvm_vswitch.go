package zvmsdk

import (
        "time"
	"github.com/terraform-provider-zvmsdk/logger"
        "github.com/hashicorp/terraform/helper/schema"
	zvmsdkgolib "github.com/zvmsdk-go"
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
                                ForceNew: true,
                        },
                        "rdev": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: false,
                        },

                },
        }
}

func resourceZVMVSwitchCreate(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

        var vswitchname,vswitchrdev string
	if name, ok := d.GetOk("name"); ok {
		vswitchname = name.(string)
	}

        if rdev, ok := d.GetOk("rdev"); ok {
                vswitchrdev = rdev.(string)
        }


        d.SetId(vswitchname)

	var body zvmsdkgolib.VswitchCreateBody
	body.Name = vswitchname
	body.Rdev = vswitchrdev

	zvmsdkgolib.VswitchCreate(url, body)

        return nil
}

func resourceZVMVSwitchExists(d *schema.ResourceData, meta interface{}) (bool, error) {
        var vswitchname string
        if name, ok := d.GetOk("name"); ok {
                vswitchname = name.(string)
        }
	logger.Log.Printf("Read %s", vswitchname)
        return true, nil
}


func resourceZVMVSwitchDelete(d *schema.ResourceData, meta interface{}) error {

	url := meta.(*Client).url

        var vswitchname string
        if name, ok := d.GetOk("name"); ok {
                vswitchname = name.(string)
        }

	zvmsdkgolib.VswitchDelete(url, vswitchname)

        return nil
}

func resourceZVMVSwitchUpdate(d *schema.ResourceData, meta interface{}) error {
        var vswitchname string
        if name, ok := d.GetOk("name"); ok {
                vswitchname = name.(string)
        }
	logger.Log.Printf("Read %s", vswitchname)

        return nil
}


func resourceZVMVSwitchRead(d *schema.ResourceData, meta interface{}) error {
        var vswitchname string
        if name, ok := d.GetOk("name"); ok {
                vswitchname = name.(string)
        }
        logger.Log.Printf("Read %s", vswitchname)

        return nil
}

