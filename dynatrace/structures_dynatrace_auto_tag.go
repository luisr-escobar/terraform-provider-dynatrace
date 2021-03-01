package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandAutoTag(d *schema.ResourceData) (*dynatraceConfigV1.AutoTag, error) {

	var dtAutoTag dynatraceConfigV1.AutoTag

	if name, ok := d.GetOk("name"); ok {
		dtAutoTag.SetName(name.(string))
	}

	if rules, ok := d.GetOk("rule"); ok {
		dtAutoTag.SetRules(expandAutoTagRules(rules.([]interface{})))
	}

	return &dtAutoTag, nil
}

func expandAutoTagRules(rules []interface{}) []dynatraceConfigV1.AutoTagRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.AutoTagRule{}
	}

	ats := make([]dynatraceConfigV1.AutoTagRule, len(rules))

	for i, rule := range rules {

		m := rule.(map[string]interface{})

		var dtAutoTagRule dynatraceConfigV1.AutoTagRule

		if trType, ok := m["type"].(string); ok {
			dtAutoTagRule.SetType(trType)
		}

		if enabled, ok := m["enabled"].(bool); ok {
			dtAutoTagRule.SetEnabled(enabled)
		}

		if valueFormat, ok := m["value_format"].(string); ok {
			dtAutoTagRule.SetValueFormat(valueFormat)
		}

		if propagationTypes, ok := m["propagation_types"]; ok {
			dtAutoTagRule.SetPropagationTypes(expandTagRulePropagationTypes(propagationTypes.([]interface{})))
		}

		if conditions, ok := m["condition"].([]interface{}); ok {
			dtAutoTagRule.SetConditions(expandConditions(conditions))
		}

		ats[i] = dtAutoTagRule
	}

	return ats

}

func expandTagRulePropagationTypes(propagationTypes []interface{}) []string {
	pts := make([]string, len(propagationTypes))

	for i, v := range propagationTypes {
		pts[i] = v.(string)
	}

	return pts

}

func flattenAutoTagRulesData(autoTagRules *[]dynatraceConfigV1.AutoTagRule) []interface{} {
	if autoTagRules != nil {
		ars := make([]interface{}, len(*autoTagRules))

		for i, autoTagRules := range *autoTagRules {
			ar := make(map[string]interface{})

			ar["type"] = autoTagRules.Type
			ar["enabled"] = autoTagRules.Enabled
			ar["value_format"] = autoTagRules.ValueFormat
			ar["propagation_types"] = flattenAutoTagPropagationTypes(autoTagRules.PropagationTypes)
			ar["condition"] = flattenConditionsData(&autoTagRules.Conditions)
			ars[i] = ar

		}
		return ars
	}

	return make([]interface{}, 0)

}

func flattenAutoTagPropagationTypes(propagationTypes *[]string) interface{} {
	if propagationTypes == nil {
		return nil
	}

	pts := make([]string, len(*propagationTypes))

	for i, e := range *propagationTypes {
		pts[i] = e
	}

	return pts
}
