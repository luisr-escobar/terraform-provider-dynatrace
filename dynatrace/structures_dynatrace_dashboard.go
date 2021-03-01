package dynatrace

import (
	"encoding/json"
	"log"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandDashboard(d *schema.ResourceData) (*dynatraceConfigV1.Dashboard, error) {

	var dtDashboard dynatraceConfigV1.Dashboard

	if dashboardMetadata, ok := d.GetOk("dashboard_metadata"); ok {
		dtDashboard.SetDashboardMetadata(expandDashboardMetadata(dashboardMetadata.([]interface{})))
	}

	if tiles, ok := d.GetOk("tile"); ok {
		dtDashboard.SetTiles(expandDashboardTiles(tiles.([]interface{})))
	}

	return &dtDashboard, nil

}

func expandExistingDashboard(d *schema.ResourceData, id string) (*dynatraceConfigV1.Dashboard, error) {

	var dtDashboard dynatraceConfigV1.Dashboard

	dtDashboard.SetId(id)

	if dashboardMetadata, ok := d.GetOk("dashboard_metadata"); ok {
		dtDashboard.SetDashboardMetadata(expandDashboardMetadata(dashboardMetadata.([]interface{})))
	}

	if tiles, ok := d.GetOk("tile"); ok {
		dtDashboard.SetTiles(expandDashboardTiles(tiles.([]interface{})))
	}

	return &dtDashboard, nil
}

func expandDashboardMetadata(dashboardMetadata []interface{}) dynatraceConfigV1.DashboardMetadata {
	if len(dashboardMetadata) == 0 || dashboardMetadata[0] == nil {
		return dynatraceConfigV1.DashboardMetadata{}
	}

	m := dashboardMetadata[0].(map[string]interface{})

	dtDashboardMetadata := dynatraceConfigV1.NewDashboardMetadataWithDefaults()

	if name, ok := m["name"].(string); ok {
		dtDashboardMetadata.SetName(name)
	}

	if shared, ok := m["shared"].(bool); ok {
		dtDashboardMetadata.SetShared(shared)
	}

	if owner, ok := m["owner"].(string); ok {
		dtDashboardMetadata.SetOwner(owner)
	}

	if sharingDetails, ok := m["sharing_details"].([]interface{}); ok && len(sharingDetails) != 0 {
		dtDashboardMetadata.SetSharingDetails(expandDashboardSharingDetails(sharingDetails))
	}

	if dashboardFilter, ok := m["dashboard_filter"].([]interface{}); ok && len(dashboardFilter) != 0 {
		dtDashboardMetadata.SetDashboardFilter(expandDashboardFilter(dashboardFilter))
	}

	if tags, ok := m["tags"]; ok {
		dtDashboardMetadata.SetTags(expandDashboardTags(tags.([]interface{})))
	}

	if preset, ok := m["preset"].(bool); ok {
		dtDashboardMetadata.SetPreset(preset)
	}

	if validFilterKeys, ok := m["valid_filter_keys"]; ok {
		dtDashboardMetadata.SetValidFilterKeys(expandDashboardFilterKeys(validFilterKeys.([]interface{})))
	}

	return *dtDashboardMetadata

}

func expandDashboardSharingDetails(sharingDetails []interface{}) dynatraceConfigV1.SharingInfo {
	if len(sharingDetails) == 0 || sharingDetails[0] == nil {
		return dynatraceConfigV1.SharingInfo{}
	}

	m := sharingDetails[0].(map[string]interface{})

	dtSharingDetails := dynatraceConfigV1.NewSharingInfoWithDefaults()

	if linkShared, ok := m["link_shared"].(bool); ok {
		dtSharingDetails.SetLinkShared(linkShared)
	}

	if published, ok := m["published"].(bool); ok {
		dtSharingDetails.SetPublished(published)
	}

	return *dtSharingDetails
}

func expandDashboardFilter(filters []interface{}) dynatraceConfigV1.DashboardFilter {
	if len(filters) == 0 || filters[0] == nil {
		return dynatraceConfigV1.DashboardFilter{}
	}

	m := filters[0].(map[string]interface{})

	dtDashboardFilter := dynatraceConfigV1.NewDashboardFilterWithDefaults()

	if timeframe, ok := m["timeframe"].(string); ok && len(timeframe) != 0 {
		dtDashboardFilter.SetTimeframe(timeframe)
	}

	if managementZone, ok := m["management_zone"].([]interface{}); ok && len(managementZone) != 0 {
		dtDashboardFilter.SetManagementZone(expandDashboardManagementZone(managementZone))
	}

	return *dtDashboardFilter

}

func expandDashboardTags(tags []interface{}) []string {
	ddt := make([]string, len(tags))

	for i, v := range tags {
		ddt[i] = v.(string)
	}

	return ddt

}

func expandDashboardFilterKeys(filterKeys []interface{}) []string {
	dfk := make([]string, len(filterKeys))

	for i, v := range filterKeys {
		dfk[i] = v.(string)
	}

	return dfk

}

func expandDashboardManagementZone(managementZone []interface{}) dynatraceConfigV1.EntityShortRepresentation {
	if len(managementZone) == 0 || managementZone[0] == nil {
		return dynatraceConfigV1.EntityShortRepresentation{}
	}

	m := managementZone[0].(map[string]interface{})

	dtDashboardManagementZone := dynatraceConfigV1.NewEntityShortRepresentationWithDefaults()

	if id, ok := m["id"].(string); ok {
		dtDashboardManagementZone.SetId(id)
	}

	if name, ok := m["name"].(string); ok && len(name) != 0 {
		dtDashboardManagementZone.SetName(name)
	}

	return *dtDashboardManagementZone

}

func expandDashboardTiles(dashboardTiles []interface{}) []dynatraceConfigV1.Tile {
	if len(dashboardTiles) < 1 {
		return []dynatraceConfigV1.Tile{}
	}

	dts := make([]dynatraceConfigV1.Tile, len(dashboardTiles))

	for i, tile := range dashboardTiles {

		m := tile.(map[string]interface{})

		var dtDashboardTile dynatraceConfigV1.Tile

		if name, ok := m["name"].(string); ok {
			dtDashboardTile.SetName(name)
		}

		if tileType, ok := m["tile_type"].(string); ok {
			dtDashboardTile.SetTileType(tileType)
		}

		if configured, ok := m["configured"].(bool); ok {
			dtDashboardTile.SetConfigured(configured)
		}

		if bounds, ok := m["bounds"].([]interface{}); ok && len(bounds) != 0 {
			dtDashboardTile.SetBounds(expandTileBounds(bounds))
		}

		if tileFilter, ok := m["tile_filter"].([]interface{}); ok && len(tileFilter) != 0 {
			dtDashboardTile.SetTileFilter(expandTileFilter(tileFilter))
		}

		if assignedEntities, ok := m["assigned_entities"]; ok {
			dtDashboardTile.SetAssignedEntities(expandAssignedEntities(assignedEntities.([]interface{})))
		}

		if metric, ok := m["metric"].(string); ok && len(metric) != 0 {
			dtDashboardTile.SetMetric(metric)
		}

		if filterConfig, ok := m["filter_config"].([]interface{}); ok && len(filterConfig) != 0 {
			dtDashboardTile.SetFilterConfig(expandFilterConfig(filterConfig))
		}

		if chartVisible, ok := m["chart_visible"].(bool); ok {
			dtDashboardTile.SetChartVisible(chartVisible)
		}

		if markdown, ok := m["markdown"].(string); ok {
			dtDashboardTile.SetMarkdown(markdown)
		}

		if excludeMaintenanceWindows, ok := m["exclude_maintenance_windows"].(bool); ok {
			dtDashboardTile.SetExcludeMaintenanceWindows(excludeMaintenanceWindows)
		}

		if customName, ok := m["custom_name"].(string); ok && len(customName) != 0 {
			dtDashboardTile.SetCustomName(customName)
		}

		if query, ok := m["query"].(string); ok && len(query) != 0 {
			dtDashboardTile.SetQuery(query)
		}

		if ddType, ok := m["type"].(string); ok && len(ddType) != 0 {
			dtDashboardTile.SetType(ddType)
		}

		if timeFrameShift, ok := m["timeframe_shift"].(string); ok && len(timeFrameShift) != 0 {
			dtDashboardTile.SetTimeFrameShift(timeFrameShift)
		}

		if visualizationConfig, ok := m["visualization_config"].([]interface{}); ok && len(visualizationConfig) != 0 {
			dtDashboardTile.SetVisualizationConfig(expandVisualizationConfig(visualizationConfig))
		}

		if limit, ok := m["limit"].(int); ok && limit != 0 {
			dtDashboardTile.SetLimit(int32(limit))
		}

		dts[i] = dtDashboardTile
	}

	return dts
}

func expandTileBounds(bounds []interface{}) dynatraceConfigV1.TileBounds {
	if len(bounds) == 0 || bounds[0] == nil {
		return dynatraceConfigV1.TileBounds{}
	}

	m := bounds[0].(map[string]interface{})

	dtb := dynatraceConfigV1.TileBounds{}

	if top, ok := m["top"].(int); ok {
		dtb.SetTop(int32(top))
	}

	if left, ok := m["left"].(int); ok {
		dtb.SetLeft(int32(left))
	}

	if width, ok := m["width"].(int); ok {
		dtb.SetWidth(int32(width))
	}

	if height, ok := m["height"].(int); ok {
		dtb.SetHeight(int32(height))
	}

	return dtb

}

func expandTileFilter(filter []interface{}) dynatraceConfigV1.TileFilter {
	if len(filter) == 0 || filter[0] == nil {
		return dynatraceConfigV1.TileFilter{}
	}

	m := filter[0].(map[string]interface{})

	dtTileFilter := dynatraceConfigV1.NewTileFilterWithDefaults()

	if timeframe, ok := m["timeframe"].(string); ok {
		dtTileFilter.SetTimeframe(timeframe)
	}

	if managementZone, ok := m["management_zone"].([]interface{}); ok && len(managementZone) != 0 {
		dtTileFilter.SetManagementZone(expandDashboardManagementZone(managementZone))
	}

	return *dtTileFilter

}

func expandAssignedEntities(assignedEntities []interface{}) []string {
	dae := make([]string, len(assignedEntities))

	for i, v := range assignedEntities {
		dae[i] = v.(string)
	}

	return dae

}

func expandVisualizationConfig(config []interface{}) dynatraceConfigV1.UserSessionQueryTileConfiguration {
	if len(config) == 0 || config[0] == nil {
		return dynatraceConfigV1.UserSessionQueryTileConfiguration{}
	}

	m := config[0].(map[string]interface{})

	dtUserSessionQueryTileConfiguration := dynatraceConfigV1.NewUserSessionQueryTileConfigurationWithDefaults()

	if hasAxisBucketing, ok := m["has_axis_bucketing"].(bool); ok && hasAxisBucketing {
		dtUserSessionQueryTileConfiguration.SetHasAxisBucketing(hasAxisBucketing)
	}

	return *dtUserSessionQueryTileConfiguration

}

func expandFilterConfig(filterConfig []interface{}) dynatraceConfigV1.CustomFilterConfig {
	m := filterConfig[0].(map[string]interface{})

	dfc := dynatraceConfigV1.CustomFilterConfig{}

	if dfType, ok := m["type"]; ok {
		dfc.Type = dfType.(string)
	}

	if customName, ok := m["custom_name"]; ok {
		dfc.CustomName = customName.(string)
	}

	if defaultName, ok := m["default_name"]; ok {
		dfc.DefaultName = defaultName.(string)
	}

	if chartConfig, ok := m["chart_config"]; ok {
		dfc.ChartConfig = expandDashboardChartConfig(chartConfig.([]interface{}))
	}

	if filtersPerEntityType, ok := m["filters_per_entity_type"]; ok {
		dfc.FiltersPerEntityType = expandFiltersPerEntityType(filtersPerEntityType.(string))
	}

	return dfc

}

func expandDashboardChartConfig(chartConfig []interface{}) dynatraceConfigV1.CustomFilterChartConfig {
	if len(chartConfig) == 0 || chartConfig[0] == nil {
		return dynatraceConfigV1.CustomFilterChartConfig{}
	}

	m := chartConfig[0].(map[string]interface{})

	dtDashboardChartConfig := dynatraceConfigV1.NewCustomFilterChartConfigWithDefaults()

	if legendShown, ok := m["legend_shown"].(bool); ok {
		dtDashboardChartConfig.SetLegendShown(legendShown)
	}

	if dcType, ok := m["type"].(string); ok {
		dtDashboardChartConfig.SetType(dcType)
	}

	if series, ok := m["series"].([]interface{}); ok {
		dtDashboardChartConfig.SetSeries(expandCustomChartSeries(series))
	}

	if resultMetadata, ok := m["result_metadata"].(string); ok {
		dtDashboardChartConfig.SetResultMetadata(expandResultMetadata(resultMetadata))
	}

	if axisLimits, ok := m["axis_limits"].(string); ok && len(axisLimits) != 0 {
		dtDashboardChartConfig.SetAxisLimits(expandAxisLimits(axisLimits))
	}

	if leftAxisCustomUnit, ok := m["left_axis_custom_unit"].(string); ok && len(leftAxisCustomUnit) != 0 {
		dtDashboardChartConfig.SetLeftAxisCustomUnit(leftAxisCustomUnit)
	}

	if rightAxisCustomUnit, ok := m["right_axis_custom_unit"].(string); ok && len(rightAxisCustomUnit) != 0 {
		dtDashboardChartConfig.SetRightAxisCustomUnit(rightAxisCustomUnit)
	}

	return *dtDashboardChartConfig

}

func expandCustomChartSeries(chartSeries []interface{}) []dynatraceConfigV1.CustomFilterChartSeriesConfig {
	if len(chartSeries) < 1 {
		return []dynatraceConfigV1.CustomFilterChartSeriesConfig{}
	}

	ccs := make([]dynatraceConfigV1.CustomFilterChartSeriesConfig, len(chartSeries))

	for i, series := range chartSeries {

		m := series.(map[string]interface{})

		var dtCustomFilterChartSeriesConfig dynatraceConfigV1.CustomFilterChartSeriesConfig

		if metric, ok := m["metric"].(string); ok {
			dtCustomFilterChartSeriesConfig.SetMetric(metric)
		}

		if aggregation, ok := m["aggregation"].(string); ok {
			dtCustomFilterChartSeriesConfig.SetAggregation(aggregation)
		}

		if percentile, ok := m["percentile"].(int); ok && percentile != 0 {
			dtCustomFilterChartSeriesConfig.SetPercentile(int64(percentile))
		}

		if ccType, ok := m["type"].(string); ok {
			dtCustomFilterChartSeriesConfig.SetType(ccType)
		}

		if entityType, ok := m["entity_type"].(string); ok {
			dtCustomFilterChartSeriesConfig.SetEntityType(entityType)
		}

		if dimensions, ok := m["dimensions"].([]interface{}); ok {
			dtCustomFilterChartSeriesConfig.SetDimensions(expandSeriesDimensions(dimensions))
		}

		if sortAscending, ok := m["sort_ascending"].(bool); ok {
			dtCustomFilterChartSeriesConfig.SetSortAscending(sortAscending)
		}

		if sortColumn, ok := m["sort_column"].(bool); ok {
			dtCustomFilterChartSeriesConfig.SetSortColumn(sortColumn)
		}

		if aggregationRate, ok := m["aggregation_rate"].(string); ok && len(aggregationRate) != 0 {
			dtCustomFilterChartSeriesConfig.SetAggregationRate(aggregationRate)
		}

		ccs[i] = dtCustomFilterChartSeriesConfig
	}

	return ccs
}

func expandSeriesDimensions(dimensions []interface{}) []dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig {
	if len(dimensions) < 1 {
		return []dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig{}
	}

	csd := make([]dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig, len(dimensions))

	for i, dimension := range dimensions {

		m := dimension.(map[string]interface{})

		var dtCustomFilterChartSeriesDimensionConfig dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig

		if id, ok := m["id"].(string); ok {
			dtCustomFilterChartSeriesDimensionConfig.SetId(id)
		}

		if name, ok := m["name"].(string); ok && len(name) != 0 {
			dtCustomFilterChartSeriesDimensionConfig.SetName(name)
		}

		if values, ok := m["values"]; ok {
			dtCustomFilterChartSeriesDimensionConfig.SetValues(expandDimensionsValues(values.([]interface{})))
		}

		if entityDimension, ok := m["entity_dimension"].(bool); ok {
			dtCustomFilterChartSeriesDimensionConfig.SetEntityDimension(entityDimension)
		}
		csd[i] = dtCustomFilterChartSeriesDimensionConfig
	}

	return csd

}

func expandDimensionsValues(values []interface{}) []string {
	ddv := make([]string, len(values))

	for i, v := range values {
		ddv[i] = v.(string)
	}

	return ddv

}

func expandResultMetadata(metadata interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(metadata.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal result metadata values %s: %v", metadata.(string), err)
		return nil
	}

	return val
}

func expandAxisLimits(limits interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(limits.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal axis limits values %s: %v", limits.(string), err)
		return nil
	}
	return val
}

func expandFiltersPerEntityType(filters interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(filters.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal filters per entity type values %s: %v", filters.(string), err)
		return nil
	}

	return val
}

func flattenDashboardMetadata(dashboardMetadata *dynatraceConfigV1.DashboardMetadata) []interface{} {
	if dashboardMetadata == nil {
		return []interface{}{dashboardMetadata}
	}

	m := make(map[string]interface{})

	m["name"] = dashboardMetadata.Name
	m["shared"] = dashboardMetadata.Shared
	m["owner"] = dashboardMetadata.Owner
	m["sharing_details"] = flattenSharingDetails(dashboardMetadata.SharingDetails)
	m["dashboard_filter"] = flattenDashboardFilter(dashboardMetadata.DashboardFilter)
	m["tags"] = flattenTags(dashboardMetadata.Tags)
	m["preset"] = dashboardMetadata.Preset
	m["valid_filter_keys"] = flattenValidFilterKeys(dashboardMetadata.ValidFilterKeys)

	return []interface{}{m}

}

func flattenSharingDetails(sharingDetails *dynatraceConfigV1.SharingInfo) []interface{} {
	if sharingDetails == nil {
		return []interface{}{sharingDetails}
	}

	s := make(map[string]interface{})

	s["link_shared"] = sharingDetails.LinkShared
	s["published"] = sharingDetails.Published

	return []interface{}{s}
}

func flattenDashboardFilter(dashboardFilter *dynatraceConfigV1.DashboardFilter) []interface{} {
	if dashboardFilter == nil {
		return []interface{}{dashboardFilter}
	}

	f := make(map[string]interface{})

	f["timeframe"] = dashboardFilter.Timeframe
	f["management_zone"] = flattenManagementZone(dashboardFilter.ManagementZone)

	return []interface{}{f}

}

func flattenManagementZone(dashboardManagementZone *dynatraceConfigV1.EntityShortRepresentation) []interface{} {
	if dashboardManagementZone == nil {
		return nil
	}

	m := make(map[string]interface{})

	m["id"] = dashboardManagementZone.Id
	m["name"] = dashboardManagementZone.Name

	return []interface{}{m}

}

func flattenDashboardTilesData(dashboardTiles []dynatraceConfigV1.Tile) []interface{} {
	if dashboardTiles != nil {
		dts := make([]interface{}, len(dashboardTiles), len(dashboardTiles))

		for i, dashboardTile := range dashboardTiles {
			dt := make(map[string]interface{})

			dt["name"] = dashboardTile.Name
			dt["tile_type"] = dashboardTile.TileType
			dt["configured"] = dashboardTile.Configured
			dt["bounds"] = flattenTileBounds(&dashboardTile.Bounds)
			dt["tile_filter"] = flattenTileFilter(dashboardTile.TileFilter)
			dt["assigned_entities"] = flattenAssignedEntities(dashboardTile.AssignedEntities)
			dt["filter_config"] = flattenFilterConfig(dashboardTile.FilterConfig)
			dt["chart_visible"] = dashboardTile.ChartVisible
			dt["markdown"] = dashboardTile.Markdown
			dt["exclude_maintenance_windows"] = dashboardTile.ExcludeMaintenanceWindows
			dt["custom_name"] = dashboardTile.CustomName
			dt["query"] = dashboardTile.Query
			dt["type"] = dashboardTile.Type
			dt["timeframe_shift"] = dashboardTile.TimeFrameShift
			dt["visualization_config"] = flattenVisualizationConfig(dashboardTile.VisualizationConfig)
			dt["limit"] = dashboardTile.Limit
			dts[i] = dt

		}
		return dts

	}

	return make([]interface{}, 0)
}

func flattenTileBounds(tileBounds *dynatraceConfigV1.TileBounds) []interface{} {
	if tileBounds == nil {
		return []interface{}{tileBounds}
	}

	b := make(map[string]interface{})

	b["top"] = tileBounds.Top
	b["left"] = tileBounds.Left
	b["width"] = tileBounds.Width
	b["height"] = tileBounds.Height

	return []interface{}{b}
}

func flattenTileFilter(tileFilter *dynatraceConfigV1.TileFilter) []interface{} {
	if tileFilter == nil {
		return []interface{}{tileFilter}
	}

	f := make(map[string]interface{})

	f["timeframe"] = tileFilter.Timeframe
	f["management_zone"] = flattenManagementZone(tileFilter.ManagementZone)

	return []interface{}{f}

}

func flattenFilterConfig(filterConfig *dynatraceConfigV1.CustomFilterConfig) []interface{} {
	if filterConfig == nil {
		return nil
	}

	f := make(map[string]interface{})

	f["type"] = filterConfig.Type
	f["custom_name"] = filterConfig.CustomName
	f["default_name"] = filterConfig.DefaultName
	f["chart_config"] = flattenChartConfig(&filterConfig.ChartConfig)
	f["filters_per_entity_type"] = flattenFiltersPerEntityType(&filterConfig.FiltersPerEntityType)

	return []interface{}{f}

}

func flattenChartConfig(chartConfig *dynatraceConfigV1.CustomFilterChartConfig) []interface{} {
	if chartConfig == nil {
		return []interface{}{chartConfig}
	}

	c := make(map[string]interface{})

	c["legend_shown"] = chartConfig.LegendShown
	c["type"] = chartConfig.Type
	c["series"] = flattenFilterSeries(&chartConfig.Series)
	c["axis_limits"] = flattenAxisLimits(&chartConfig.AxisLimits)
	c["result_metadata"] = flattenResultMetadata(&chartConfig.ResultMetadata)
	c["left_axis_custom_unit"] = chartConfig.LeftAxisCustomUnit
	c["right_axis_custom_unit"] = chartConfig.RightAxisCustomUnit
	c["type"] = chartConfig.Type

	return []interface{}{c}

}

func flattenFilterSeries(filterSeries *[]dynatraceConfigV1.CustomFilterChartSeriesConfig) []interface{} {
	if filterSeries != nil {
		csc := make([]interface{}, len(*filterSeries), len(*filterSeries))

		for i, filterSeries := range *filterSeries {
			cs := make(map[string]interface{})

			cs["metric"] = filterSeries.Metric
			cs["aggregation"] = filterSeries.Aggregation
			cs["percentile"] = filterSeries.Percentile
			cs["type"] = filterSeries.Type
			cs["entity_type"] = filterSeries.EntityType
			cs["dimensions"] = flattenChartDimensions(&filterSeries.Dimensions)
			cs["sort_ascending"] = filterSeries.SortAscending
			cs["sort_column"] = filterSeries.SortColumn
			cs["aggregation_rate"] = filterSeries.AggregationRate
			csc[i] = cs

		}
		return csc

	}

	return make([]interface{}, 0)
}

func flattenChartDimensions(chartDimensions *[]dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig) []interface{} {
	if chartDimensions != nil {
		csd := make([]interface{}, len(*chartDimensions), len(*chartDimensions))

		for i, chartDimension := range *chartDimensions {
			cd := make(map[string]interface{})

			cd["id"] = chartDimension.Id
			cd["name"] = chartDimension.Name
			cd["values"] = flattenDimensionValues(chartDimension.Values)
			cd["entity_dimension"] = chartDimension.EntityDimension
			csd[i] = cd

		}
		return csd

	}

	return make([]interface{}, 0)
}

func flattenDimensionValues(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}

func flattenTags(values *[]string) interface{} {
	if values == nil {
		return nil
	}

	dvs := make([]interface{}, len(*values))

	for i, e := range *values {
		dvs[i] = e
	}

	return &dvs
}

func flattenAssignedEntities(values *[]string) *[]string {
	if values == nil {
		return nil
	}

	aes := make([]string, len(*values))

	for i, e := range *values {
		aes[i] = e
	}

	return &aes
}

func flattenValidFilterKeys(values *[]string) *[]string {
	if values == nil {
		return nil
	}

	fks := make([]string, len(*values))

	for i, e := range *values {
		fks[i] = e
	}

	return &fks
}

func flattenVisualizationConfig(visualizationConfig *dynatraceConfigV1.UserSessionQueryTileConfiguration) []interface{} {
	if visualizationConfig == nil {
		return nil
	}

	v := make(map[string]interface{})

	v["has_axis_bucketing"] = visualizationConfig.HasAxisBucketing

	return []interface{}{v}
}

func flattenFiltersPerEntityType(filters interface{}) interface{} {
	json, err := json.Marshal(filters)
	if err != nil {
		log.Printf("[ERROR] Could not marshal filters per entity type value %s: %v", filters.(string), err)
		return nil
	}
	return string(json)
}

func flattenAxisLimits(limits interface{}) interface{} {
	json, err := json.Marshal(limits)
	if err != nil {
		log.Printf("[ERROR] Could not marshal axis limits values %s: %v", limits.(string), err)
		return nil
	}
	return string(json)
}

func flattenResultMetadata(metadata interface{}) interface{} {
	json, err := json.Marshal(metadata)
	if err != nil {
		log.Printf("[ERROR] Could not marshal result metadata values %s: %v", metadata.(string), err)
		return nil
	}
	return string(json)
}
