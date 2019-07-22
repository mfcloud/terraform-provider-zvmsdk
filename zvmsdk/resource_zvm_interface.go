package zvmsdk

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	zvmsdkgolib "github.com/mfcloud/zvmsdk-go"
)

func resourceZVMInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceZVMInterfaceCreate,
		Delete: resourceZVMInterfaceDelete,
		Read:   resourceZVMInterfaceRead,
		//Exists: resourceZVMInterfaceExists,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"osversion": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"active": {
                                Type:     schema.TypeInt,
				Default:  1,
                                Optional: true,
                                ForceNew: true,
                        },
			"networks": {
				Type:     schema.TypeList,
                                Required: true,
                                ForceNew: true,
                                Elem: &schema.Resource{
                                        Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"cidr": {
                                                        Type:     schema.TypeString,
                                                        Optional: true,
                                                        ForceNew: true,
                                                },
						"gateway": {
                                                        Type:     schema.TypeString,
                                                        Optional: true,
                                                        ForceNew: true,
                                                },
						"vdev": {
                                                        Type:     schema.TypeString,
                                                        Required: true,
                                                        ForceNew: true,
                                                },
						"mac": {
                                                        Type:     schema.TypeString,
                                                        Optional: true,
                                                        ForceNew: true,
                                                },
						"nicid": {
                                                        Type:     schema.TypeString,
                                                        Optional: true,
                                                        ForceNew: true,
                                                },
						"osa": {
                                                        Type:     schema.TypeString,
                                                        Optional: true,
                                                        ForceNew: true,
                                                },
						"dns": {
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

func resourceZVMInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	var userid string
	if item, ok := d.GetOk("userid"); ok {
		userid = item.(string)
	}

	var osversion string
        if item, ok := d.GetOk("osversion"); ok {
                osversion = item.(string)
        }

	var active int
        if item, ok := d.GetOk("active"); ok {
                active = item.(int)
        }

	url := meta.(*Client).url
	d.SetId(userid)

	var body zvmsdkgolib.GuestInterfaceCreateBody
	body.Userid = userid
	body.If.Osversion = osversion
	body.If.Active = active

	for i := 0; i < d.Get("networks.#").(int); i++ {

                var interf zvmsdkgolib.GuestNetwork

                prefix := fmt.Sprintf("networks.%d", i)
                if item, ok := d.GetOk(prefix + ".ip"); ok {
                        interf.IP = item.(string)
                }
                if item, ok := d.GetOk(prefix + ".cidr"); ok {
                        interf.Cidr = item.(string)
                }
                if item, ok := d.GetOk(prefix + ".gateway"); ok {
                        interf.Gateway = item.(string)
                }
		if item, ok := d.GetOk(prefix + ".vdev"); ok {
                        interf.Vdev = item.(string)
                }
                if item, ok := d.GetOk(prefix + ".mac"); ok {
                        interf.Mac = item.(string)
                }
                if item, ok := d.GetOk(prefix + ".nicid"); ok {
                        interf.NicID = item.(string)
                }
		if item, ok := d.GetOk(prefix + ".osa"); ok {
                        interf.Osa = item.(string)
                }
                if item, ok := d.GetOk(prefix + ".dns"); ok {
                        interf.Dns = item.(string)
                }

                body.If.Networks = append(body.If.Networks, interf)
        }
	zvmsdkgolib.GuestInterfaceCreate(url, body)
	return nil
}

func resourceZVMInterfaceRead(d *schema.ResourceData, meta interface{}) error {

        return nil
}

func resourceZVMInterfaceDelete(d *schema.ResourceData, meta interface{}) error {

        return nil
}
