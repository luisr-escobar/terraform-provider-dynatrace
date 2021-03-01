package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandDynatraceNotification(d *schema.ResourceData) (*dynatraceConfigV1.NotificationConfig, error) {

	var dtNotificationConfig dynatraceConfigV1.NotificationConfig

	if name, ok := d.GetOk("name"); ok {
		dtNotificationConfig.SetName(name.(string))
	}

	if alertingProfile, ok := d.GetOk("alerting_profile"); ok {
		dtNotificationConfig.SetAlertingProfile(alertingProfile.(string))
	}

	dtNotificationConfig.SetActive(d.Get("active").(bool))

	if dnType, ok := d.GetOk("type"); ok {
		dtNotificationConfig.SetType(dnType.(string))
	}

	if jobTemplateURL, ok := d.GetOk("job_template_url"); ok {
		dtNotificationConfig.SetJobTemplateURL(jobTemplateURL.(string))
	}

	if jobTemplateID, ok := d.GetOk("job_template_id"); ok {
		dtNotificationConfig.SetJobTemplateID(int32(jobTemplateID.(int)))
	}

	if customMessage, ok := d.GetOk("custom_message"); ok {
		dtNotificationConfig.SetCustomMessage(customMessage.(string))
	}

	if subject, ok := d.GetOk("subject"); ok {
		dtNotificationConfig.SetSubject(subject.(string))
	}

	if body, ok := d.GetOk("body"); ok {
		dtNotificationConfig.SetBody(body.(string))
	}

	if receivers, ok := d.GetOk("receivers"); ok {
		dtNotificationConfig.SetReceivers(expandNotificationReceivers(receivers.([]interface{})))
	}

	if ccReceivers, ok := d.GetOk("cc_receivers"); ok {
		dtNotificationConfig.SetCcReceivers(expandNotificationReceivers(ccReceivers.([]interface{})))
	}

	if bccReceivers, ok := d.GetOk("bcc_receivers"); ok {
		dtNotificationConfig.SetBccReceivers(expandNotificationReceivers(bccReceivers.([]interface{})))
	}

	if projectKey, ok := d.GetOk("project_key"); ok {
		dtNotificationConfig.SetProjectKey(projectKey.(string))
	}

	if issueType, ok := d.GetOk("issue_type"); ok {
		dtNotificationConfig.SetIssueType(issueType.(string))
	}

	if summary, ok := d.GetOk("summary"); ok {
		dtNotificationConfig.SetSummary(summary.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		dtNotificationConfig.SetDescription(description.(string))
	}

	if apiKey, ok := d.GetOk("api_key"); ok {
		dtNotificationConfig.SetApiKey(apiKey.(string))
	}

	if domain, ok := d.GetOk("domain"); ok {
		dtNotificationConfig.SetDomain(domain.(string))
	}

	if account, ok := d.GetOk("account"); ok {
		dtNotificationConfig.SetAccount(account.(string))
	}

	if serviceAPIKey, ok := d.GetOk("service_api_key"); ok {
		dtNotificationConfig.SetServiceApiKey(serviceAPIKey.(string))
	}

	if serviceName, ok := d.GetOk("service_name"); ok {
		dtNotificationConfig.SetServiceName(serviceName.(string))
	}

	if instanceName, ok := d.GetOk("instance_name"); ok {
		dtNotificationConfig.SetInstanceName(instanceName.(string))
	}

	if url, ok := d.GetOk("url"); ok {
		dtNotificationConfig.SetUrl(url.(string))
	}

	dtNotificationConfig.SetAcceptAnyCertificate(d.Get("accept_any_certificate").(bool))

	if payload, ok := d.GetOk("payload"); ok {
		dtNotificationConfig.SetPayload(payload.(string))
	}

	if username, ok := d.GetOk("username"); ok {
		dtNotificationConfig.SetUsername(username.(string))
	}

	if password, ok := d.GetOk("password"); ok {
		dtNotificationConfig.SetPassword(password.(string))
	}

	if message, ok := d.GetOk("message"); ok {
		dtNotificationConfig.SetMessage(message.(string))
	}

	dtNotificationConfig.SetSendIncidents(d.Get("send_incidents").(bool))

	dtNotificationConfig.SetSendEvents(d.Get("send_events").(bool))

	if channel, ok := d.GetOk("channel"); ok {
		dtNotificationConfig.SetChannel(channel.(string))
	}

	if title, ok := d.GetOk("title"); ok {
		dtNotificationConfig.SetTitle(title.(string))
	}

	if applicationKey, ok := d.GetOk("application_key"); ok {
		dtNotificationConfig.SetApplicationKey(applicationKey.(string))
	}

	if authorizationToken, ok := d.GetOk("authorization_token"); ok {
		dtNotificationConfig.SetAuthorizationToken(authorizationToken.(string))
	}

	if boardID, ok := d.GetOk("board_id"); ok {
		dtNotificationConfig.SetBoardId(boardID.(string))
	}

	if listID, ok := d.GetOk("list_id"); ok {
		dtNotificationConfig.SetListId(listID.(string))
	}

	if resolvedListID, ok := d.GetOk("resolved_list_id"); ok {
		dtNotificationConfig.SetResolvedListId(resolvedListID.(string))
	}

	if text, ok := d.GetOk("text"); ok {
		dtNotificationConfig.SetText(text.(string))
	}

	if routingKey, ok := d.GetOk("routing_key"); ok {
		dtNotificationConfig.SetRoutingKey(routingKey.(string))
	}

	dtNotificationConfig.SetNotifyEventMergesEnabled(d.Get("notify_event_merges_enabled").(bool))

	if headers, ok := d.GetOk("headers"); ok {
		dtNotificationConfig.SetHeaders(expandNotificationHeaders(headers.([]interface{})))
	}

	return &dtNotificationConfig, nil

}

func expandNotificationReceivers(receivers []interface{}) []string {
	nrs := make([]string, len(receivers))

	for i, v := range receivers {
		nrs[i] = v.(string)
	}

	return nrs

}

func expandNotificationHeaders(headers []interface{}) []dynatraceConfigV1.HttpHeader {

	nhs := make([]dynatraceConfigV1.HttpHeader, len(headers))

	for i, header := range headers {
		m := header.(map[string]interface{})

		var dtHeader dynatraceConfigV1.HttpHeader

		if name, ok := m["name"].(string); ok {
			dtHeader.SetName(name)
		}

		if value, ok := m["value"].(string); ok && len(value) != 0 {
			dtHeader.SetValue(value)
		}

		nhs[i] = dtHeader

	}

	return nhs

}

func flattenDynatraceNotification(notification dynatraceConfigV1.NotificationConfig, d *schema.ResourceData) diag.Diagnostics {

	notificationReceivers := flattenNotificationReceivers(&notification.Receivers)
	if err := d.Set("receivers", notificationReceivers); err != nil {
		return diag.FromErr(err)
	}

	notificationCCReceivers := flattenNotificationReceivers(&notification.CcReceivers)
	if err := d.Set("cc_receivers", notificationCCReceivers); err != nil {
		return diag.FromErr(err)
	}

	notificationBCCReceivers := flattenNotificationReceivers(&notification.BccReceivers)
	if err := d.Set("bcc_receivers", notificationBCCReceivers); err != nil {
		return diag.FromErr(err)

	}

	notificationHeaders := flattenNotificationHeaders(notification.Headers)
	if err := d.Set("headers", notificationHeaders); err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", &notification.Name)
	d.Set("alerting_profile", &notification.AlertingProfile)
	d.Set("active", &notification.Active)
	d.Set("type", &notification.Type)
	d.Set("job_template_url", &notification.JobTemplateURL)
	d.Set("job_template_id", &notification.JobTemplateID)
	d.Set("custom_message", &notification.CustomMessage)
	d.Set("subject", &notification.Subject)
	d.Set("body", &notification.Body)
	d.Set("project_key", &notification.ProjectKey)
	d.Set("issue_type", &notification.IssueType)
	d.Set("summary", &notification.Summary)
	d.Set("description", &notification.Description)
	d.Set("api_key", &notification.ApiKey)
	d.Set("domain", &notification.Domain)
	d.Set("account", &notification.Account)
	d.Set("service_api_key", &notification.ServiceApiKey)
	d.Set("service_name", &notification.ServiceName)
	d.Set("instance_name", &notification.InstanceName)
	d.Set("url", &notification.Url)
	d.Set("accept_any_certificate", &notification.AcceptAnyCertificate)
	d.Set("username", &notification.Username)
	d.Set("password", &notification.Password)
	d.Set("message", &notification.Message)
	d.Set("send_incidents", &notification.SendIncidents)
	d.Set("send_events", &notification.SendEvents)
	d.Set("channel", &notification.Channel)
	d.Set("title", &notification.Title)
	d.Set("application_key", &notification.ApplicationKey)
	d.Set("authorization_token", &notification.AuthorizationToken)
	d.Set("board_id", &notification.BoardId)
	d.Set("list_id", &notification.ListId)
	d.Set("resolved_list_id", &notification.ResolvedListId)
	d.Set("text", &notification.Text)
	d.Set("routing_key", &notification.RoutingKey)
	d.Set("notify_event_merges_enabled", &notification.NotifyEventMergesEnabled)

	return nil

}

func flattenNotificationReceivers(receivers *[]string) *[]string {
	if receivers == nil {
		return nil
	}

	pts := make([]string, len(*receivers))

	for i, e := range *receivers {
		pts[i] = e
	}

	return &pts
}

func flattenNotificationHeaders(headers *[]dynatraceConfigV1.HttpHeader) []interface{} {
	if headers != nil {
		nhs := make([]interface{}, len(*headers))

		for i, headers := range *headers {
			nh := make(map[string]interface{})

			nh["name"] = headers.Name
			nh["value"] = headers.Value
			nhs[i] = nh

		}
		return nhs
	}

	return make([]interface{}, 0)
}
