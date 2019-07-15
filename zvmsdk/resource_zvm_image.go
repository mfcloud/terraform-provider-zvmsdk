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
		},
	}
}

func resourceZVMImageCreate(d *schema.ResourceData, meta interface{}) error {
	var imagename string
	if name, ok := d.GetOk("name"); ok {
		imagename = name.(string)
	}

	url := meta.(*Client).url

	d.SetId(imagename)

	var body zvmsdkgolib.ImageCreateBody
	body.Name = imagename
	body.RemoteHost = "abc"
	body.Meta = nil
	body.URL = "url"

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
