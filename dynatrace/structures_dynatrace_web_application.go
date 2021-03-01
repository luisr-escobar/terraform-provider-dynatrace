package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandWebApplication(d *schema.ResourceData) (*dynatraceConfigV1.WebApplicationConfig, error) {

	var dtWebApplicationConfig dynatraceConfigV1.WebApplicationConfig

	if name, ok := d.GetOk("name"); ok {
		dtWebApplicationConfig.SetName(name.(string))
	}

	if waType, ok := d.GetOk("type"); ok {
		dtWebApplicationConfig.SetType(waType.(string))
	}

	if realUserMonitoringEnabled, ok := d.GetOk("real_user_monitoring_enabled"); ok {
		dtWebApplicationConfig.SetRealUserMonitoringEnabled(realUserMonitoringEnabled.(bool))
	}

	if costControlUSerSessionPercentage, ok := d.GetOk("cost_control_user_session_percentage"); ok {
		dtWebApplicationConfig.SetCostControlUserSessionPercentage(float32(costControlUSerSessionPercentage.(int)))
	}

	if loadActionKeyPerformanceMetric, ok := d.GetOk("load_action_key_performance_metric"); ok {
		dtWebApplicationConfig.SetLoadActionKeyPerformanceMetric(loadActionKeyPerformanceMetric.(string))
	}

	if xhrActionKeyPerformanceMetric, ok := d.GetOk("xhr_action_key_performance_metric"); ok {
		dtWebApplicationConfig.SetXhrActionKeyPerformanceMetric(xhrActionKeyPerformanceMetric.(string))
	}

	if urlInjectionPattern, ok := d.GetOk("url_injection_pattern"); ok {
		dtWebApplicationConfig.SetUrlInjectionPattern(urlInjectionPattern.(string))
	}

	if sessionReplayConfig, ok := d.GetOk("session_replay_config"); ok {
		dtWebApplicationConfig.SetSessionReplayConfig(expandSessionReplayConfig(sessionReplayConfig.([]interface{})))
	}

	if loadActionApdexSettings, ok := d.GetOk("load_action_apdex_settings"); ok {
		dtWebApplicationConfig.SetLoadActionApdexSettings(expandApdexSettings(loadActionApdexSettings.([]interface{})))
	}

	if xhrActionApdexSettings, ok := d.GetOk("xhr_action_apdex_settings"); ok {
		dtWebApplicationConfig.SetXhrActionApdexSettings(expandApdexSettings(xhrActionApdexSettings.([]interface{})))
	}

	if customActionApdexSettings, ok := d.GetOk("custom_action_apdex_settings"); ok {
		dtWebApplicationConfig.SetCustomActionApdexSettings(expandApdexSettings(customActionApdexSettings.([]interface{})))
	}

	if waterfallSettings, ok := d.GetOk("waterfall_settings"); ok {
		dtWebApplicationConfig.SetWaterfallSettings(expandWaterfallSettings(waterfallSettings.([]interface{})))
	}

	if monitoringSettings, ok := d.GetOk("monitoring_settings"); ok {
		dtWebApplicationConfig.SetMonitoringSettings(expandMonitoringSettings(monitoringSettings.([]interface{})))
	}

	if userTags, ok := d.GetOk("user_tag"); ok {
		dtWebApplicationConfig.SetUserTags(expandUserTags(userTags.([]interface{})))
	}

	if userActionAndSessionProperties, ok := d.GetOk("user_action_and_session_property"); ok {
		dtWebApplicationConfig.SetUserActionAndSessionProperties(expandUserSessionProperties(userActionAndSessionProperties.([]interface{})))
	}

	if userActionNamingSettings, ok := d.GetOk("user_action_naming_settings"); ok {
		dtWebApplicationConfig.SetUserActionNamingSettings(expanduserActionNamingSettings(userActionNamingSettings.([]interface{})))
	}

	if metaDataCaptureSettings, ok := d.GetOk("metadata_capture_setting"); ok {
		dtWebApplicationConfig.SetMetaDataCaptureSettings(expandmetaDataCaptureSettings(metaDataCaptureSettings.([]interface{})))
	}

	if conversionGoals, ok := d.GetOk("conversion_goal"); ok {
		dtWebApplicationConfig.SetConversionGoals(expandconversionGoals(conversionGoals.([]interface{})))
	}

	return &dtWebApplicationConfig, nil

}

func expandSessionReplayConfig(replayConfig []interface{}) dynatraceConfigV1.SessionReplaySetting {
	if len(replayConfig) == 0 || replayConfig[0] == nil {
		return dynatraceConfigV1.SessionReplaySetting{}
	}

	dtSessionReplaySetting := dynatraceConfigV1.NewSessionReplaySettingWithDefaults()

	m := replayConfig[0].(map[string]interface{})

	if enabled, ok := m["enabled"].(bool); ok {
		dtSessionReplaySetting.SetEnabled(enabled)
	}

	if costControlPercentage, ok := m["cost_control_percentage"].(int); ok {
		dtSessionReplaySetting.SetCostControlPercentage(int32(costControlPercentage))
	}

	return *dtSessionReplaySetting

}

func expandconversionGoals(goals []interface{}) []dynatraceConfigV1.ConversionGoal {
	if len(goals) < 1 {
		return []dynatraceConfigV1.ConversionGoal{}
	}

	cgs := make([]dynatraceConfigV1.ConversionGoal, len(goals))

	for i, goal := range goals {

		m := goal.(map[string]interface{})

		var dtConversionGoal dynatraceConfigV1.ConversionGoal

		if name, ok := m["name"].(string); ok {
			dtConversionGoal.SetName(name)
		}

		if id, ok := m["id"].(string); ok && len(id) != 0 {
			dtConversionGoal.SetId(id)
		}

		if cgType, ok := m["type"].(string); ok && len(cgType) != 0 {
			dtConversionGoal.SetType(cgType)
		}

		if destinationDetails, ok := m["destination_details"].([]interface{}); ok && len(destinationDetails) != 0 {
			dtConversionGoal.SetDestinationDetails(expandDestinationDetails(destinationDetails))
		}

		if userActionDetails, ok := m["user_action_details"].([]interface{}); ok && len(userActionDetails) != 0 {
			dtConversionGoal.SetUserActionDetails(expandUserActionDetails(userActionDetails))
		}

		if visitDurationDetails, ok := m["visit_duration_details"].([]interface{}); ok && len(visitDurationDetails) != 0 {
			dtConversionGoal.SetVisitDurationDetails(expandVisitDurationDetails(visitDurationDetails))
		}

		if visitNumActionDetails, ok := m["visit_num_action_details"].([]interface{}); ok && len(visitNumActionDetails) != 0 {
			dtConversionGoal.SetVisitNumActionDetails(expandVisitNumActionDetails(visitNumActionDetails))
		}
		cgs[i] = dtConversionGoal
	}

	return cgs

}

func expandDestinationDetails(destinationDetails []interface{}) dynatraceConfigV1.DestinationDetails {
	if len(destinationDetails) == 0 || destinationDetails[0] == nil {
		return dynatraceConfigV1.DestinationDetails{}
	}

	dtDestinationDetails := dynatraceConfigV1.NewDestinationDetailsWithDefaults()

	m := destinationDetails[0].(map[string]interface{})

	if urlOrPath, ok := m["url_or_path"].(string); ok {
		dtDestinationDetails.SetUrlOrPath(urlOrPath)
	}

	if matchType, ok := m["match_type"].(string); ok && len(matchType) != 0 {
		dtDestinationDetails.SetMatchType(matchType)
	}

	if caseSensitive, ok := m["case_sensitive"].(bool); ok {
		dtDestinationDetails.SetCaseSensitive(caseSensitive)
	}

	return *dtDestinationDetails

}

func expandUserActionDetails(actionDetails []interface{}) dynatraceConfigV1.UserActionDetails {
	if len(actionDetails) == 0 || actionDetails[0] == nil {
		return dynatraceConfigV1.UserActionDetails{}
	}

	dtUserActionDetails := dynatraceConfigV1.NewUserActionDetailsWithDefaults()

	m := actionDetails[0].(map[string]interface{})

	if value, ok := m["url_or_path"].(string); ok && len(value) != 0 {
		dtUserActionDetails.SetValue(value)
	}

	if caseSensitive, ok := m["case_sensitive"].(bool); ok {
		dtUserActionDetails.SetCaseSensitive(caseSensitive)
	}

	if matchType, ok := m["match_type"].(string); ok && len(matchType) != 0 {
		dtUserActionDetails.SetMatchType(matchType)
	}

	if matchEntity, ok := m["match_entity"].(string); ok && len(matchEntity) != 0 {
		dtUserActionDetails.SetMatchEntity(matchEntity)
	}

	if actionType, ok := m["action_type"].(string); ok && len(actionType) != 0 {
		dtUserActionDetails.SetActionType(actionType)
	}

	return *dtUserActionDetails

}

func expandVisitDurationDetails(durationDetails []interface{}) dynatraceConfigV1.VisitDurationDetails {
	if len(durationDetails) == 0 || durationDetails[0] == nil {
		return dynatraceConfigV1.VisitDurationDetails{}
	}

	dtVisitDurationDetails := dynatraceConfigV1.NewVisitDurationDetailsWithDefaults()

	m := durationDetails[0].(map[string]interface{})

	if durationInMillis, ok := m["duration_in_millis"].(int); ok {
		dtVisitDurationDetails.SetDurationInMillis(int64(durationInMillis))
	}

	return *dtVisitDurationDetails

}

