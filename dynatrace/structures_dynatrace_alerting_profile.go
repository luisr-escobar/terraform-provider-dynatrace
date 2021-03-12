package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandAlertingProfile(d *schema.ResourceData) (*dynatraceConfigV1.AlertingProfile, error) {

	var dtAlertingProfile dynatraceConfigV1.AlertingProfile

	if displayName, ok := d.GetOk("display_name"); ok {
		dtAlertingProfile.SetDisplayName(displayName.(string))
	}

	if mzID, ok := d.GetOk("mz_id"); ok {
		dtAlertingProfile.SetMzId(mzID.(string))
	}

	if rule, ok := d.GetOk("rule"); ok {
		dtAlertingProfile.SetRules(expandAlertingProfileRules(rule.([]interface{})))
	}

	if eventTypeFilter, ok := d.GetOk("event_type_filter"); ok {
		dtAlertingProfile.SetEventTypeFilters(expandEventTypeFilters(eventTypeFilter.([]interface{})))
	}

	return &dtAlertingProfile, nil
}

func expandAlertingProfileRules(rules []interface{}) []dynatraceConfigV1.AlertingProfileSeverityRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.AlertingProfileSeverityRule{}
	}

	ars := make([]dynatraceConfigV1.AlertingProfileSeverityRule, len(rules))

	for i, rule := range rules {
		m := rule.(map[string]interface{})

		var dtAlertingProfileRule dynatraceConfigV1.AlertingProfileSeverityRule

		if severityLevel, ok := m["severity_level"].(string); ok {
			dtAlertingProfileRule.SetSeverityLevel(severityLevel)
		}

		if tagFilter, ok := m["tag_filters"].([]interface{}); ok {
			dtAlertingProfileRule.SetTagFilter(expandTagFilter(tagFilter))
		}

		if delayMinutes, ok := m["delay_in_minutes"].(int); ok {
			dtAlertingProfileRule.SetDelayInMinutes(int32(delayMinutes))
		}
		ars[i] = dtAlertingProfileRule
	}
	return ars

}

func expandTagFilter(tagFilter []interface{}) dynatraceConfigV1.AlertingProfileTagFilter {
	if len(tagFilter) == 0 || tagFilter[0] == nil {
		return dynatraceConfigV1.AlertingProfileTagFilter{}
	}

	dtAlertingProfileTagFilter := dynatraceConfigV1.NewAlertingProfileTagFilterWithDefaults()

	m := tagFilter[0].(map[string]interface{})

	if includeMode, ok := m["include_mode"].(string); ok {
		dtAlertingProfileTagFilter.SetIncludeMode(includeMode)
	}

	if tagFilters, ok := m["tag_filter"].([]interface{}); ok {
		dtAlertingProfileTagFilter.SetTagFilters(expandAlertingProfileTagFilters(tagFilters))
	}

	return *dtAlertingProfileTagFilter

}

func expandAlertingProfileTagFilters(tagFilters []interface{}) []dynatraceConfigV1.TagFilter {
	if len(tagFilters) == 0 || tagFilters[0] == nil {
		return []dynatraceConfigV1.TagFilter{}
	}

	tfs := make([]dynatraceConfigV1.TagFilter, len(tagFilters))

	for i, filter := range tagFilters {

		m := filter.(map[string]interface{})
		var dtTagFilter dynatraceConfigV1.TagFilter

		if context, ok := m["context"].(string); ok {
			dtTagFilter.SetContext(context)
		}

		if key, ok := m["key"].(string); ok {
			dtTagFilter.SetKey(key)
		}

		if value, ok := m["value"].(string); ok && len(value) != 0 {
			dtTagFilter.SetValue(value)
		}

		tfs[i] = dtTagFilter

	}

	return tfs
}

func expandEventTypeFilters(typeFilters []interface{}) []dynatraceConfigV1.AlertingEventTypeFilter {
	if len(typeFilters) == 0 || typeFilters[0] == nil {
		return []dynatraceConfigV1.AlertingEventTypeFilter{}
	}

	etfs := make([]dynatraceConfigV1.AlertingEventTypeFilter, len(typeFilters))

	for i, typeFilter := range typeFilters {

		m := typeFilter.(map[string]interface{})

		if v, ok := m["predefined_event_filter"]; ok && len(m["predefined_event_filter"].([]interface{})) != 0 {
			etfs[i] = dynatraceConfigV1.AlertingEventTypeFilter{
				PredefinedEventFilter: expandPredefinedEventFilter(v.([]interface{})),
			}
		}

		if v, ok := m["custom_event_filter"]; ok && len(m["custom_event_filter"].([]interface{})) != 0 {
			cef := v.([]interface{})[0]
			customEventFilter := cef.(map[string]interface{})

			etfs[i] = dynatraceConfigV1.AlertingEventTypeFilter{
				CustomEventFilter: &dynatraceConfigV1.AlertingCustomEventFilter{
					CustomTitleFilter:       expandCustomTextFilter(customEventFilter["custom_title_filter"].([]interface{})),
					CustomDescriptionFilter: expandCustomTextFilter(customEventFilter["custom_description_filter"].([]interface{})),
				},
			}
		}
	}

	return etfs
}

func expandPredefinedEventFilter(predefinedEventFilter []interface{}) *dynatraceConfigV1.AlertingPredefinedEventFilter {
	if len(predefinedEventFilter) == 0 || predefinedEventFilter[0] == nil {
		return nil
	}

	m := predefinedEventFilter[0].(map[string]interface{})

	pef := &dynatraceConfigV1.AlertingPredefinedEventFilter{}

	if negate, ok := m["negate"]; ok {
		pef.Negate = negate.(bool)
	}

	if eventType, ok := m["event_type"]; ok {
		pef.EventType = eventType.(string)
	}

	return pef

}

