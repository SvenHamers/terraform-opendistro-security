package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SvenHamers/go-opendistro"
	"github.com/SvenHamers/go-opendistro/security"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backend_roles": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"opendistro_roles": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}

}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {

	ctx := context.TODO()
	name := d.Get("user_name").(string)

	client, _ := opendistro.NewClient(m.(*opendistro.ClientConfig))

	user := &security.UserCreate{
		Password:        d.Get("password").(string),
		BackendRoles:    expandStringSet(d.Get("backend_roles").([]interface{})),
		OpenDistroRoles: expandStringSet(d.Get("opendistro_roles").([]interface{})),
	}

	err := client.Security.Users.Create(ctx, name, user)
	if err != nil {
		return fmt.Errorf("failed to create user (%s): %s", name, err)
	}

	d.SetId(name)

	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {

	ctx := context.TODO()

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Users.Delete(ctx, d.Get("user_name").(string))
	if reqerr != nil {
		log.Print(reqerr)
	}

	return nil
}