func expandVisitNumActionDetails(actionDetails []interface{}) dynatraceConfigV1.VisitNumActionDetails {
	if len(actionDetails) == 0 || actionDetails[0] == nil {
		return dynatraceConfigV1.VisitNumActionDetails{}
	}

	dtVisitNumActionDetails := dynatraceConfigV1.NewVisitNumActionDetailsWithDefaults()

	m := actionDetails[0].(map[string]interface{})

	if numUserActions, ok := m["num_user_actions"].(int); ok {
		dtVisitNumActionDetails.SetNumUserActions(int32(numUserActions))
	}

	return *dtVisitNumActionDetails

}

func expandmetaDataCaptureSettings(settings []interface{}) []dynatraceConfigV1.MetaDataCapturing {
	if len(settings) < 1 {
		return []dynatraceConfigV1.MetaDataCapturing{}
	}

	mcs := make([]dynatraceConfigV1.MetaDataCapturing, len(settings))

	for i, setting := range settings {

		m := setting.(map[string]interface{})

		var dtMetaDataCapturing dynatraceConfigV1.MetaDataCapturing

		if mdType, ok := m["type"].(string); ok {
			dtMetaDataCapturing.SetType(mdType)
		}

		if capturingName, ok := m["capturing_name"].(string); ok {
			dtMetaDataCapturing.SetCapturingName(capturingName)
		}

		if name, ok := m["name"].(string); ok {
			dtMetaDataCapturing.SetName(name)
		}

		if uniqueID, ok := m["unique_id"].(int); ok && uniqueID != 0 {
			dtMetaDataCapturing.SetUniqueId(int32(uniqueID))
		}

		if publicMetadata, ok := m["public_metadata"].(bool); ok {
			dtMetaDataCapturing.SetPublicMetadata(publicMetadata)
		}
		mcs[i] = dtMetaDataCapturing
	}

	return mcs
}

func expandUserTags(tags []interface{}) []dynatraceConfigV1.UserTag {
	if len(tags) < 1 {
		return []dynatraceConfigV1.UserTag{}
	}

	uts := make([]dynatraceConfigV1.UserTag, len(tags))

	for i, tag := range tags {

		m := tag.(map[string]interface{})

		var dtUserTags dynatraceConfigV1.UserTag

		if uniqueID, ok := m["unique_id"].(int); ok {
			dtUserTags.SetUniqueId(int32(uniqueID))
		}

		if metadataID, ok := m["metadata_id"].(int); ok && metadataID != 0 {
			dtUserTags.SetMetadataId(int32(metadataID))
		}

		if cleanupRule, ok := m["cleanup_rule"].(string); ok && len(cleanupRule) != 0 {
			dtUserTags.SetCleanupRule(cleanupRule)
		}

		if serverSideRequestAttribute, ok := m["server_side_request_attribute"].(string); ok && len(serverSideRequestAttribute) != 0 {
			dtUserTags.SetServerSideRequestAttribute(serverSideRequestAttribute)
		}

		if ignoreCase, ok := m["ignore_case"].(bool); ok {
			dtUserTags.SetIgnoreCase(ignoreCase)
		}

		uts[i] = dtUserTags
	}

	return uts
}

func expandUserSessionProperties(properties []interface{}) []dynatraceConfigV1.UserActionAndSessionProperties {
	if len(properties) < 1 {
		return []dynatraceConfigV1.UserActionAndSessionProperties{}
	}

	sps := make([]dynatraceConfigV1.UserActionAndSessionProperties, len(properties))

	for i, property := range properties {

		m := property.(map[string]interface{})

		var dtUserActionAndSessionProperties dynatraceConfigV1.UserActionAndSessionProperties

		if displayName, ok := m["display_name"].(string); ok && len(displayName) != 0 {
			dtUserActionAndSessionProperties.SetDisplayName(displayName)
		}

		if spType, ok := m["type"].(string); ok {
			dtUserActionAndSessionProperties.SetType(spType)
		}

		if origin, ok := m["origin"].(string); ok {
			dtUserActionAndSessionProperties.SetOrigin(origin)
		}

		if aggregation, ok := m["aggregation"].(string); ok && len(aggregation) != 0 {
			dtUserActionAndSessionProperties.SetAggregation(aggregation)
		}

		if storeAsUserActionProperty, ok := m["store_as_user_action_property"].(bool); ok {
			dtUserActionAndSessionProperties.SetStoreAsUserActionProperty(storeAsUserActionProperty)
		}

		if storeAsSessionProperty, ok := m["store_as_session_property"].(bool); ok {
			dtUserActionAndSessionProperties.SetStoreAsSessionProperty(storeAsSessionProperty)
		}

		if cleanupRule, ok := m["cleanup_rule"].(string); ok && len(cleanupRule) != 0 {
			dtUserActionAndSessionProperties.SetCleanupRule(cleanupRule)
		}

		if serverSideRequestAttribute, ok := m["server_side_request_attribute"].(string); ok && len(serverSideRequestAttribute) != 0 {
			dtUserActionAndSessionProperties.SetServerSideRequestAttribute(serverSideRequestAttribute)
		}

		if uniqueID, ok := m["unique_id"].(int); ok {
			dtUserActionAndSessionProperties.SetUniqueId(int32(uniqueID))
		}

		if key, ok := m["key"].(string); ok && len(key) != 0 {
			dtUserActionAndSessionProperties.SetKey(key)
		}

		if metadataID, ok := m["metadata_id"].(int); ok && metadataID != 0 {
			dtUserActionAndSessionProperties.SetMetadataId(int32(metadataID))
		}

		if ignoreCase, ok := m["ignore_case"].(bool); ok {
			dtUserActionAndSessionProperties.SetIgnoreCase(ignoreCase)
		}
		sps[i] = dtUserActionAndSessionProperties
	}

	return sps
}

func expanduserActionNamingSettings(namingSettings []interface{}) dynatraceConfigV1.UserActionNamingSettings {
	if len(namingSettings) == 0 || namingSettings[0] == nil {
		return dynatraceConfigV1.UserActionNamingSettings{}
	}

	dtUserActionNamingSettings := dynatraceConfigV1.NewUserActionNamingSettingsWithDefaults()

	m := namingSettings[0].(map[string]interface{})

	if ignoreCase, ok := m["ignore_case"].(bool); ok {
		dtUserActionNamingSettings.SetIgnoreCase(ignoreCase)
	}

	if useFirstDetectedLoadAction, ok := m["use_first_detected_load_action"].(bool); ok {
		dtUserActionNamingSettings.SetUseFirstDetectedLoadAction(useFirstDetectedLoadAction)
	}

	if splitUserActionsByDomain, ok := m["split_user_actions_by_domain"].(bool); ok {
		dtUserActionNamingSettings.SetSplitUserActionsByDomain(splitUserActionsByDomain)
	}

	if placeholders, ok := m["placeholder"].([]interface{}); ok && len(placeholders) != 0 {
		dtUserActionNamingSettings.SetPlaceholders(expandPlaceholders(placeholders))
	}

	if loadActionNamingRules, ok := m["load_action_naming_rule"].([]interface{}); ok && len(loadActionNamingRules) != 0 {
		dtUserActionNamingSettings.SetLoadActionNamingRules(expandUserActionNamingRule(loadActionNamingRules))
	}

	if xhrActionNamingRules, ok := m["xhr_action_naming_rule"].([]interface{}); ok && len(xhrActionNamingRules) != 0 {
		dtUserActionNamingSettings.SetXhrActionNamingRules(expandUserActionNamingRule(xhrActionNamingRules))
	}

	if customActionNamingRules, ok := m["custom_action_naming_rule"].([]interface{}); ok && len(customActionNamingRules) != 0 {
		dtUserActionNamingSettings.SetCustomActionNamingRules(expandUserActionNamingRule(customActionNamingRules))
	}

	return *dtUserActionNamingSettings

}

func expandPlaceholders(placeholders []interface{}) []dynatraceConfigV1.UserActionNamingPlaceholder {
	if len(placeholders) < 1 {
		return []dynatraceConfigV1.UserActionNamingPlaceholder{}
	}

	phs := make([]dynatraceConfigV1.UserActionNamingPlaceholder, len(placeholders))

	for i, placeholder := range placeholders {

		m := placeholder.(map[string]interface{})

		var dtPlaceholder dynatraceConfigV1.UserActionNamingPlaceholder

		if name, ok := m["name"].(string); ok {
			dtPlaceholder.SetName(name)
		}

		if input, ok := m["input"].(string); ok {
			dtPlaceholder.SetInput(input)
		}

		if processingPart, ok := m["processing_part"].(string); ok {
			dtPlaceholder.SetProcessingPart(processingPart)
		}

		if metadataID, ok := m["metadata_id"].(int); ok && metadataID != 0 {
			dtPlaceholder.SetMetadataId(int32(metadataID))
		}

		if useGuessedElementIdentifier, ok := m["use_guessed_element_identifier"].(bool); ok {
			dtPlaceholder.SetUseGuessedElementIdentifier(useGuessedElementIdentifier)
		}

		if processingSteps, ok := m["processing_steps"].([]interface{}); ok && len(processingSteps) != 0 {
			dtPlaceholder.SetProcessingSteps(expandProcessingSteps(processingSteps))
		}
		phs[i] = dtPlaceholder
	}

	return phs

}

