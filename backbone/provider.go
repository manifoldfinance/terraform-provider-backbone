package backbone

import (
	backboneRest "github.com/eddiezane/backbone-rest-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Your Todoist API key",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODOIST_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"backbone_task": resourceTodoistTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"backbone_project": dataSourceTodoistProject(),
		},
		ConfigureFunc: configureFunc(),
	}
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		client := backboneRest.NewClient(d.Get("api_key").(string))
		return client, nil
	}
}
