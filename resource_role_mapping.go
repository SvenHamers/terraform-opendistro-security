package main

import (
	"context"
	"log"

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

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))
	if err != nil {
		log.Print(err)
	}

	securityRoleMapping := &security.RoleMappingRelations{
		BackendRoles: expandStringSet(d.Get("backend_roles").([]interface{})),
		Hosts:        expandStringSet(d.Get("hosts").([]interface{})),
		Users:        expandStringSet(d.Get("hosts").([]interface{})),
	}

	reqerr := client.Security.Rolesmapping.Create(ctx, d.Get("rolemapping_name").(string), securityRoleMapping)
	if reqerr != nil {
		log.Print(reqerr)
	}

	d.SetId(d.Get("rolemapping_name").(string))
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

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))
	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Rolesmapping.Delete(ctx, d.Get("rolemapping_name").(string))
	if reqerr != nil {
		log.Print(reqerr)
	}

	return nil
}
