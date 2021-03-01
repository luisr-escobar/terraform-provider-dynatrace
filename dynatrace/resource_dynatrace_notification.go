package dynatrace

import (
	"context"

	"github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceNotification() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceNotificationCreate,
		ReadContext:   resourceDynatraceNotificationRead,
		UpdateContext: resourceDynatraceNotificationUpdate,
		DeleteContext: resourceDynatraceNotificationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the notification configuration.",
				Required:    true,
			},
			"alerting_profile": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the associated alerting profile.",
				Optional:    true,
				Default:     "c21f969b-5f03-333d-83e0-4f8f136e7682",
			},
			"active": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The configuration is enabled (true) or disabled (false).",
				Required:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Defines the actual set of fields depending on the value.",
				Required:    true,
			},
			"job_template_url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The URL of the target Ansible Tower job template.",
				Optional:    true,
			},
			"job_template_id": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The ID of the target Ansible Tower job template.",
				Optional:    true,
			},
			"custom_message": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The custom message of the notification.",
				Optional:    true,
			},
			"subject": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The subject of the email notification.",
				Optional:    true,
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The template of the email notification.",
				Optional:    true,
			},
			"receivers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The list of the email recipients.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cc_receivers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The list of the email CC-recipients",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"bcc_receivers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The list of the email BCC-recipients",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"project_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The project key of the Jira issue to be created by this notification.",
				Optional:    true,
				Sensitive:   true,
			},
			"issue_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The type of the Jira issue to be created by this notification.",
				Optional:    true,
			},
			"summary": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The summary of the Jira issue to be created by this notification.",
				Optional:    true,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The description of the notification.",
				Optional:    true,
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The API key of the target account.",
				Optional:    true,
				Sensitive:   true,
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The region domain of the OpsGenie.",
				Optional:    true,
			},
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the PagerDuty account.",
				Optional:    true,
			},
			"service_api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The API key to access PagerDuty.",
				Optional:    true,
				Sensitive:   true,
			},
			"service_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the service.",
				Optional:    true,
			},
			"instance_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.",
				Optional:    true,
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The URL of the notification endpoint.",
				Optional:    true,
				Sensitive:   true,
			},
			"accept_any_certificate": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Accept any, including self-signed and invalid, SSL certificate (true) or only trusted (false) certificates.",
				Optional:    true,
			},
			"payload": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The content of the notification message.",
				Optional:    true,
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The username required for authentication.",
				Optional:    true,
				Sensitive:   true,
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The password required for authentication.",
				Optional:    true,
				Sensitive:   true,
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The content of the message.",
				Optional:    true,
			},
			"send_incidents": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Send incidents into ServiceNow ITSM.",
				Optional:    true,
			},
			"send_events": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Send events into ServiceNow ITOM.",
				Optional:    true,
			},
			"channel": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.",
				Optional:    true,
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The content of the message.",
				Optional:    true,
				ForceNew:    true,
			},
			"application_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The application key for the Trello account.",
				Optional:    true,
				Sensitive:   true,
			},
			"authorization_token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The application token for the Trello account.",
				Optional:    true,
				Sensitive:   true,
			},
			"board_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Trello board to which the card should be assigned.",
				Optional:    true,
			},
			"list_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Trello list to which the card should be assigned.",
				Optional:    true,
				ForceNew:    true,
			},
			"resolved_list_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Trello list to which the card of the resolved problem should be assigned.",
				Optional:    true,
			},
			"text": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The text of the generated Trello card.",
				Optional:    true,
			},
			"routing_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The routing key, defining the group to be notified.",
				Optional:    true,
				ForceNew:    true,
			},
			"notify_event_merges_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Call webhook if new events merge into existing problems.",
				Optional:    true,
			},
			"headers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of the additional HTTP headers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The name of the HTTP header.",
							Required:    true,
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The value of the HTTP header. May contain an empty value.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDynatraceNotificationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	dn, err := expandDynatraceNotification(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = dynatraceConfigClientV1.NotificationsApi.CreateNotificationConfig(authConfigV1).NotificationConfig(*dn).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create notification",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	// Get id of notification using name and type since there is no response body when notification is created
	name, nameOk := d.GetOk("name")
	dnType, dnTypeOk := d.GetOk("type")

	notification, _, err := dynatraceConfigClientV1.NotificationsApi.ListNotificationConfigs(authConfigV1).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get dynatrace notifications",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	dns := dynatrace.NotificationConfigStub{}

	if nameOk && dnTypeOk {
		for _, a := range *notification.Values {
			if *a.Name == name.(string) && *a.Type == dnType.(string) {
				dns = a
				break
			}
		}
	}

	d.SetId(dns.Id)

	resourceDynatraceNotificationRead(ctx, d, m)

	return diags

}

func resourceDynatraceNotificationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	notificationID := d.Id()

	notification, _, err := dynatraceConfigClientV1.NotificationsApi.GetNotificationConfig(authConfigV1, notificationID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace notification",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenDynatraceNotification(notification, d)

	return diags
}

func resourceDynatraceNotificationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	notificationID := d.Id()

	dn, err := expandDynatraceNotification(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = dynatraceConfigClientV1.NotificationsApi.UpdateNotificationConfig(authConfigV1, notificationID).NotificationConfig(*dn).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update dynatrace notification",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceNotificationRead(ctx, d, m)

}

func resourceDynatraceNotificationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	notificationID := d.Id()

	_, err := dynatraceConfigClientV1.NotificationsApi.DeleteNotificationConfig(authConfigV1, notificationID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace notification",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
