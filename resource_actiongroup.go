package main

import (
	"context"
	"log"

	"github.com/WhizUs/go-opendistro"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceActionGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceActionGroupCreate,
		Read:   resourceActionGroupRead,
		Update: resourceActionGroupUpdate,
		Delete: resourceActionGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed_actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}

}

func resourceActionGroupCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()
	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	rolesArr := []string{}
	for _, element := range d.Get("backend_roles").([]interface{}) {
		rolesArr = append(rolesArr, element.(string))
	}

	reqerr := client.Security.Actiongroups.Create()
	if reqerr != nil {
		log.Print(reqerr)
	}
	d.SetId(d.Get("name").(string))

	return resourceActionGroupRead(d, m)
}

func resourceActionGroupRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceActionGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceActionGroupRead(d, m)
}

func resourceActionGroupDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
