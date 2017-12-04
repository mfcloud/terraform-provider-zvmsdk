package zvmsdk

import (
        "fmt"
        "time"
        "os"

        "github.com/hashicorp/terraform/helper/schema"
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
                        "name": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: false,
                        },
                },
        }
}

func resourceZVMGuestCreate(d *schema.ResourceData, meta interface{}) error {
        var file string
	if name, ok := d.GetOk("name"); ok {
		file  = "/tmp/" + name.(string)
	}
        f, _ := os.Create("/tmp/abc")
        f.Close()
        f, _ = os.Create(file)
        f.Close()

        var id = "00000000-0000-0000-0000000000000001"
        d.SetId(id)

        // the domain ID must always be saved, otherwise it won't be possible to cleanup a domain
        // if something bad happens at provisioning time
        d.Partial(true)
        d.Set("id", id)
        d.SetPartial("id")
        d.Partial(false)


        return nil
}

func resourceZVMGuestExists(d *schema.ResourceData, meta interface{}) (bool, error) {

        return true, nil
}


func resourceZVMGuestDelete(d *schema.ResourceData, meta interface{}) error {
        if name, ok := d.GetOk("name"); ok {
                fmt.Printf("%s", name)
        }
        f, _ := os.Create("/tmp/data3")
        f.Close()

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