func expandProcessingSteps(steps []interface{}) []dynatraceConfigV1.UserActionNamingPlaceholderProcessingStep {
	if len(steps) < 1 {
		return []dynatraceConfigV1.UserActionNamingPlaceholderProcessingStep{}
	}

	dps := make([]dynatraceConfigV1.UserActionNamingPlaceholderProcessingStep, len(steps))

	for i, step := range steps {

		m := step.(map[string]interface{})

		var dtActionNamingPlaceholderProcessingStep dynatraceConfigV1.UserActionNamingPlaceholderProcessingStep

		if psType, ok := m["type"].(string); ok {
			dtActionNamingPlaceholderProcessingStep.SetType(psType)
		}

		if patternBefore, ok := m["pattern_before"].(string); ok && len(patternBefore) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetPatternBefore(patternBefore)
		}

		if patternBeforeSearchType, ok := m["pattern_before_search_type"].(string); ok && len(patternBeforeSearchType) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetPatternBeforeSearchType(patternBeforeSearchType)
		}

		if patternAfter, ok := m["pattern_after"].(string); ok && len(patternAfter) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetPatternAfter(patternAfter)
		}

		if patternAfterSearchType, ok := m["pattern_after_search_type"].(string); ok && len(patternAfterSearchType) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetPatternAfterSearchType(patternAfterSearchType)
		}

		if replacement, ok := m["replacement"].(string); ok && len(replacement) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetReplacement(replacement)
		}

		if patternToReplace, ok := m["pattern_to_replace"].(string); ok && len(patternToReplace) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetPatternToReplace(patternToReplace)
		}

		if regularExpression, ok := m["regular_expression"].(string); ok && len(regularExpression) != 0 {
			dtActionNamingPlaceholderProcessingStep.SetRegularExpression(regularExpression)
		}

		if fallbackToInput, ok := m["fallback_to_input"].(bool); ok {
			dtActionNamingPlaceholderProcessingStep.SetFallbackToInput(fallbackToInput)
		}
		dps[i] = dtActionNamingPlaceholderProcessingStep
	}

	return dps

}

func expandUserActionNamingRule(rules []interface{}) []dynatraceConfigV1.UserActionNamingRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.UserActionNamingRule{}
	}

	nrs := make([]dynatraceConfigV1.UserActionNamingRule, len(rules))

	for i, rule := range rules {

		m := rule.(map[string]interface{})

		var dtUserActionNamingRule dynatraceConfigV1.UserActionNamingRule

		if template, ok := m["template"].(string); ok && len(template) != 0 {
			dtUserActionNamingRule.SetTemplate(template)
		}

		if conditions, ok := m["condition"].([]interface{}); ok && len(conditions) != 0 {
			dtUserActionNamingRule.SetConditions(expandNamingRuleConditions(conditions))
		}
		nrs[i] = dtUserActionNamingRule
	}

	return nrs
}

func expandNamingRuleConditions(conditions []interface{}) []dynatraceConfigV1.UserActionNamingRuleCondition {
	if len(conditions) < 1 {
		return []dynatraceConfigV1.UserActionNamingRuleCondition{}
	}

	rcs := make([]dynatraceConfigV1.UserActionNamingRuleCondition, len(conditions))

	for i, condition := range conditions {

		m := condition.(map[string]interface{})

		var dtUserActionNamingRuleCondition dynatraceConfigV1.UserActionNamingRuleCondition

		if operand1, ok := m["operand1"].(string); ok {
			dtUserActionNamingRuleCondition.SetOperand1(operand1)
		}

		if operand2, ok := m["operand2"].(string); ok && len(operand2) != 0 {
			dtUserActionNamingRuleCondition.SetOperand2(operand2)
		}

		if operator, ok := m["operator"].(string); ok {
			dtUserActionNamingRuleCondition.SetOperator(operator)
		}
		rcs[i] = dtUserActionNamingRuleCondition
	}

	return rcs
}

func expandApdexSettings(apdexSettings []interface{}) dynatraceConfigV1.Apdex {
	if len(apdexSettings) == 0 || apdexSettings[0] == nil {
		return dynatraceConfigV1.Apdex{}
	}

	dtapdexSettings := dynatraceConfigV1.NewApdexWithDefaults()

	m := apdexSettings[0].(map[string]interface{})

	if threshold, ok := m["threshold"].(int); ok && threshold != 0 {
		dtapdexSettings.SetThreshold(float32(threshold))
	}

	if toleratedThreshold, ok := m["tolerated_threshold"].(int); ok && toleratedThreshold != 0 {
		dtapdexSettings.SetToleratedThreshold(int32(toleratedThreshold))
	}

	if frustratingThreshold, ok := m["frustrating_threshold"].(int); ok && frustratingThreshold != 0 {
		dtapdexSettings.SetFrustratingThreshold(int32(frustratingThreshold))
	}

	if toleratedFallbackThreshold, ok := m["tolerated_fallback_threshold"].(int); ok && toleratedFallbackThreshold != 0 {
		dtapdexSettings.SetToleratedFallbackThreshold(int32(toleratedFallbackThreshold))
	}

	if frustratingFallbackThreshold, ok := m["frustrating_fallback_threshold"].(int); ok && frustratingFallbackThreshold != 0 {
		dtapdexSettings.SetFrustratingFallbackThreshold(int32(frustratingFallbackThreshold))
	}

	return *dtapdexSettings

}

func expandWaterfallSettings(waterfallSettings []interface{}) dynatraceConfigV1.WaterfallSettings {
	if len(waterfallSettings) == 0 || waterfallSettings[0] == nil {
		return dynatraceConfigV1.WaterfallSettings{}
	}

	dtWaterfallSettings := dynatraceConfigV1.NewWaterfallSettingsWithDefaults()

	m := waterfallSettings[0].(map[string]interface{})

	if uncompressedResourcesThreshold, ok := m["uncompressed_resources_threshold"].(int); ok {
		dtWaterfallSettings.SetUncompressedResourcesThreshold(int32(uncompressedResourcesThreshold))
	}

	if resourcesThreshold, ok := m["resources_threshold"].(int); ok {
		dtWaterfallSettings.SetResourcesThreshold(int32(resourcesThreshold))
	}

	if resourceBrowserCachingThreshold, ok := m["resources_browser_caching_threshold"].(int); ok {
		dtWaterfallSettings.SetResourceBrowserCachingThreshold(int32(resourceBrowserCachingThreshold))
	}

	if slowFirstPartyResourcesThreshold, ok := m["slow_first_party_resources_threshold"].(int); ok {
		dtWaterfallSettings.SetSlowFirstPartyResourcesThreshold(int32(slowFirstPartyResourcesThreshold))
	}

	if slowThirdPartyResourcesThreshold, ok := m["slow_third_party_resources_threshold"].(int); ok {
		dtWaterfallSettings.SetSlowThirdPartyResourcesThreshold(int32(slowThirdPartyResourcesThreshold))
	}

	if slowCdnResourcesThreshold, ok := m["slow_cdn_resources_threshold"].(int); ok {
		dtWaterfallSettings.SetSlowCdnResourcesThreshold(int32(slowCdnResourcesThreshold))
	}

	if speedIndexVisuallyCompleteRatioThreshold, ok := m["speed_index_visually_complete_ratio_threshold"].(int); ok {
		dtWaterfallSettings.SetSpeedIndexVisuallyCompleteRatioThreshold(int32(speedIndexVisuallyCompleteRatioThreshold))
	}

	return *dtWaterfallSettings

}