func expandCustomTextFilter(customTextFilter []interface{}) *dynatraceConfigV1.AlertingCustomTextFilter {
	if len(customTextFilter) == 0 || customTextFilter[0] == nil {
		return nil
	}

	m := customTextFilter[0].(map[string]interface{})

	ctf := &dynatraceConfigV1.AlertingCustomTextFilter{}

	if enabled, ok := m["enabled"]; ok {
		ctf.Enabled = enabled.(bool)
	}

	if value, ok := m["value"]; ok {
		ctf.Value = value.(string)
	}

	if operator, ok := m["operator"]; ok {
		ctf.Operator = operator.(string)
	}

	if negate, ok := m["negate"]; ok {
		ctf.Negate = negate.(bool)
	}

	if caseInsensitive, ok := m["case_insensitive"]; ok {
		ctf.CaseInsensitive = caseInsensitive.(bool)
	}

	return ctf

}

func flattenAlertingProfileRulesData(alertingProfileRules *[]dynatraceConfigV1.AlertingProfileSeverityRule) []interface{} {
	if alertingProfileRules != nil {
		ars := make([]interface{}, len(*alertingProfileRules), len(*alertingProfileRules))

		for i, alertingProfileRules := range *alertingProfileRules {
			ar := make(map[string]interface{})

			ar["severity_level"] = alertingProfileRules.SeverityLevel
			ar["delay_in_minutes"] = alertingProfileRules.DelayInMinutes
			ar["tag_filters"] = flattenAlertingProfileTagFilter(&alertingProfileRules.TagFilter)
			ars[i] = ar
		}

		return ars
	}

	return make([]interface{}, 0)
}

func flattenAlertingProfileTagFilter(alertingProfileTagFilter *dynatraceConfigV1.AlertingProfileTagFilter) []interface{} {
	if alertingProfileTagFilter == nil {
		return []interface{}{alertingProfileTagFilter}
	}
	t := make(map[string]interface{})

	t["include_mode"] = alertingProfileTagFilter.IncludeMode
	t["tag_filter"] = flattenAlertingProfileTagFilters(alertingProfileTagFilter.TagFilters)

	return []interface{}{t}

}

func flattenAlertingProfileTagFilters(alertingProfileTagFilters *[]dynatraceConfigV1.TagFilter) []interface{} {
	if alertingProfileTagFilters != nil {
		tfs := make([]interface{}, len(*alertingProfileTagFilters), len(*alertingProfileTagFilters))

		for i, alertingProfileTagFilters := range *alertingProfileTagFilters {
			tf := make(map[string]interface{})

			tf["context"] = alertingProfileTagFilters.Context
			tf["key"] = alertingProfileTagFilters.Key
			tf["value"] = alertingProfileTagFilters.Value
			tfs[i] = tf
		}

		return tfs
	}

	return make([]interface{}, 0)
}

func flattenAlertingProfileEventTypeFiltersData(alertingProfileEventTypeFilters *[]dynatraceConfigV1.AlertingEventTypeFilter) []interface{} {
	if alertingProfileEventTypeFilters != nil {

		efs := make([]interface{}, len(*alertingProfileEventTypeFilters), len(*alertingProfileEventTypeFilters))

		for i, alertingProfileEventTypeFilters := range *alertingProfileEventTypeFilters {
			ef := make(map[string]interface{})

			ef["predefined_event_filter"] = flattenPredefinedEventFilter(alertingProfileEventTypeFilters.PredefinedEventFilter)
			ef["custom_event_filter"] = flattenCustomEventFilter(alertingProfileEventTypeFilters.CustomEventFilter)
			efs[i] = ef

		}
		return efs
	}

	return make([]interface{}, 0)
}

func flattenPredefinedEventFilter(alertingProfilePredefinedEventFilters *dynatraceConfigV1.AlertingPredefinedEventFilter) []interface{} {
	if alertingProfilePredefinedEventFilters == nil {
		return nil
	}

	pef := make(map[string]interface{})

	pef["event_type"] = alertingProfilePredefinedEventFilters.EventType
	pef["negate"] = alertingProfilePredefinedEventFilters.Negate

	return []interface{}{pef}

}

func flattenCustomEventFilter(alertingProfileCustomEventFilters *dynatraceConfigV1.AlertingCustomEventFilter) []interface{} {
	if alertingProfileCustomEventFilters == nil {
		return nil
	}

	cef := make(map[string]interface{})

	cef["custom_title_filter"] = flattenCustomTextFilter(alertingProfileCustomEventFilters.CustomTitleFilter)
	cef["custom_description_filter"] = flattenCustomTextFilter(alertingProfileCustomEventFilters.CustomDescriptionFilter)

	return []interface{}{cef}

}

func flattenCustomTextFilter(alertingProfileCustomTextFilters *dynatraceConfigV1.AlertingCustomTextFilter) []interface{} {
	if alertingProfileCustomTextFilters == nil {
		return nil
	}

	ctf := make(map[string]interface{})

	ctf["enabled"] = alertingProfileCustomTextFilters.Enabled
	ctf["value"] = alertingProfileCustomTextFilters.Value
	ctf["operator"] = alertingProfileCustomTextFilters.Operator
	ctf["negate"] = alertingProfileCustomTextFilters.Negate
	ctf["case_insensitive"] = alertingProfileCustomTextFilters.CaseInsensitive

	return []interface{}{ctf}
}
