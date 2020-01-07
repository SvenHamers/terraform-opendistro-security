package main

import (
	"github.com/WhizUs/go-opendistro"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPENDISTRO_USER", nil),
				Description: "Opendistro user",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPENDISTRO_PASSWORD", nil),
				Description: "Opendistro password",
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPENDISTRO_URL", nil),
				Description: "opendistro base url",
			},
			"allow_insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "allow of insecure trafic",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"opendistro_user":         resourceUser(),
			"opendistro_role":         resourceRole(),
			"opendistro_role_mapping": resourcerolemapping(),
			"opendistro_tenant":       resourceTenant(),
			"opendistro_action_group": resourceActionGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"opendistro_health": dataSourceHealthCurrent(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	tlscfg := &opendistro.TLSConfig{
		Insecure: d.Get("allow_insecure").(bool),
	}

	cfg := &opendistro.ClientConfig{
		Username:  d.Get("user").(string),
		Password:  d.Get("password").(string),
		BaseURL:   d.Get("base_url").(string),
		TLSConfig: tlscfg,
	}

	return cfg, nil
}