func expandMonitoringSettings(monitoringSettings []interface{}) dynatraceConfigV1.MonitoringSettings {
	if len(monitoringSettings) == 0 || monitoringSettings[0] == nil {
		return dynatraceConfigV1.MonitoringSettings{}
	}

	dtMonitoringSettings := dynatraceConfigV1.NewMonitoringSettingsWithDefaults()

	m := monitoringSettings[0].(map[string]interface{})

	if fetchRequests, ok := m["fetch_requests"].(bool); ok {
		dtMonitoringSettings.SetFetchRequests(fetchRequests)
	}

	if xmlHTTPRequest, ok := m["xml_http_request"].(bool); ok {
		dtMonitoringSettings.SetXmlHttpRequest(xmlHTTPRequest)
	}

	if excludeXhrRegex, ok := m["exclude_xhr_regex"].(string); ok {
		dtMonitoringSettings.SetExcludeXhrRegex(excludeXhrRegex)
	}

	if correlationHeaderInclusionRegex, ok := m["correlation_header_inclusion_regex"].(string); ok && len(correlationHeaderInclusionRegex) != 0 {
		dtMonitoringSettings.SetCorrelationHeaderInclusionRegex(correlationHeaderInclusionRegex)
	}

	if injectionMode, ok := m["injection_mode"].(string); ok {
		dtMonitoringSettings.SetInjectionMode(injectionMode)
	}

	if addCrossOriginAnonymousAttribute, ok := m["add_cross_origin_anonymous_attribute"].(bool); ok {
		dtMonitoringSettings.SetAddCrossOriginAnonymousAttribute(addCrossOriginAnonymousAttribute)
	}

	if scriptTagCacheDurationInHours, ok := m["script_tag_cache_duration_in_hours"].(int); ok {
		dtMonitoringSettings.SetScriptTagCacheDurationInHours(int32(scriptTagCacheDurationInHours))
	}

	if libraryFileLocation, ok := m["library_file_location"].(string); ok {
		dtMonitoringSettings.SetLibraryFileLocation(libraryFileLocation)
	}

	if monitoringDataPath, ok := m["monitoring_data_path"].(string); ok {
		dtMonitoringSettings.SetMonitoringDataPath(monitoringDataPath)
	}

	if customConfigurationProperties, ok := m["custom_configuration_properties"].(string); ok {
		dtMonitoringSettings.SetCustomConfigurationProperties(customConfigurationProperties)
	}

	if serverRequestPathID, ok := m["server_request_path_id"].(string); ok {
		dtMonitoringSettings.SetServerRequestPathId(serverRequestPathID)
	}

	if secureCookieAttribute, ok := m["secure_cookie_attribute"].(bool); ok {
		dtMonitoringSettings.SetSecureCookieAttribute(secureCookieAttribute)
	}

	if cookiePlacementDomain, ok := m["cookie_placement_domain"].(string); ok {
		dtMonitoringSettings.SetCookiePlacementDomain(cookiePlacementDomain)
	}

	if cacheControlHeaderOptimizations, ok := m["cache_control_header_optimizations"].(bool); ok {
		dtMonitoringSettings.SetCacheControlHeaderOptimizations(cacheControlHeaderOptimizations)
	}

	if javaScriptFrameworkSupport, ok := m["javascript_framework_support"].([]interface{}); ok {
		dtMonitoringSettings.SetJavaScriptFrameworkSupport(expandJavaScriptFrameworkSupport(javaScriptFrameworkSupport))
	}

	if contentCapture, ok := m["content_capture"].([]interface{}); ok {
		dtMonitoringSettings.SetContentCapture(expandContentCapture(contentCapture))
	}

	if advancedJavaScriptTagSettings, ok := m["advanced_javascript_tag_settings"].([]interface{}); ok {
		dtMonitoringSettings.SetAdvancedJavaScriptTagSettings(expandAdvancedJavaScriptTagSettings(advancedJavaScriptTagSettings))
	}

	if browserRestrictionSettings, ok := m["browser_restriction_settings"].([]interface{}); ok && len(browserRestrictionSettings) != 0 {
		dtMonitoringSettings.SetBrowserRestrictionSettings(expandBrowserRestrictionSettings(browserRestrictionSettings))
	}

	if ipAddressRestrictionSettings, ok := m["ip_address_restriction_settings"].([]interface{}); ok && len(ipAddressRestrictionSettings) != 0 {
		dtMonitoringSettings.SetIpAddressRestrictionSettings(expandIPAddressRestrictionSettings(ipAddressRestrictionSettings))
	}

	if javaScriptInjectionRules, ok := m["javascript_injection_rule"].([]interface{}); ok && len(javaScriptInjectionRules) != 0 {
		dtMonitoringSettings.SetJavaScriptInjectionRules(expandJavaScriptInjectionRules(javaScriptInjectionRules))
	}

	return *dtMonitoringSettings

}

func expandJavaScriptFrameworkSupport(fwSupport []interface{}) dynatraceConfigV1.JavaScriptFrameworkSupport {
	if len(fwSupport) == 0 || fwSupport[0] == nil {
		return dynatraceConfigV1.JavaScriptFrameworkSupport{}
	}

	dtJavaScriptFrameworkSupport := dynatraceConfigV1.NewJavaScriptFrameworkSupportWithDefaults()

	m := fwSupport[0].(map[string]interface{})

	if angular, ok := m["angular"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetAngular(angular)
	}

	if dojo, ok := m["dojo"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetDojo(dojo)
	}

	if extJs, ok := m["ext_js"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetExtJS(extJs)
	}

	if icefaces, ok := m["icefaces"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetIcefaces(icefaces)
	}

	if jQuery, ok := m["jquery"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetJQuery(jQuery)
	}

	if mooTools, ok := m["moo_tools"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetMooTools(mooTools)
	}

	if prototype, ok := m["prototype"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetPrototype(prototype)
	}

	if activeXObject, ok := m["activex_object"].(bool); ok {
		dtJavaScriptFrameworkSupport.SetActiveXObject(activeXObject)
	}

	return *dtJavaScriptFrameworkSupport

}

func expandContentCapture(contentCapture []interface{}) dynatraceConfigV1.ContentCapture {
	if len(contentCapture) == 0 || contentCapture[0] == nil {
		return dynatraceConfigV1.ContentCapture{}
	}

	dtContentCapture := dynatraceConfigV1.NewContentCaptureWithDefaults()

	m := contentCapture[0].(map[string]interface{})

	if javaScriptErrors, ok := m["javascript_errors"].(bool); ok {
		dtContentCapture.SetJavaScriptErrors(javaScriptErrors)
	}

	if visuallyCompleteAndSpeedIndex, ok := m["visually_complete_and_speed_index"].(bool); ok {
		dtContentCapture.SetVisuallyCompleteAndSpeedIndex(visuallyCompleteAndSpeedIndex)
	}

	if resourceTimingSettings, ok := m["resource_timing_settings"].([]interface{}); ok {
		dtContentCapture.SetResourceTimingSettings(expandResourceTimingSettings(resourceTimingSettings))
	}

	if timeoutSettings, ok := m["timeout_settings"].([]interface{}); ok {
		dtContentCapture.SetTimeoutSettings(expandTimeoutSettings(timeoutSettings))
	}

	if visuallyComplete2Settings, ok := m["visually_complete_2_settings"].([]interface{}); ok && len(visuallyComplete2Settings) != 0 {
		dtContentCapture.SetVisuallyComplete2Settings(expandVisuallyComplete2Settings(visuallyComplete2Settings))
	}

	return *dtContentCapture

}

func expandResourceTimingSettings(timingSettings []interface{}) dynatraceConfigV1.ResourceTimingSettings {
	if len(timingSettings) == 0 || timingSettings[0] == nil {
		return dynatraceConfigV1.ResourceTimingSettings{}
	}

	dtResourceTimingSettings := dynatraceConfigV1.NewResourceTimingSettingsWithDefaults()

	m := timingSettings[0].(map[string]interface{})

	if w3cResourceTimings, ok := m["w3c_resource_timings"].(bool); ok {
		dtResourceTimingSettings.SetW3cResourceTimings(w3cResourceTimings)
	}

	if nonW3cResourceTimings, ok := m["non_w3c_resource_timings"].(bool); ok {
		dtResourceTimingSettings.SetNonW3cResourceTimings(nonW3cResourceTimings)
	}

	if nonW3cResourceTimingsInstrumentationDelay, ok := m["non_w3c_resource_timings_instrumentation_delay"].(int); ok {
		dtResourceTimingSettings.SetNonW3cResourceTimingsInstrumentationDelay(int32(nonW3cResourceTimingsInstrumentationDelay))
	}

	if resourceTimingCaptureType, ok := m["resource_timing_capture_type"].(string); ok {
		dtResourceTimingSettings.SetResourceTimingCaptureType(resourceTimingCaptureType)
	}

	if resourceTimingsDomainLimit, ok := m["resource_timings_domain_limit"].(int); ok {
		dtResourceTimingSettings.SetResourceTimingsDomainLimit(int32(resourceTimingsDomainLimit))
	}

	return *dtResourceTimingSettings

}

func expandTimeoutSettings(timeoutSettings []interface{}) dynatraceConfigV1.TimeoutSettings {
	if len(timeoutSettings) == 0 || timeoutSettings[0] == nil {
		return dynatraceConfigV1.TimeoutSettings{}
	}

	dtTimeoutSettings := dynatraceConfigV1.NewTimeoutSettingsWithDefaults()

	m := timeoutSettings[0].(map[string]interface{})

	if timedActionSupport, ok := m["timed_action_support"].(bool); ok {
		dtTimeoutSettings.SetTimedActionSupport(timedActionSupport)
	}

	if temporaryActionLimit, ok := m["temporary_action_limit"].(int); ok {
		dtTimeoutSettings.SetTemporaryActionLimit(int32(temporaryActionLimit))
	}

	if temporaryActionTotalTimeout, ok := m["temporary_action_total_timeout"].(int); ok {
		dtTimeoutSettings.SetTemporaryActionTotalTimeout(int32(temporaryActionTotalTimeout))
	}

	return *dtTimeoutSettings

}

func expandVisuallyComplete2Settings(visuallyComplete2Settings []interface{}) dynatraceConfigV1.VisuallyComplete2Settings {
	if len(visuallyComplete2Settings) == 0 || visuallyComplete2Settings[0] == nil {
		return dynatraceConfigV1.VisuallyComplete2Settings{}
	}

	dtVisuallyComplete2Settings := dynatraceConfigV1.NewVisuallyComplete2SettingsWithDefaults()

	m := visuallyComplete2Settings[0].(map[string]interface{})

	if imageURLBlacklist, ok := m["image_url_blacklist"].(string); ok {
		dtVisuallyComplete2Settings.SetImageUrlBlacklist(imageURLBlacklist)
	}

	if mutationBlacklist, ok := m["mutation_blacklist"].(string); ok {
		dtVisuallyComplete2Settings.SetMutationBlacklist(mutationBlacklist)
	}

	if mutationTimeout, ok := m["mutation_timeout"].(int); ok {
		dtVisuallyComplete2Settings.SetMutationTimeout(int32(mutationTimeout))
	}

	if inactivityTimeout, ok := m["inactivity_timeout"].(int); ok {
		dtVisuallyComplete2Settings.SetInactivityTimeout(int32(inactivityTimeout))
	}

	if threshold, ok := m["threshold"].(int); ok {
		dtVisuallyComplete2Settings.SetThreshold(int32(threshold))
	}

	return *dtVisuallyComplete2Settings

}

