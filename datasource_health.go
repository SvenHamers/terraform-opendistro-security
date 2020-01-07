package main

import (
	"context"
	"log"

	"github.com/WhizUs/go-opendistro"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHealthCurrent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHealthCurrentRead,

		Schema: map[string]*schema.Schema{
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHealthCurrentRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.TODO()

	client, err := opendistro.NewClient(m.(*opendistro.ClientConfig))

	if err != nil {
		log.Print(err)
	}

	data, reqerr := client.Security.Health.Get(ctx)
	if reqerr != nil {
		log.Print(reqerr)
	}

	d.SetId(data.Mode)
	d.Set("message", data.Message)
	d.Set("mode", data.Mode)
	d.Set("status", data.Status)
	return nil
}
