---
layout: "awx"
page_title: "AWX: awx_notification_template"
sidebar_current: "docs-awx-resource-notification-template"
description: |-
  Manages a notification template in AWX.
---

# awx_notification_template

Manages a notification template in AWX. Notification templates define how and where notifications are sent when certain events occur, such as job completions, failures, or approvals. Supported notification types include email, Slack, webhook, PagerDuty, IRC, and others.

## Example Usage

```hcl
resource "awx_notification_template" "default" {
    name                      = "schedule-test"
    notification_type         = "webhook"
    organization_id           = data.awx_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this notification template.
* `notification_type` - (Required) The type of notification (e.g., "email", "slack", "webhook", "pagerduty", "irc", "mattermost", "rocketchat", "twilio").
* `organization_id` - (Required) The ID of the organization that this notification template belongs to.
* `description` - (Optional) Optional description of this notification template.
* `notification_configuration` - (Optional) The notification configuration in JSON format. The expected fields depend on the notification type.
