package main

import (
	"context"
	"fmt"

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

	name := d.Get("name").(string)

	client, _ := opendistro.NewClient(m.(*opendistro.ClientConfig))

	err := client.Security.Tenants.Create(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to create tenant (%s): %s", name, err)
	}

	d.SetId(name)

	return resourceTenantRead(d, m)
}

func resourceTenantRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceTenantDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()
	name := d.Get("name").(string)
	client, _ := opendistro.NewClient(m.(*opendistro.ClientConfig))

	err := client.Security.Tenants.Delete(ctx, name)

	if err != nil {
		return fmt.Errorf("failed to delete tenant (%s): %s", name, err)
	}

	return nil
}
