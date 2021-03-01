package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandApplicationDetectionRule(d *schema.ResourceData) (*dynatraceConfigV1.ApplicationDetectionRuleConfig, error) {

	var dtApplicationDetectionRuleConfig dynatraceConfigV1.ApplicationDetectionRuleConfig

	if order, ok := d.GetOk("order"); ok {
		dtApplicationDetectionRuleConfig.SetOrder(order.(string))
	}

	if applicationIdentifier, ok := d.GetOk("application_identifier"); ok {
		dtApplicationDetectionRuleConfig.SetApplicationIdentifier(applicationIdentifier.(string))
	}

	if name, ok := d.GetOk("name"); ok {
		dtApplicationDetectionRuleConfig.SetName(name.(string))
	}

	if filterConfig, ok := d.GetOk("filter_config"); ok {
		dtApplicationDetectionRuleConfig.SetFilterConfig(expandDetectionRuleFilterConfig(filterConfig.([]interface{})))
	}

	return &dtApplicationDetectionRuleConfig, nil

}

func expandDetectionRuleFilterConfig(filterConfig []interface{}) dynatraceConfigV1.ApplicationFilter {
	if len(filterConfig) == 0 || filterConfig[0] == nil {
		return dynatraceConfigV1.ApplicationFilter{}
	}

	dtApplicationFilter := dynatraceConfigV1.NewApplicationFilterWithDefaults()

	m := filterConfig[0].(map[string]interface{})

	if pattern, ok := m["pattern"].(string); ok {
		dtApplicationFilter.SetPattern(pattern)
	}

	if applicationMatchType, ok := m["application_match_type"].(string); ok {
		dtApplicationFilter.SetApplicationMatchType(applicationMatchType)
	}

	if applicationMatchTarget, ok := m["application_match_target"].(string); ok {
		dtApplicationFilter.SetApplicationMatchTarget(applicationMatchTarget)
	}

	return *dtApplicationFilter

}

func flattenApplicationDetectionRule(detectionRule dynatraceConfigV1.ApplicationDetectionRuleConfig, d *schema.ResourceData) diag.Diagnostics {

	d.Set("order", &detectionRule.Order)
	d.Set("application_identifier", &detectionRule.ApplicationIdentifier)
	d.Set("name", &detectionRule.Name)

	filterConfig := flattenDetectionRuleFilterConfig(&detectionRule.FilterConfig)
	if err := d.Set("filter_config", filterConfig); err != nil {
		return diag.FromErr(err)
	}

	return nil

}

func flattenDetectionRuleFilterConfig(filterConfig *dynatraceConfigV1.ApplicationFilter) []interface{} {
	if filterConfig == nil {
		return []interface{}{filterConfig}
	}

	m := make(map[string]interface{})

	m["pattern"] = filterConfig.Pattern
	m["application_match_type"] = filterConfig.ApplicationMatchType
	m["application_match_target"] = filterConfig.ApplicationMatchTarget

	return []interface{}{m}

}
