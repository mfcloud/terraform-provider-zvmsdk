package zvmsdk

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	zvmsdkgolib "github.com/mfcloud/zvmsdk-go"
)

func resourceZVMImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceZVMImageCreate,
		Delete: resourceZVMImageDelete,
		Read:   resourceZVMImageRead,
		Exists: resourceZVMImageExists,
		Update: resourceZVMImageUpdate,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"meta": {
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"osversion": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"md5sum": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"remotehost": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceZVMImageCreate(d *schema.ResourceData, meta interface{}) error {
	var name string
	if item, ok := d.GetOk("name"); ok {
		name = item.(string)
	}

	var remotehost string
	if item, ok := d.GetOk("remotehost"); ok {
		remotehost = item.(string)
	}

	var locationurl string
	if item, ok := d.GetOk("url"); ok {
		locationurl = item.(string)
	}

	url := meta.(*Client).url

	d.SetId(name)

	m := make(map[string]string)
	if item, ok := d.GetOk("meta.osversion"); ok {
		// FIXME: make it a struct later
		m["os_version"] = item.(string)
	}
	if item, ok := d.GetOk("meta.md5sum"); ok {
		// FIXME: make it a struct later
		m["md5sum"] = item.(string)
	}
	var body zvmsdkgolib.ImageCreateBody
	body.Name = name
	body.RemoteHost = remotehost
	body.Meta = m
	body.URL = locationurl

	zvmsdkgolib.ImageCreate(url, body)

	return nil
}

func resourceZVMImageExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	return true, nil
}

func resourceZVMImageDelete(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

	var imagename string
	if name, ok := d.GetOk("name"); ok {
		imagename = name.(string)
	}

	zvmsdkgolib.ImageDelete(url, imagename)

	return nil
}

func resourceZVMImageUpdate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceZVMImageRead(d *schema.ResourceData, meta interface{}) error {
	url := meta.(*Client).url

	var imagename string
	if name, ok := d.GetOk("name"); ok {
		imagename = name.(string)
	}

	zvmsdkgolib.ImageGet(url, imagename)

	return nil
}
