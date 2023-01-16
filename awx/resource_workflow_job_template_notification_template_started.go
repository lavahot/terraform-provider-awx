/*
Add a notification on the workflow job template at startup.

Example Usage

```hcl
resource "awx_workflow_job_template_notification_template_started" "baseconfig" {
    workflow_job_template_id   = data.awx_workflow_job_template.baseconfig.id
    notification_template_id   = data.awx_notification_template.default.id
}
```

*/

package awx

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkflowJobTemplateNotificationTemplateStarted() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowJobTemplateNotificationTemplateCreateForType("started"),
		DeleteContext: resourceWorkflowJobTemplateNotificationTemplateDeleteForType("started"),
		ReadContext:   resourceWorkflowJobTemplateNotificationTemplateRead,

		Schema: map[string]*schema.Schema{
			"workflow_job_template_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"notification_template_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