func expandAdvancedJavaScriptTagSettings(jsTagSettings []interface{}) dynatraceConfigV1.AdvancedJavaScriptTagSettings {
	if len(jsTagSettings) == 0 || jsTagSettings[0] == nil {
		return dynatraceConfigV1.AdvancedJavaScriptTagSettings{}
	}

	dtAdvancedJavaScriptTagSettings := dynatraceConfigV1.NewAdvancedJavaScriptTagSettingsWithDefaults()

	m := jsTagSettings[0].(map[string]interface{})

	if syncBeaconFirefox, ok := m["sync_beacon_firefox"].(bool); ok {
		dtAdvancedJavaScriptTagSettings.SetSyncBeaconFirefox(syncBeaconFirefox)
	}

	if syncBeaconInternetExplorer, ok := m["sync_beacon_internet_explorer"].(bool); ok {
		dtAdvancedJavaScriptTagSettings.SetSyncBeaconInternetExplorer(syncBeaconInternetExplorer)
	}

	if instrumentUnsupportedAjaxFrameworks, ok := m["instrument_unsupported_ajax_frameworks"].(bool); ok {
		dtAdvancedJavaScriptTagSettings.SetInstrumentUnsupportedAjaxFrameworks(instrumentUnsupportedAjaxFrameworks)
	}

	if specialCharactersToEscape, ok := m["special_characters_to_escape"].(string); ok {
		dtAdvancedJavaScriptTagSettings.SetSpecialCharactersToEscape(specialCharactersToEscape)
	}

	if maxActionNameLength, ok := m["max_action_name_length"].(int); ok {
		dtAdvancedJavaScriptTagSettings.SetMaxActionNameLength(int32(maxActionNameLength))
	}

	if maxErrorsToCapture, ok := m["max_errors_to_capture"].(int); ok {
		dtAdvancedJavaScriptTagSettings.SetMaxErrorsToCapture(int32(maxErrorsToCapture))
	}

	if additionalEventHandlers, ok := m["additional_event_handlers"].([]interface{}); ok {
		dtAdvancedJavaScriptTagSettings.SetAdditionalEventHandlers(expandAdditionalEventHandlers(additionalEventHandlers))
	}

	if eventWrapperSettings, ok := m["event_wrapper_settings"].([]interface{}); ok {
		dtAdvancedJavaScriptTagSettings.SetEventWrapperSettings(expandEventWrapperSettings(eventWrapperSettings))
	}

	if globalEventCaptureSettings, ok := m["global_event_capture_settings"].([]interface{}); ok {
		dtAdvancedJavaScriptTagSettings.SetGlobalEventCaptureSettings(expandGlobalEventCaptureSettings(globalEventCaptureSettings))
	}

	return *dtAdvancedJavaScriptTagSettings

}

func expandAdditionalEventHandlers(eventHandlers []interface{}) dynatraceConfigV1.AdditionalEventHandlers {
	if len(eventHandlers) == 0 || eventHandlers[0] == nil {
		return dynatraceConfigV1.AdditionalEventHandlers{}
	}

	dtAdditionalEventHandlers := dynatraceConfigV1.NewAdditionalEventHandlersWithDefaults()

	m := eventHandlers[0].(map[string]interface{})

	if userMouseupEventForClicks, ok := m["user_mouseup_event_for_clicks"].(bool); ok {
		dtAdditionalEventHandlers.SetUserMouseupEventForClicks(userMouseupEventForClicks)
	}

	if clickEventHandler, ok := m["click_event_handler"].(bool); ok {
		dtAdditionalEventHandlers.SetBlurEventHandler(clickEventHandler)
	}

	if mouseupEventHandler, ok := m["mouseup_event_handler"].(bool); ok {
		dtAdditionalEventHandlers.SetMouseupEventHandler(mouseupEventHandler)
	}

	if blurEventHandler, ok := m["blur_event_handler"].(bool); ok {
		dtAdditionalEventHandlers.SetBlurEventHandler(blurEventHandler)
	}

	if changeEventHandler, ok := m["change_event_handler"].(bool); ok {
		dtAdditionalEventHandlers.SetChangeEventHandler(changeEventHandler)
	}

	if toStringMethod, ok := m["to_string_method"].(bool); ok {
		dtAdditionalEventHandlers.SetToStringMethod(toStringMethod)
	}

	if maxDomNodesToInstrument, ok := m["max_dom_nodes_to_instrument"].(int); ok {
		dtAdditionalEventHandlers.SetMaxDomNodesToInstrument(int32(maxDomNodesToInstrument))
	}

	return *dtAdditionalEventHandlers

}

func expandEventWrapperSettings(wrapperSettings []interface{}) dynatraceConfigV1.EventWrapperSettings {
	if len(wrapperSettings) == 0 || wrapperSettings[0] == nil {
		return dynatraceConfigV1.EventWrapperSettings{}
	}

	dtEventWrapperSettings := dynatraceConfigV1.NewEventWrapperSettingsWithDefaults()

	m := wrapperSettings[0].(map[string]interface{})

	if click, ok := m["click"].(bool); ok {
		dtEventWrapperSettings.SetClick(click)
	}

	if mouseUp, ok := m["mouse_up"].(bool); ok {
		dtEventWrapperSettings.SetMouseUp(mouseUp)
	}

	if change, ok := m["change"].(bool); ok {
		dtEventWrapperSettings.SetChange(change)
	}

	if blur, ok := m["blur"].(bool); ok {
		dtEventWrapperSettings.SetBlur(blur)
	}

	if touchStart, ok := m["touch_start"].(bool); ok {
		dtEventWrapperSettings.SetTouchStart(touchStart)
	}

	if touchEnd, ok := m["touch_end"].(bool); ok {
		dtEventWrapperSettings.SetTouchEnd(touchEnd)
	}

	return *dtEventWrapperSettings

}

func expandGlobalEventCaptureSettings(captureSettings []interface{}) dynatraceConfigV1.GlobalEventCaptureSettings {
	if len(captureSettings) == 0 || captureSettings[0] == nil {
		return dynatraceConfigV1.GlobalEventCaptureSettings{}
	}

	dtGlobalEventCaptureSettings := dynatraceConfigV1.NewGlobalEventCaptureSettingsWithDefaults()

	m := captureSettings[0].(map[string]interface{})

	if mouseUp, ok := m["mouse_up"].(bool); ok {
		dtGlobalEventCaptureSettings.SetMouseUp(mouseUp)
	}

	if mouseDown, ok := m["mouse_down"].(bool); ok {
		dtGlobalEventCaptureSettings.SetMouseDown(mouseDown)
	}

	if click, ok := m["click"].(bool); ok {
		dtGlobalEventCaptureSettings.SetClick(click)
	}

	if doubleClick, ok := m["double_click"].(bool); ok {
		dtGlobalEventCaptureSettings.SetDoubleClick(doubleClick)
	}

	if keyUp, ok := m["key_up"].(bool); ok {
		dtGlobalEventCaptureSettings.SetKeyUp(keyUp)
	}

	if keyDown, ok := m["key_down"].(bool); ok {
		dtGlobalEventCaptureSettings.SetKeyDown(keyDown)
	}

	if scroll, ok := m["scroll"].(bool); ok {
		dtGlobalEventCaptureSettings.SetScroll(scroll)
	}

	if additionalEventCapturedAsUserInput, ok := m["additional_event_captured_as_user_input"].(string); ok {
		dtGlobalEventCaptureSettings.SetAdditionalEventCapturedAsUserInput(additionalEventCapturedAsUserInput)
	}

	return *dtGlobalEventCaptureSettings

}

func expandBrowserRestrictionSettings(restrictionSettings []interface{}) dynatraceConfigV1.WebApplicationConfigBrowserRestrictionSettings {
	if len(restrictionSettings) == 0 || restrictionSettings[0] == nil {
		return dynatraceConfigV1.WebApplicationConfigBrowserRestrictionSettings{}
	}

	dtWebApplicationConfigBrowserRestrictionSettings := dynatraceConfigV1.NewWebApplicationConfigBrowserRestrictionSettingsWithDefaults()

	m := restrictionSettings[0].(map[string]interface{})

	if mode, ok := m["mode"].(string); ok {
		dtWebApplicationConfigBrowserRestrictionSettings.SetMode(mode)
	}

	if browserRestrictions, ok := m["browser_restriction"].([]interface{}); ok && len(browserRestrictions) != 0 {
		dtWebApplicationConfigBrowserRestrictionSettings.SetBrowserRestrictions(expandBrowserRestrictions(browserRestrictions))
	}

	return *dtWebApplicationConfigBrowserRestrictionSettings
}

func expandBrowserRestrictions(restrictions []interface{}) []dynatraceConfigV1.WebApplicationConfigBrowserRestriction {
	if len(restrictions) < 1 {
		return []dynatraceConfigV1.WebApplicationConfigBrowserRestriction{}
	}

	brs := make([]dynatraceConfigV1.WebApplicationConfigBrowserRestriction, len(restrictions))

	for i, restriction := range restrictions {

		m := restriction.(map[string]interface{})

		var dtWebApplicationConfigBrowserRestriction dynatraceConfigV1.WebApplicationConfigBrowserRestriction

		if browserVersion, ok := m["browser_version"].(string); ok && len(browserVersion) != 0 {
			dtWebApplicationConfigBrowserRestriction.SetBrowserVersion(browserVersion)
		}

		if browserType, ok := m["browser_type"].(string); ok {
			dtWebApplicationConfigBrowserRestriction.SetBrowserType(browserType)
		}

		if platform, ok := m["platform"].(string); ok && len(platform) != 0 {
			dtWebApplicationConfigBrowserRestriction.SetPlatform(platform)
		}

		if comparator, ok := m["comparator"].(string); ok && len(comparator) != 0 {
			dtWebApplicationConfigBrowserRestriction.SetComparator(comparator)
		}
		brs[i] = dtWebApplicationConfigBrowserRestriction
	}

	return brs

}

