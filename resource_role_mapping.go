package main

import (
	"context"
	"fmt"

	"github.com/SvenHamers/go-opendistro"
	"github.com/SvenHamers/go-opendistro/security"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcerolemapping() *schema.Resource {
	return &schema.Resource{
		Create: resourcerolemappingCreate,
		Read:   resourcerolemappingRead,
		Update: resourcerolemappingUpdate,
		Delete: resourcerolemappingDelete,

		Schema: map[string]*schema.Schema{
			"rolemapping_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rolemapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backend_roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}

}

func resourcerolemappingCreate(d *schema.ResourceData, m interface{}) error {

	ctx := context.TODO()
	name := d.Get("rolemapping_name").(string)
	client, _ := opendistro.NewClient(m.(*opendistro.ClientConfig))

	securityRoleMapping := &security.RoleMappingRelations{
		BackendRoles: expandStringSet(d.Get("backend_roles").([]interface{})),
		Hosts:        expandStringSet(d.Get("hosts").([]interface{})),
		Users:        expandStringSet(d.Get("hosts").([]interface{})),
	}

	err := client.Security.Rolesmapping.Create(ctx, name, securityRoleMapping)
	if err != nil {
		return fmt.Errorf("failed to create role mapping (%s): %s", name, err)
	}
	d.SetId(name)
	return resourcerolemappingRead(d, m)
}

func resourcerolemappingRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourcerolemappingUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcerolemappingRead(d, m)
}

func resourcerolemappingDelete(d *schema.ResourceData, m interface{}) error {

	ctx := context.TODO()
	name := d.Get("rolemapping_name").(string)
	client, _ := opendistro.NewClient(m.(*opendistro.ClientConfig))

	err := client.Security.Rolesmapping.Delete(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to delete role mapping (%s): %s", name, err)
	}

	return nil
}
