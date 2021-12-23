/*
*TBD*

Example Usage

```hcl
resource "awx_workflow_job_template_schedule" "default" {
  workflow_job_template_id      = awx_workflow_job_template.default.id

  name                      = "schedule-test"
  rrule                     = "DTSTART;TZID=Europe/Paris:20211214T120000 RRULE:INTERVAL=1;FREQ=DAILY"
}
```

*/
package awx

import (
	"context"
	"fmt"
	"log"
	"strconv"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkflowJobTemplateSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowJobTemplateScheduleCreate,
		ReadContext:   resourceScheduleRead,
		UpdateContext: resourceScheduleUpdate,
		DeleteContext: resourceScheduleDelete,
		Schema: map[string]*schema.Schema{

			"workflow_job_template_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The workflow_job_template id for this schedule",
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rrule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: false,
				Default:  true,
			},
		},
	}
}

func resourceWorkflowJobTemplateScheduleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.WorkflowJobTemplateScheduleService

	workflowJobTemplateID := d.Get("workflow_job_template_id").(int)

	result, err := awxService.CreateWorkflowJobTemplateSchedule(workflowJobTemplateID, map[string]interface{}{
		"name":        d.Get("name").(string),
		"rrule":       d.Get("rrule").(int),
		"description": d.Get("description").(string),
		"enabled":     d.Get("enabled").(bool),
	}, map[string]string{})
	if err != nil {
		log.Printf("Fail to Create Schedule for WorkflowJobTemplate %d: %v", workflowJobTemplateID, err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Schedule",
			Detail:   fmt.Sprintf("Schedule failed to create %s", err.Error()),
		})
		return diags
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceScheduleRead(ctx, d, m)
}