func expandIPAddressRestrictionSettings(restrictionSettings []interface{}) dynatraceConfigV1.WebApplicationConfigIpAddressRestrictionSettings {
	if len(restrictionSettings) == 0 || restrictionSettings[0] == nil {
		return dynatraceConfigV1.WebApplicationConfigIpAddressRestrictionSettings{}
	}

	dtWebApplicationConfigIPAddressRestrictionSettings := dynatraceConfigV1.NewWebApplicationConfigIpAddressRestrictionSettingsWithDefaults()

	m := restrictionSettings[0].(map[string]interface{})

	if mode, ok := m["mode"].(string); ok {
		dtWebApplicationConfigIPAddressRestrictionSettings.SetMode(mode)
	}

	if ipAddressRestrictions, ok := m["ip_address_restriction"].([]interface{}); ok && len(ipAddressRestrictions) != 0 {
		dtWebApplicationConfigIPAddressRestrictionSettings.SetIpAddressRestrictions(expandIPAddressRestrictions(ipAddressRestrictions))
	}

	return *dtWebApplicationConfigIPAddressRestrictionSettings

}

func expandIPAddressRestrictions(restrictions []interface{}) []dynatraceConfigV1.IpAddressRange {
	if len(restrictions) < 1 {
		return []dynatraceConfigV1.IpAddressRange{}
	}

	irs := make([]dynatraceConfigV1.IpAddressRange, len(restrictions))

	for i, restriction := range restrictions {

		m := restriction.(map[string]interface{})

		var dtIPAddressRange dynatraceConfigV1.IpAddressRange

		if subnetMask, ok := m["subnet_mask"].(int); ok && subnetMask != 0 {
			dtIPAddressRange.SetSubnetMask(int32(subnetMask))
		}

		if address, ok := m["address"].(string); ok {
			dtIPAddressRange.SetAddress(address)
		}

		if addressTo, ok := m["address_to"].(string); ok && len(addressTo) != 0 {
			dtIPAddressRange.SetAddressTo(addressTo)
		}

		irs[i] = dtIPAddressRange
	}

	return irs
}

func expandJavaScriptInjectionRules(injectionRules []interface{}) []dynatraceConfigV1.JavaScriptInjectionRules {
	if len(injectionRules) < 1 {
		return []dynatraceConfigV1.JavaScriptInjectionRules{}
	}

	jrs := make([]dynatraceConfigV1.JavaScriptInjectionRules, len(injectionRules))

	for i, injectionRule := range injectionRules {

		m := injectionRule.(map[string]interface{})

		var dtJavaScriptInjectionRules dynatraceConfigV1.JavaScriptInjectionRules

		if enabled, ok := m["enabled"].(bool); ok {
			dtJavaScriptInjectionRules.SetEnabled(enabled)
		}

		if urlOperator, ok := m["url_operator"].(string); ok {
			dtJavaScriptInjectionRules.SetUrlOperator(urlOperator)
		}

		if urlPattern, ok := m["url_pattern"].(string); ok && len(urlPattern) != 0 {
			dtJavaScriptInjectionRules.SetUrlPattern(urlPattern)
		}

		if rule, ok := m["rule"].(string); ok {
			dtJavaScriptInjectionRules.SetRule(rule)
		}

		if htmlPattern, ok := m["html_pattern"].(string); ok && len(htmlPattern) != 0 {
			dtJavaScriptInjectionRules.SetHtmlPattern(htmlPattern)
		}

		jrs[i] = dtJavaScriptInjectionRules

	}
	return jrs

}

func flattenDynatraceWebApplication(webApp dynatraceConfigV1.WebApplicationConfig, d *schema.ResourceData) diag.Diagnostics {

	d.Set("name", &webApp.Name)
	d.Set("type", &webApp.Type)
	d.Set("real_user_monitoring_enabled", &webApp.RealUserMonitoringEnabled)
	d.Set("cost_control_user_session_percentage", &webApp.CostControlUserSessionPercentage)
	d.Set("load_action_key_performance_metric", &webApp.LoadActionKeyPerformanceMetric)
	d.Set("xhr_action_key_performance_metric", &webApp.XhrActionKeyPerformanceMetric)
	d.Set("url_injection_pattern", &webApp.UrlInjectionPattern)

	sessionReplayConfig := flattenSessionReplayConfig(webApp.SessionReplayConfig)
	if err := d.Set("session_replay_config", sessionReplayConfig); err != nil {
		return diag.FromErr(err)
	}

	loadActionApdexSettings := flattenApdexSettings(&webApp.LoadActionApdexSettings)
	if err := d.Set("load_action_apdex_settings", loadActionApdexSettings); err != nil {
		return diag.FromErr(err)
	}

	xhrActionApdexSettings := flattenApdexSettings(&webApp.XhrActionApdexSettings)
	if err := d.Set("xhr_action_apdex_settings", xhrActionApdexSettings); err != nil {
		return diag.FromErr(err)
	}

	customActionApdexSettings := flattenApdexSettings(&webApp.CustomActionApdexSettings)
	if err := d.Set("custom_action_apdex_settings", customActionApdexSettings); err != nil {
		return diag.FromErr(err)
	}

	waterfallSettings := flattenWaterfallSettings(&webApp.WaterfallSettings)
	if err := d.Set("waterfall_settings", waterfallSettings); err != nil {
		return diag.FromErr(err)
	}

	monitoringSettings := flattenMonitoringSettings(&webApp.MonitoringSettings)
	if err := d.Set("monitoring_settings", monitoringSettings); err != nil {
		return diag.FromErr(err)
	}

	userTags := flattenUserTags(webApp.UserTags)
	if err := d.Set("user_tag", userTags); err != nil {
		return diag.FromErr(err)
	}

	userActionAndSessionProperties := flattenUserActionAndSessionProperties(webApp.UserActionAndSessionProperties)
	if err := d.Set("user_action_and_session_property", userActionAndSessionProperties); err != nil {
		return diag.FromErr(err)
	}

	userActionNamingSettings := flattenUserActionNamingSettings(webApp.UserActionNamingSettings)
	if err := d.Set("user_action_naming_settings", userActionNamingSettings); err != nil {
		return diag.FromErr(err)
	}

	metaDataCaptureSettings := flattenMetaDataCaptureSettings(webApp.MetaDataCaptureSettings)
	if err := d.Set("metadata_capture_setting", metaDataCaptureSettings); err != nil {
		return diag.FromErr(err)
	}

	conversionGoals := flattenConversionGoals(webApp.ConversionGoals)
	if err := d.Set("conversion_goal", conversionGoals); err != nil {
		return diag.FromErr(err)
	}

	return nil

}

func flattenSessionReplayConfig(replayConfig *dynatraceConfigV1.SessionReplaySetting) []interface{} {
	if replayConfig == nil {
		return []interface{}{replayConfig}
	}

	m := make(map[string]interface{})

	m["enabled"] = replayConfig.Enabled
	m["cost_control_percentage"] = replayConfig.CostControlPercentage

	return []interface{}{m}
}

func flattenApdexSettings(apdexSettings *dynatraceConfigV1.Apdex) []interface{} {
	if apdexSettings == nil {
		return []interface{}{apdexSettings}
	}

	a := make(map[string]interface{})

	a["threshold"] = apdexSettings.Threshold
	a["tolerated_threshold"] = apdexSettings.ToleratedThreshold
	a["frustrating_threshold"] = apdexSettings.FrustratingThreshold
	a["tolerated_fallback_threshold"] = apdexSettings.ToleratedFallbackThreshold
	a["frustrating_fallback_threshold"] = apdexSettings.FrustratingFallbackThreshold

	return []interface{}{a}

}

func flattenWaterfallSettings(waterfallSettings *dynatraceConfigV1.WaterfallSettings) []interface{} {
	if waterfallSettings == nil {
		return []interface{}{waterfallSettings}
	}

	w := make(map[string]interface{})

	w["uncompressed_resources_threshold"] = waterfallSettings.UncompressedResourcesThreshold
	w["resources_threshold"] = waterfallSettings.ResourcesThreshold
	w["resources_browser_caching_threshold"] = waterfallSettings.ResourceBrowserCachingThreshold
	w["slow_first_party_resources_threshold"] = waterfallSettings.SlowFirstPartyResourcesThreshold
	w["slow_third_party_resources_threshold"] = waterfallSettings.SlowThirdPartyResourcesThreshold
	w["slow_cdn_resources_threshold"] = waterfallSettings.SlowCdnResourcesThreshold
	w["speed_index_visually_complete_ratio_threshold"] = waterfallSettings.SpeedIndexVisuallyCompleteRatioThreshold

	return []interface{}{w}

}

