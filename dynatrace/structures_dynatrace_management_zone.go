package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandManagementZone(d *schema.ResourceData) (*dynatraceConfigV1.ManagementZone, error) {

	var dtManagementZone dynatraceConfigV1.ManagementZone

	if name, ok := d.GetOk("name"); ok {
		dtManagementZone.SetName(name.(string))
	}

	if rules, ok := d.GetOk("rule"); ok {
		dtManagementZone.SetRules(expandManagementZoneRules(rules.([]interface{})))
	}

	if dimensionalRules, ok := d.GetOk("dimensional_rule"); ok {
		dtManagementZone.SetDimensionalRules(expandDimensionalRules(dimensionalRules.([]interface{})))
	}

	return &dtManagementZone, nil

}

func expandDimensionalRules(dimensionalRules []interface{}) []dynatraceConfigV1.DimensionalManagementZoneRuleDto {

	drs := make([]dynatraceConfigV1.DimensionalManagementZoneRuleDto, len(dimensionalRules))

	for i, dimensionalRule := range dimensionalRules {
		m := dimensionalRule.(map[string]interface{})

		var dtDimensionalRule dynatraceConfigV1.DimensionalManagementZoneRuleDto

		if enabled, ok := m["enabled"].(bool); ok {
			dtDimensionalRule.SetEnabled(enabled)
		}

		if appliesTo, ok := m["applies_to"].(string); ok {
			dtDimensionalRule.SetAppliesTo(appliesTo)
		}

		if conditions, ok := m["condition"].([]interface{}); ok {
			dtDimensionalRule.SetConditions(expandDimensionalConditions(conditions))
		}
		drs[i] = dtDimensionalRule

	}

	return drs
}

func expandDimensionalConditions(conditions []interface{}) []dynatraceConfigV1.DimensionalManagementZoneConditionDto {

	dcs := make([]dynatraceConfigV1.DimensionalManagementZoneConditionDto, len(conditions))

	for i, condition := range conditions {

		m := condition.(map[string]interface{})

		var dtDimensionalCondition dynatraceConfigV1.DimensionalManagementZoneConditionDto

		if conditionType, ok := m["condition_type"].(string); ok {
			dtDimensionalCondition.SetConditionType(conditionType)
		}

		if ruleMatcher, ok := m["rule_matcher"].(string); ok {
			dtDimensionalCondition.SetRuleMatcher(ruleMatcher)
		}

		if key, ok := m["key"].(string); ok {
			dtDimensionalCondition.SetKey(key)
		}

		if value, ok := m["value"].(string); ok && len(value) != 0 {
			dtDimensionalCondition.SetValue(value)
		}
		dcs[i] = dtDimensionalCondition
	}

	return dcs

}

func expandManagementZoneRules(rules []interface{}) []dynatraceConfigV1.ManagementZoneRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.ManagementZoneRule{}
	}

	mrs := make([]dynatraceConfigV1.ManagementZoneRule, len(rules))

	for i, rule := range rules {

		m := rule.(map[string]interface{})

		var dtManagementZoneRule dynatraceConfigV1.ManagementZoneRule

		if mzType, ok := m["type"].(string); ok {
			dtManagementZoneRule.SetType(mzType)
		}

		if enabled, ok := m["enabled"].(bool); ok {
			dtManagementZoneRule.SetEnabled(enabled)
		}

		if propagationTypes, ok := m["propagation_types"]; ok {
			dtManagementZoneRule.SetPropagationTypes(expandManagementZonePropagationTypes(propagationTypes.([]interface{})))
		}

		if conditions, ok := m["condition"].([]interface{}); ok {
			dtManagementZoneRule.SetConditions(expandConditions(conditions))
		}

		mrs[i] = dtManagementZoneRule
	}

	return mrs
}

func expandManagementZonePropagationTypes(propagationTypes []interface{}) []string {
	pts := make([]string, len(propagationTypes))

	for i, v := range propagationTypes {
		pts[i] = v.(string)
	}

	return pts

}

func flattenManagementZoneRulesData(managementZoneRules *[]dynatraceConfigV1.ManagementZoneRule) []interface{} {
	if managementZoneRules != nil {
		mrs := make([]interface{}, len(*managementZoneRules))

		for i, managementZoneRules := range *managementZoneRules {
			mr := make(map[string]interface{})

			mr["type"] = managementZoneRules.Type
			mr["enabled"] = managementZoneRules.Enabled
			mr["propagation_types"] = flattenManagementZonePropagationTypes(managementZoneRules.PropagationTypes)
			mr["condition"] = flattenConditionsData(&managementZoneRules.Conditions)
			mrs[i] = mr

		}
		return mrs
	}

	return make([]interface{}, 0)
}

func flattenManagementZonePropagationTypes(propagationTypes *[]string) *[]string {
	if propagationTypes == nil {
		return nil
	}

	pts := make([]string, len(*propagationTypes))

	for i, e := range *propagationTypes {
		pts[i] = e
	}

	return &pts
}
