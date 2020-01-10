package main

import (
	"context"
	"log"

	"github.com/SvenHamers/go-opendistro"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceTenantCreate,
		Read:   resourceTenantRead,
		Delete: resourceTenantDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}

}

func resourceTenantCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Tenants.Create(ctx, d.Get("name").(string))
	if reqerr != nil {
		log.Print(reqerr)
	}

	d.SetId(d.Get("name").(string))

	return resourceTenantRead(d, m)
}

func resourceTenantRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceTenantDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Tenants.Delete(ctx, d.Get("name").(string))
	if reqerr != nil {
		log.Print(reqerr)
	}

	return nil
}