func flattenMonitoringSettings(monitoringSettings *dynatraceConfigV1.MonitoringSettings) []interface{} {
	if monitoringSettings == nil {
		return []interface{}{monitoringSettings}
	}

	m := make(map[string]interface{})

	m["fetch_requests"] = monitoringSettings.FetchRequests
	m["xml_http_request"] = monitoringSettings.XmlHttpRequest
	m["exclude_xhr_regex"] = monitoringSettings.ExcludeXhrRegex
	m["correlation_header_inclusion_regex"] = monitoringSettings.CorrelationHeaderInclusionRegex
	m["injection_mode"] = monitoringSettings.InjectionMode
	m["add_cross_origin_anonymous_attribute"] = monitoringSettings.AddCrossOriginAnonymousAttribute
	m["script_tag_cache_duration_in_hours"] = monitoringSettings.ScriptTagCacheDurationInHours
	m["library_file_location"] = monitoringSettings.LibraryFileLocation
	m["monitoring_data_path"] = monitoringSettings.MonitoringDataPath
	m["custom_configuration_properties"] = monitoringSettings.CustomConfigurationProperties
	m["server_request_path_id"] = monitoringSettings.ServerRequestPathId
	m["secure_cookie_attribute"] = monitoringSettings.SecureCookieAttribute
	m["cookie_placement_domain"] = monitoringSettings.CookiePlacementDomain
	m["cache_control_header_optimizations"] = monitoringSettings.CacheControlHeaderOptimizations
	m["javascript_framework_support"] = flattenJavascriptFrameworkSupport(&monitoringSettings.JavaScriptFrameworkSupport)
	m["content_capture"] = flattenContentCapture(&monitoringSettings.ContentCapture)
	m["advanced_javascript_tag_settings"] = flattenAdvancedJavaScriptTagSettings(&monitoringSettings.AdvancedJavaScriptTagSettings)
	m["browser_restriction_settings"] = flattenBrowserRestrictionSettings(monitoringSettings.BrowserRestrictionSettings)
	m["ip_address_restriction_settings"] = flattenIPAddressRestrictionSettings(monitoringSettings.IpAddressRestrictionSettings)
	m["javascript_injection_rule"] = flattenJavaScriptInjectionRules(monitoringSettings.JavaScriptInjectionRules)

	return []interface{}{m}

}

func flattenJavascriptFrameworkSupport(frameworkSupport *dynatraceConfigV1.JavaScriptFrameworkSupport) []interface{} {
	if frameworkSupport == nil {
		return []interface{}{frameworkSupport}
	}

	m := make(map[string]interface{})

	m["angular"] = frameworkSupport.Angular
	m["dojo"] = frameworkSupport.Dojo
	m["ext_js"] = frameworkSupport.ExtJS
	m["icefaces"] = frameworkSupport.Icefaces
	m["jquery"] = frameworkSupport.JQuery
	m["moo_tools"] = frameworkSupport.MooTools
	m["prototype"] = frameworkSupport.Prototype
	m["activex_object"] = frameworkSupport.ActiveXObject

	return []interface{}{m}
}

func flattenContentCapture(contentCapture *dynatraceConfigV1.ContentCapture) []interface{} {
	if contentCapture == nil {
		return []interface{}{contentCapture}
	}

	m := make(map[string]interface{})

	m["javascript_errors"] = contentCapture.JavaScriptErrors
	m["visually_complete_and_speed_index"] = contentCapture.VisuallyCompleteAndSpeedIndex
	m["resource_timing_settings"] = flattenResourceTimingSettings(&contentCapture.ResourceTimingSettings)
	m["timeout_settings"] = flattenTimeoutSettings(&contentCapture.TimeoutSettings)
	m["visually_complete_2_settings"] = flattenVisuallyComplete2Settings(contentCapture.VisuallyComplete2Settings)

	return []interface{}{m}
}

func flattenResourceTimingSettings(timingSettings *dynatraceConfigV1.ResourceTimingSettings) []interface{} {
	if timingSettings == nil {
		return []interface{}{timingSettings}
	}

	m := make(map[string]interface{})

	m["w3c_resource_timings"] = timingSettings.W3cResourceTimings
	m["non_w3c_resource_timings"] = timingSettings.NonW3cResourceTimings
	m["non_w3c_resource_timings_instrumentation_delay"] = timingSettings.NonW3cResourceTimingsInstrumentationDelay
	m["resource_timing_capture_type"] = timingSettings.ResourceTimingCaptureType
	m["resource_timings_domain_limit"] = timingSettings.ResourceTimingsDomainLimit

	return []interface{}{m}

}

func flattenVisuallyComplete2Settings(complete2Settings *dynatraceConfigV1.VisuallyComplete2Settings) []interface{} {
	if complete2Settings == nil {
		return []interface{}{complete2Settings}
	}

	m := make(map[string]interface{})

	m["image_url_blacklist"] = complete2Settings.ImageUrlBlacklist
	m["mutation_blacklist"] = complete2Settings.MutationBlacklist
	m["mutation_timeout"] = complete2Settings.MutationTimeout
	m["inactivity_timeout"] = complete2Settings.InactivityTimeout
	m["threshold"] = complete2Settings.Threshold

	return []interface{}{m}

}

func flattenTimeoutSettings(timeoutSettings *dynatraceConfigV1.TimeoutSettings) []interface{} {
	if timeoutSettings == nil {
		return []interface{}{timeoutSettings}
	}

	m := make(map[string]interface{})

	m["timed_action_support"] = timeoutSettings.TimedActionSupport
	m["temporary_action_limit"] = timeoutSettings.TemporaryActionLimit
	m["temporary_action_total_timeout"] = timeoutSettings.TemporaryActionTotalTimeout

	return []interface{}{m}

}

func flattenAdvancedJavaScriptTagSettings(tagSettings *dynatraceConfigV1.AdvancedJavaScriptTagSettings) []interface{} {
	if tagSettings == nil {
		return []interface{}{tagSettings}
	}

	m := make(map[string]interface{})

	m["sync_beacon_firefox"] = tagSettings.SyncBeaconFirefox
	m["sync_beacon_internet_explorer"] = tagSettings.SyncBeaconInternetExplorer
	m["instrument_unsupported_ajax_frameworks"] = tagSettings.InstrumentUnsupportedAjaxFrameworks
	m["special_characters_to_escape"] = tagSettings.SpecialCharactersToEscape
	m["max_action_name_length"] = tagSettings.MaxActionNameLength
	m["max_errors_to_capture"] = tagSettings.MaxErrorsToCapture
	m["additional_event_handlers"] = flattenAdditionalEventHandlers(&tagSettings.AdditionalEventHandlers)
	m["event_wrapper_settings"] = flattenWrapperSettings(&tagSettings.EventWrapperSettings)
	m["global_event_capture_settings"] = flattenGlobalEventCaptureSettings(&tagSettings.GlobalEventCaptureSettings)

	return []interface{}{m}
}

func flattenAdditionalEventHandlers(eventHandlers *dynatraceConfigV1.AdditionalEventHandlers) []interface{} {
	if eventHandlers == nil {
		return []interface{}{eventHandlers}
	}

	m := make(map[string]interface{})

	m["user_mouseup_event_for_clicks"] = eventHandlers.UserMouseupEventForClicks
	m["click_event_handler"] = eventHandlers.ClickEventHandler
	m["mouseup_event_handler"] = eventHandlers.MouseupEventHandler
	m["blur_event_handler"] = eventHandlers.BlurEventHandler
	m["change_event_handler"] = eventHandlers.ChangeEventHandler
	m["to_string_method"] = eventHandlers.ToStringMethod
	m["max_dom_nodes_to_instrument"] = eventHandlers.MaxDomNodesToInstrument

	return []interface{}{m}

}

func flattenWrapperSettings(settings *dynatraceConfigV1.EventWrapperSettings) []interface{} {
	if settings == nil {
		return []interface{}{settings}
	}

	m := make(map[string]interface{})

	m["click"] = settings.Click
	m["mouse_up"] = settings.MouseUp
	m["change"] = settings.Change
	m["blur"] = settings.Blur
	m["touch_start"] = settings.TouchStart
	m["touch_end"] = settings.TouchEnd

	return []interface{}{m}

}

func flattenGlobalEventCaptureSettings(settings *dynatraceConfigV1.GlobalEventCaptureSettings) []interface{} {
	if settings == nil {
		return []interface{}{settings}
	}

	m := make(map[string]interface{})

	m["mouse_up"] = settings.MouseUp
	m["mouse_down"] = settings.MouseDown
	m["click"] = settings.Click
	m["double_click"] = settings.DoubleClick
	m["key_up"] = settings.KeyUp
	m["key_down"] = settings.KeyDown
	m["scroll"] = settings.Scroll
	m["additional_event_captured_as_user_input"] = settings.AdditionalEventCapturedAsUserInput

	return []interface{}{m}
}

func flattenBrowserRestrictionSettings(restrictionSettings *dynatraceConfigV1.WebApplicationConfigBrowserRestrictionSettings) []interface{} {
	if restrictionSettings == nil {
		return []interface{}{restrictionSettings}
	}

	m := make(map[string]interface{})

	m["mode"] = restrictionSettings.Mode
	m["browser_restriction"] = flattenBrowserConfigRestriction(restrictionSettings.BrowserRestrictions)

	return []interface{}{m}
}

func flattenBrowserConfigRestriction(restrictions *[]dynatraceConfigV1.WebApplicationConfigBrowserRestriction) []interface{} {
	if restrictions != nil {
		brs := make([]interface{}, len(*restrictions))

		for i, restrictions := range *restrictions {
			br := make(map[string]interface{})

			br["browser_version"] = restrictions.BrowserVersion
			br["browser_type"] = restrictions.BrowserType
			br["platform"] = restrictions.Platform
			br["comparator"] = restrictions.Comparator

			brs[i] = br
		}
		return brs
	}
	return make([]interface{}, 0)

}

func flattenIPAddressRestrictionSettings(restrictionSettings *dynatraceConfigV1.WebApplicationConfigIpAddressRestrictionSettings) []interface{} {
	if restrictionSettings == nil {
		return []interface{}{restrictionSettings}
	}

	m := make(map[string]interface{})

	m["mode"] = restrictionSettings.Mode
	m["ip_address_restriction"] = flattenIPAddressRestrictions(restrictionSettings.IpAddressRestrictions)

	return []interface{}{m}
}

