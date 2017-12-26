package zvmsdk

import (
        "fmt"
        "time"

        "github.com/hashicorp/terraform/helper/schema"
        zvmsdkgolib "github.com/zvmsdk-go"
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
                },
        }
}

func resourceZVMGuestCreate(d *schema.ResourceData, meta interface{}) error {
        var guestid string
        if name, ok := d.GetOk("userid"); ok {
                guestid = name.(string)
        }

	url := meta.(*Client).url

        d.SetId(guestid)

        var body zvmsdkgolib.GuestCreateBody
        body.Userid = guestid
        body.Vcpus = 2

        zvmsdkgolib.GuestCreate(url, body)

        return nil
}

func resourceZVMGuestExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	url := meta.(*Client).url

        var guestid string
        if name, ok := d.GetOk("userid"); ok {
                guestid = name.(string)
        }

	zvmsdkgolib.GuestQuery(url, guestid)

        return true, nil
}


func resourceZVMGuestDelete(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

        var guestid string
        if name, ok := d.GetOk("userid"); ok {
                guestid = name.(string)
        }

        zvmsdkgolib.GuestDelete(url, guestid)

        return nil
}

func resourceZVMGuestUpdate(d *schema.ResourceData, meta interface{}) error {
        return nil
}


func resourceZVMGuestRead(d *schema.ResourceData, meta interface{}) error {
        if userid, ok := d.GetOk("userid"); ok {
                fmt.Printf("%s", userid)
        }

        return nil
}

