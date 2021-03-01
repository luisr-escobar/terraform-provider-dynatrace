package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandMaintenanceWindow(d *schema.ResourceData) (*dynatraceConfigV1.MaintenanceWindow, error) {

	var dtMainteanceWindow dynatraceConfigV1.MaintenanceWindow

	if name, ok := d.GetOk("name"); ok {
		dtMainteanceWindow.SetName(name.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		dtMainteanceWindow.SetDescription(description.(string))
	}

	if mwType, ok := d.GetOk("type"); ok {
		dtMainteanceWindow.SetType(mwType.(string))
	}

	if suppression, ok := d.GetOk("suppression"); ok {
		dtMainteanceWindow.SetSuppression(suppression.(string))
	}

	if scope, ok := d.GetOk("scope"); ok {
		dtMainteanceWindow.SetScope(expandMaintenanceWindowScope(scope.([]interface{})))
	}

	if schedule, ok := d.GetOk("schedule"); ok {
		dtMainteanceWindow.SetSchedule(expandMaintenanceWindowSchedule(schedule.([]interface{})))
	}
	return &dtMainteanceWindow, nil
}

func expandMaintenanceWindowScope(scope []interface{}) dynatraceConfigV1.Scope {
	if len(scope) == 0 || scope[0] == nil {
		return dynatraceConfigV1.Scope{}
	}

	m := scope[0].(map[string]interface{})

	mws := dynatraceConfigV1.Scope{}

	if entities, ok := m["entities"]; ok {
		mws.Entities = expandEntities(entities.([]interface{}))
	}

	if matches, ok := m["match"]; ok {
		mws.Matches = expandMatches(matches.([]interface{}))
	}

	return mws

}

func expandEntities(entities []interface{}) []string {
	mwe := make([]string, len(entities))

	for i, v := range entities {
		mwe[i] = v.(string)
	}

	return mwe

}

func expandMatches(matches []interface{}) []dynatraceConfigV1.MonitoredEntityFilter {
	mwm := make([]dynatraceConfigV1.MonitoredEntityFilter, len(matches))

	for i, match := range matches {

		m := match.(map[string]interface{})

		var dtMonitoredEntityFilter dynatraceConfigV1.MonitoredEntityFilter

		if mwType, ok := m["type"].(string); ok && len(mwType) != 0 {
			dtMonitoredEntityFilter.SetType(mwType)
		}

		if mzID, ok := m["mz_id"].(string); ok && len(mzID) != 0 {
			dtMonitoredEntityFilter.SetMzId(mzID)
		}

		if tags, ok := m["tags"].([]interface{}); ok {
			dtMonitoredEntityFilter.SetTags(expandTags(tags))
		}

		if tagCombination, ok := m["tag_combination"].(string); ok && len(tagCombination) != 0 {
			dtMonitoredEntityFilter.SetTagCombination(tagCombination)
		}
		mwm[i] = dtMonitoredEntityFilter
	}

	return mwm
}

func expandTags(tags []interface{}) []dynatraceConfigV1.TagInfo {
	mwt := make([]dynatraceConfigV1.TagInfo, len(tags))

	for i, tag := range tags {
		m := tag.(map[string]interface{})

		var dtTagInfo dynatraceConfigV1.TagInfo

		if context, ok := m["context"].(string); ok {
			dtTagInfo.SetContext(context)
		}

		if key, ok := m["key"].(string); ok {
			dtTagInfo.SetKey(key)
		}

		if value, ok := m["value"].(string); ok && len(value) != 0 {
			dtTagInfo.SetValue(value)
		}

		mwt[i] = dtTagInfo
	}

	return mwt
}

func expandMaintenanceWindowSchedule(schedule []interface{}) dynatraceConfigV1.Schedule {
	if len(schedule) == 0 || schedule[0] == nil {
		return dynatraceConfigV1.Schedule{}
	}

	dtMainteanceWindowSchedule := dynatraceConfigV1.NewScheduleWithDefaults()

	m := schedule[0].(map[string]interface{})

	if recurrenceType, ok := m["recurrence_type"].(string); ok {
		dtMainteanceWindowSchedule.SetRecurrenceType(recurrenceType)
	}

	if recurrence, ok := m["recurrence"].([]interface{}); ok && len(recurrence) != 0 {
		dtMainteanceWindowSchedule.SetRecurrence(expandRecurrence(recurrence))
	}

	if start, ok := m["start"].(string); ok {
		dtMainteanceWindowSchedule.SetStart(start)
	}

	if end, ok := m["end"].(string); ok {
		dtMainteanceWindowSchedule.SetEnd(end)
	}

	if zoneID, ok := m["zone_id"].(string); ok {
		dtMainteanceWindowSchedule.SetZoneId(zoneID)
	}

	return *dtMainteanceWindowSchedule

}

func expandRecurrence(recurrence []interface{}) dynatraceConfigV1.Recurrence {
	if len(recurrence) == 0 || recurrence[0] == nil {
		return dynatraceConfigV1.Recurrence{}
	}

	m := recurrence[0].(map[string]interface{})

	var dtRecurrence dynatraceConfigV1.Recurrence

	if dayOfWeek, ok := m["day_of_week"].(string); ok {
		dtRecurrence.SetDayOfWeek(dayOfWeek)
	}

	if dayOfMonth, ok := m["day_of_month"].(int); ok && dayOfMonth > 0 {
		dtRecurrence.SetDayOfMonth(int32(dayOfMonth))
	}

	if startTime, ok := m["start_time"].(string); ok {
		dtRecurrence.SetStartTime(startTime)
	}

	if durationMinutes, ok := m["duration_minutes"].(int); ok {
		dtRecurrence.SetDurationMinutes(int32(durationMinutes))
	}

	return dtRecurrence

}

func flattenMaintenanceWindowScopeData(maintenanceWindowScope *dynatraceConfigV1.Scope) []interface{} {
	if maintenanceWindowScope == nil {
		return []interface{}{maintenanceWindowScope}
	}

	m := make(map[string]interface{})

	m["entities"] = flattenMaintenanceWindowEntities(maintenanceWindowScope.Entities)
	m["match"] = flattenMaintenanceWindowMatches(&maintenanceWindowScope.Matches)

	return []interface{}{m}
}

func flattenMaintenanceWindowEntities(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}

func flattenMaintenanceWindowMatches(maintenanceWindowMatches *[]dynatraceConfigV1.MonitoredEntityFilter) []interface{} {
	if maintenanceWindowMatches != nil {
		mwm := make([]interface{}, len(*maintenanceWindowMatches), len(*maintenanceWindowMatches))

		for i, maintenanceWindowMatches := range *maintenanceWindowMatches {
			mm := make(map[string]interface{})

			mm["type"] = maintenanceWindowMatches.Type
			mm["mz_id"] = maintenanceWindowMatches.MzId
			mm["tags"] = flattenMaintenanceWindowTags(&maintenanceWindowMatches.Tags)
			mm["tag_combination"] = maintenanceWindowMatches.TagCombination
			mwm[i] = mm
		}

		return mwm
	}

	return make([]interface{}, 0)

}

func flattenMaintenanceWindowTags(maintenanceWindowTags *[]dynatraceConfigV1.TagInfo) []interface{} {
	if maintenanceWindowTags != nil {
		mwt := make([]interface{}, len(*maintenanceWindowTags), len(*maintenanceWindowTags))

		for i, maintenanceWindowTags := range *maintenanceWindowTags {
			mt := make(map[string]interface{})

			mt["context"] = maintenanceWindowTags.Context
			mt["key"] = maintenanceWindowTags.Key
			mt["value"] = maintenanceWindowTags.Value
			mwt[i] = mt
		}

		return mwt
	}

	return make([]interface{}, 0)

}

func flattenMaintenanceWindowScheduleData(maintenanceWindowSchedule *dynatraceConfigV1.Schedule) []interface{} {
	if maintenanceWindowSchedule == nil {
		return []interface{}{maintenanceWindowSchedule}
	}

	m := make(map[string]interface{})

	m["recurrence_type"] = maintenanceWindowSchedule.RecurrenceType
	m["recurrence"] = flattenMaintenanceWindowRecurrence(maintenanceWindowSchedule.Recurrence)
	m["start"] = maintenanceWindowSchedule.Start
	m["end"] = maintenanceWindowSchedule.End
	m["zone_id"] = maintenanceWindowSchedule.ZoneId

	return []interface{}{m}
}

func flattenMaintenanceWindowRecurrence(maintenanceWindowRecurrence *dynatraceConfigV1.Recurrence) []interface{} {
	if maintenanceWindowRecurrence == nil {
		return nil
	}

	m := make(map[string]interface{})

	m["day_of_week"] = maintenanceWindowRecurrence.DayOfWeek
	m["day_of_month"] = maintenanceWindowRecurrence.DayOfMonth
	m["start_time"] = maintenanceWindowRecurrence.StartTime
	m["duration_minutes"] = maintenanceWindowRecurrence.DurationMinutes

	return []interface{}{m}
}