func flattenIPAddressRestrictions(restrictions *[]dynatraceConfigV1.IpAddressRange) []interface{} {
	if restrictions != nil {
		irs := make([]interface{}, len(*restrictions))

		for i, restrictions := range *restrictions {
			ir := make(map[string]interface{})

			ir["subnet_mask"] = restrictions.SubnetMask
			ir["address"] = restrictions.Address
			ir["address_to"] = restrictions.AddressTo

			irs[i] = ir
		}
		return irs
	}

	return make([]interface{}, 0)

}

func flattenJavaScriptInjectionRules(injectionRules *[]dynatraceConfigV1.JavaScriptInjectionRules) []interface{} {
	if injectionRules != nil {
		jrs := make([]interface{}, len(*injectionRules))

		for i, injectionRules := range *injectionRules {
			jr := make(map[string]interface{})

			jr["enabled"] = injectionRules.Enabled
			jr["url_operator"] = injectionRules.UrlOperator
			jr["url_pattern"] = injectionRules.UrlPattern
			jr["rule"] = injectionRules.Rule
			jr["html_pattern"] = injectionRules.HtmlPattern

			jrs[i] = jr
		}
		return jrs
	}

	return make([]interface{}, 0)
}

func flattenUserTags(userTags *[]dynatraceConfigV1.UserTag) []interface{} {
	if userTags != nil {
		uts := make([]interface{}, len(*userTags))

		for i, userTags := range *userTags {
			ut := make(map[string]interface{})

			ut["unique_id"] = userTags.UniqueId
			ut["metadata_id"] = userTags.MetadataId
			ut["cleanup_rule"] = userTags.CleanupRule
			ut["server_side_request_attribute"] = userTags.ServerSideRequestAttribute
			ut["ignore_case"] = userTags.IgnoreCase
			uts[i] = ut

		}
		return uts
	}

	return make([]interface{}, 0)
}

func flattenUserActionAndSessionProperties(sessionProperties *[]dynatraceConfigV1.UserActionAndSessionProperties) []interface{} {
	if sessionProperties != nil {
		sps := make([]interface{}, len(*sessionProperties))

		for i, sessionProperties := range *sessionProperties {
			sp := make(map[string]interface{})

			sp["display_name"] = sessionProperties.DisplayName
			sp["type"] = sessionProperties.Type
			sp["origin"] = sessionProperties.Origin
			sp["aggregation"] = sessionProperties.Aggregation
			sp["store_as_user_action_property"] = sessionProperties.StoreAsUserActionProperty
			sp["store_as_session_property"] = sessionProperties.StoreAsSessionProperty
			sp["cleanup_rule"] = sessionProperties.CleanupRule
			sp["server_side_request_attribute"] = sessionProperties.ServerSideRequestAttribute
			sp["unique_id"] = sessionProperties.UniqueId
			sp["key"] = sessionProperties.Key
			sp["metadata_id"] = sessionProperties.MetadataId
			sp["ignore_case"] = sessionProperties.IgnoreCase

			sps[i] = sp
		}
		return sps
	}

	return make([]interface{}, 0)
}

func flattenUserActionNamingSettings(namingSettings *dynatraceConfigV1.UserActionNamingSettings) []interface{} {
	if namingSettings == nil {
		return []interface{}{namingSettings}
	}

	n := make(map[string]interface{})

	n["ignore_case"] = namingSettings.IgnoreCase
	n["use_first_detected_load_action"] = namingSettings.UseFirstDetectedLoadAction
	n["split_user_actions_by_domain"] = namingSettings.SplitUserActionsByDomain
	n["placeholder"] = flattenPlaceholders(namingSettings.Placeholders)
	n["load_action_naming_rule"] = flattenActionNamingRule(namingSettings.LoadActionNamingRules)
	n["xhr_action_naming_rule"] = flattenActionNamingRule(namingSettings.XhrActionNamingRules)
	n["custom_action_naming_rule"] = flattenActionNamingRule(namingSettings.CustomActionNamingRules)

	return []interface{}{n}

}

func flattenPlaceholders(placeholders *[]dynatraceConfigV1.UserActionNamingPlaceholder) []interface{} {
	if placeholders != nil {
		phs := make([]interface{}, len(*placeholders))

		for i, placeholders := range *placeholders {
			ph := make(map[string]interface{})

			ph["name"] = placeholders.Name
			ph["input"] = placeholders.Input
			ph["processing_part"] = placeholders.ProcessingPart
			ph["metadata_id"] = placeholders.MetadataId
			ph["use_guessed_element_identifier"] = placeholders.UseGuessedElementIdentifier
			ph["processing_steps"] = flattenProcessingSteps(placeholders.ProcessingSteps)

			phs[i] = ph
		}
		return phs
	}

	return make([]interface{}, 0)
}

func flattenProcessingSteps(steps *[]dynatraceConfigV1.UserActionNamingPlaceholderProcessingStep) []interface{} {
	if steps != nil {
		pps := make([]interface{}, len(*steps))

		for i, steps := range *steps {
			pp := make(map[string]interface{})

			pp["type"] = steps.Type
			pp["pattern_before"] = steps.PatternBefore
			pp["pattern_before_search_type"] = steps.PatternBeforeSearchType
			pp["pattern_after"] = steps.PatternAfter
			pp["pattern_after_search_type"] = steps.PatternAfterSearchType
			pp["replacement"] = steps.Replacement
			pp["pattern_to_replace"] = steps.PatternToReplace
			pp["regular_expression"] = steps.RegularExpression
			pp["fallback_to_input"] = steps.FallbackToInput

			pps[i] = pp
		}
		return pps
	}

	return make([]interface{}, 0)

}

func flattenActionNamingRule(namingRules *[]dynatraceConfigV1.UserActionNamingRule) []interface{} {
	if namingRules != nil {
		nrs := make([]interface{}, len(*namingRules))

		for i, namingRules := range *namingRules {
			nr := make(map[string]interface{})

			nr["template"] = namingRules.Template
			nr["condition"] = flattenNamingRuleConditions(namingRules.Conditions)

			nrs[i] = nr
		}
		return nrs
	}

	return make([]interface{}, 0)
}

func flattenNamingRuleConditions(conditions *[]dynatraceConfigV1.UserActionNamingRuleCondition) []interface{} {
	if conditions != nil {
		rcs := make([]interface{}, len(*conditions))

		for i, conditions := range *conditions {
			rc := make(map[string]interface{})

			rc["operand1"] = conditions.Operand1
			rc["operand2"] = conditions.Operand2
			rc["operator"] = conditions.Operator

			rcs[i] = rc
		}
		return rcs
	}

	return make([]interface{}, 0)
}

func flattenMetaDataCaptureSettings(captureSettings *[]dynatraceConfigV1.MetaDataCapturing) []interface{} {
	if captureSettings != nil {
		mcs := make([]interface{}, len(*captureSettings))

		for i, captureSettings := range *captureSettings {
			mc := make(map[string]interface{})

			mc["type"] = captureSettings.Type
			mc["capturing_name"] = captureSettings.CapturingName
			mc["name"] = captureSettings.Name
			mc["unique_id"] = captureSettings.Name
			mc["public_metadata"] = captureSettings.PublicMetadata

			mcs[i] = mc
		}
		return mcs
	}

	return make([]interface{}, 0)

}

func flattenConversionGoals(goals *[]dynatraceConfigV1.ConversionGoal) []interface{} {
	if goals != nil {
		cgs := make([]interface{}, len(*goals))

		for i, goals := range *goals {
			cg := make(map[string]interface{})

			cg["name"] = goals.Name
			cg["id"] = goals.Id
			cg["type"] = goals.Type
			cg["destination_details"] = flattenDestinationDetails(goals.DestinationDetails)
			cg["user_action_details"] = flattenUserActionDetails(goals.UserActionDetails)
			cg["visit_duration_details"] = flattenVisitDurationDetails(goals.VisitDurationDetails)
			cg["visit_num_action_details"] = flattenVisitNumActionDetails(goals.VisitNumActionDetails)

			cgs[i] = cg
		}
		return cgs
	}

	return make([]interface{}, 0)

}

func flattenDestinationDetails(details *dynatraceConfigV1.DestinationDetails) []interface{} {
	if details == nil {
		return []interface{}{details}
	}

	m := make(map[string]interface{})

	m["url_or_path"] = details.UrlOrPath
	m["match_type"] = details.MatchType
	m["case_sensitive"] = details.CaseSensitive

	return []interface{}{m}
}

func flattenUserActionDetails(details *dynatraceConfigV1.UserActionDetails) []interface{} {
	if details == nil {
		return []interface{}{details}
	}

	m := make(map[string]interface{})

	m["value"] = details.Value
	m["case_sensitive"] = details.CaseSensitive
	m["match_type"] = details.MatchType
	m["match_entity"] = details.MatchEntity
	m["action_type"] = details.ActionType

	return []interface{}{m}
}

func flattenVisitDurationDetails(details *dynatraceConfigV1.VisitDurationDetails) []interface{} {
	if details == nil {
		return []interface{}{details}
	}

	m := make(map[string]interface{})

	m["duration_in_millis"] = details.DurationInMillis

	return []interface{}{m}
}

func flattenVisitNumActionDetails(details *dynatraceConfigV1.VisitNumActionDetails) []interface{} {
	if details == nil {
		return []interface{}{details}
	}

	m := make(map[string]interface{})

	m["num_user_actions"] = details.NumUserActions

	return []interface{}{m}
}
