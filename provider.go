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
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "Opendistro user",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "Opendistro password",
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "opendistro base url",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"opendistro_user":         resourceUser(),
			"opendistro_role":         resourceRole(),
			"opendistro_role_mapping": resourcerolemapping(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	tlscfg := &opendistro.TLSConfig{
		CACert:        "",
		CAPath:        "",
		ClientCert:    "",
		ClientKey:     "",
		TLSServerName: "",
		Insecure:      true,
	}

	cfg := &opendistro.ClientConfig{
		Username:  d.Get("user").(string),
		Password:  d.Get("password").(string),
		BaseURL:   d.Get("base_url").(string),
		TLSConfig: tlscfg,
	}

	return cfg, nil
}
