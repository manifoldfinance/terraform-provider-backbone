package backbone

import (
	backboneRest "github.com/eddiezane/backbone-rest-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTodoistProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Read: dataSourceTodoistProjectRead,
	}
}

func dataSourceTodoistProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*backboneRest.Client)

	name := d.Get("name").(string)

	project, err := client.GetProjectByName(name)
	if err != nil {
		return err
	}

	d.SetId(project.Id)

	return nil
}
