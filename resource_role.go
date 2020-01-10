package main

import (
	"context"
	"log"

	"github.com/SvenHamers/go-opendistro"
	"github.com/SvenHamers/go-opendistro/security"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,

		Schema: map[string]*schema.Schema{
			"role_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"is_static": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"is_reserved": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"is_hidden": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"cluster_permissions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"index_permissions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dls": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  true,
						},
						"index_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
						"fls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
						"masked_fields": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
						"allowed_actions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
					},
				},
			},
			"tenant_permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenant_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
						"allowed_actions": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							ForceNew: true,
						},
					},
				},
			},
		},
	}

}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()

	rolePermission := &security.RolePermissions{
		IndexPermissions:   expandIndexConfigRequest(d.Get("index_permissions").([]interface{})),
		TenantPermissions:  expandTenantConfigRequest(d.Get("tenant_permissions").([]interface{})),
		ClusterPermissions: expandStringSet(d.Get("cluster_permissions").([]interface{})),
	}

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Roles.Create(ctx, d.Get("role_name").(string), rolePermission)
	if reqerr != nil {
		log.Print(reqerr)
	}

	d.SetId(d.Get("role_name").(string))
	return resourceRoleRead(d, m)
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRoleRead(d, m)
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {

	ctx := context.TODO()

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	reqerr := client.Security.Roles.Delete(ctx, d.Get("role_name").(string))
	if reqerr != nil {
		log.Print(reqerr)
	}

	return nil
}

func expandIndexConfigRequest(l []interface{}) *[]security.IndexPermissions {
	if len(l) == 0 {
		return nil
	}

	m := l[0].(map[string]interface{})

	indexPer := security.IndexPermissions{
		Dls:            m["dls"].(string),
		IndexPatterns:  expandStringSet(m["index_patterns"].([]interface{})),
		Fls:            expandStringSet(m["fls"].([]interface{})),
		MaskedFields:   expandStringSet(m["masked_fields"].([]interface{})),
		AllowedActions: expandStringSet(m["allowed_actions"].([]interface{})),
	}

	indexPerArr := &[]security.IndexPermissions{indexPer}
	return indexPerArr

}

func expandTenantConfigRequest(l []interface{}) *[]security.TenantPermissions {
	if len(l) == 0 {
		return nil
	}

	m := l[0].(map[string]interface{})

	tenantPerr := security.TenantPermissions{
		TenantPatterns: expandStringSet(m["tenant_patterns"].([]interface{})),
		AllowedActions: expandStringSet(m["allowed_actions"].([]interface{})),
	}

	tenantPerArr := &[]security.TenantPermissions{tenantPerr}
	return tenantPerArr
}

func expandStringSet(configured []interface{}) []string {
	return expandStringList(configured)
}

func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
