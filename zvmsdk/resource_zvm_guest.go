package zvmsdk

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mfcloud/terraform-provider-zvmsdk/logger"
	zvmsdkgolib "github.com/mfcloud/zvmsdk-go"
)

func resourceZVMGuest() *schema.Resource {
	return &schema.Resource{
		Create: resourceZVMGuestCreate,
		Delete: resourceZVMGuestDelete,
		Read:   resourceZVMGuestRead,
		Exists: resourceZVMGuestExists,
		Update: resourceZVMGuestUpdate,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: false,
			},
			"diskpool": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"userprofile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"imageid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2048,
				ForceNew: false,
			},
			"disklist": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"diskpool": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"boot": {
							Type:     schema.TypeString,
							Default:  "0",
							Optional: true,
							ForceNew: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceZVMGuestCreate(d *schema.ResourceData, meta interface{}) error {
	var userid string
	if name, ok := d.GetOk("userid"); ok {
		userid = name.(string)
	}
	url := meta.(*Client).url
	logger.Log.Printf("Start to create %s", userid)

	d.SetId(userid)

	var body zvmsdkgolib.GuestCreateBody
	body.Userid = userid
	body.Vcpus = d.Get("vcpus").(int)
	body.DiskPool = d.Get("diskpool").(string)
	body.Memory = d.Get("memory").(int)
	body.UserProfile = d.Get("userprofile").(string)

	for i := 0; i < d.Get("disklist.#").(int); i++ {
		var disk zvmsdkgolib.GuestCreateDisk

		prefix := fmt.Sprintf("disklist.%d", i)
		if size, ok := d.GetOk(prefix + ".size"); ok {
			disk.Size = size.(string)
		}
		if format, ok := d.GetOk(prefix + ".format"); ok {
			disk.Format = format.(string)
		}
		if boot, ok := d.GetOk(prefix + ".boot"); ok {
			disk.Boot = boot.(string)
		}

		body.DiskList = append(body.DiskList, disk)
	}

	logger.Log.Printf("Create guest with: %+v", body)
	res, data := zvmsdkgolib.GuestCreate(url, body)
	logger.Log.Printf("Create guest ret: %+v, %+v", res, string(data))

	var deploybody zvmsdkgolib.GuestDeployBody
	deploybody.Image = d.Get("imageid").(string)
	deploybody.Vdev = "100"
	logger.Log.Printf("Deploy guest with: %+v", body)
	res, data = zvmsdkgolib.GuestDeploy(url, userid, deploybody)
	logger.Log.Printf("Deploy guest ret: %+v, %+v", res, string(data))

	return nil
}

func resourceZVMGuestExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	url := meta.(*Client).url

	var userid string
	if name, ok := d.GetOk("userid"); ok {
		userid = name.(string)
	}

	zvmsdkgolib.GuestGet(url, userid)

	return true, nil
}

func resourceZVMGuestDelete(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

	var userid string
	if name, ok := d.GetOk("userid"); ok {
		userid = name.(string)
	}

	zvmsdkgolib.GuestDelete(url, userid)

	return nil
}

func resourceZVMGuestUpdate(d *schema.ResourceData, meta interface{}) error {
	var userid string
	if name, ok := d.GetOk("userid"); ok {
		userid = name.(string)
	}

	url := meta.(*Client).url

	var body zvmsdkgolib.GuestCreateBody
	body.Userid = userid
	body.Vcpus = 2

	zvmsdkgolib.GuestCreate(url, body)
	return nil
}

func resourceZVMGuestRead(d *schema.ResourceData, meta interface{}) error {
	var userid string
	if name, ok := d.GetOk("userid"); ok {
		userid = name.(string)
	}

	url := meta.(*Client).url

	zvmsdkgolib.GuestGet(url, userid)

	return nil
}
