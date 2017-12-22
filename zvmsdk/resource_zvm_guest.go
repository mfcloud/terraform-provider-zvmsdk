package zvmsdk

import (
        "fmt"
        "time"
        "os"

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

        d.SetId(guestid)

        var body zvmsdkgolib.GuestCreateBody
        body.Userid = guestid
        body.Vcpus = 2

        zvmsdkgolib.GuestCreate(body)

        return nil
}

func resourceZVMGuestExists(d *schema.ResourceData, meta interface{}) (bool, error) {

        return true, nil
}


func resourceZVMGuestDelete(d *schema.ResourceData, meta interface{}) error {
        var guestid string
        if name, ok := d.GetOk("userid"); ok {
                guestid = name.(string)
        }

        zvmsdkgolib.GuestDelete(guestid)

        return nil
}

func resourceZVMGuestUpdate(d *schema.ResourceData, meta interface{}) error {
        if name, ok := d.GetOk("name"); ok {
                fmt.Printf("%s", name)
        }
        return nil
}


func resourceZVMGuestRead(d *schema.ResourceData, meta interface{}) error {
        if name, ok := d.GetOk("name"); ok {
                fmt.Printf("%s", name)
        }
        f, _ := os.Create("/tmp/data4")
        f.Close()

        return nil
}

