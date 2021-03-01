package dynatrace

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDynatraceDashboard() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceDashboardCreate,
		ReadContext:   resourceDynatraceDashboardRead,
		UpdateContext: resourceDynatraceDashboardUpdate,
		DeleteContext: resourceDynatraceDashboardDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dashboard_metadata": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Parameters of a dashboard.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The name of the dashboard.",
							Required:    true,
						},
						"shared": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "The dashboard is shared (true) or private (false).",
							Optional:    true,
						},
						"owner": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The owner of the dashboard.",
							Optional:    true,
							// suppress to ignore user assigned post dashboard creation
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return true
							},
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A set of tags assigned to the dashboard.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preset": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "The dashboard is a preset (true)",
							Optional:    true,
						},
						"valid_filter_keys": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A set of all possible global dashboard filters that can be applied to dashboard",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sharing_details": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Sharing configuration of a dashboard.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"link_shared": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "If true, the dashboard is shared via link and authenticated users with the link can view.",
										Optional:    true,
										Default:     false,
									},
									"published": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "If true, the dashboard is published to anyone on this environment.",
										Optional:    true,
										Default:     false,
									},
								},
							},
						},
						"dashboard_filter": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Filters, applied to a dashboard.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The default timeframe of the dashboard.",
										Optional:    true,
									},
									"management_zone": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The short representation of a Dynatrace entity.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The ID of the Dynatrace entity.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the Dynatrace entity.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"tile": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Configuration of a tile. The actual set of fields depends on the type of the tile.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The name of the tile.",
						},
						"tile_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Defines the actual set of fields depending on the value.",
							StateFunc: func(val interface{}) string {
								return strings.ToUpper(val.(string))
							},
						},
						"configured": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The tile is configured and ready to use (true) or just placed on the dashboard (false).",
						},
						"bounds": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The position and size of a tile.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"top": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.",
									},
									"left": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.",
									},
									"width": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The width of the tile, in pixels.",
									},
									"height": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The height of the tile, in pixels.",
									},
								},
							},
						},
						"tile_filter": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A filter applied to a tile, it overrides the dashboard's filter.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The default timeframe of the tile.",
									},
									"management_zone": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The short representation of a Dynatrace entity.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The ID of the Dynatrace entity.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the Dynatrace entity.",
												},
											},
										},
									},
								},
							},
						},
						"assigned_entities": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The list of Dynatrace entities, assigned to the tile.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"metric": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tile is visible and ready to use (true) or just placed on the dashboard (false).",
						},
						"filter_config": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of the custom filter of a tile.",
							Optional:    true,
							Elem:        filterConfigSchema(),
						},
						"chart_visible": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The tile is visible and ready to use (true) or just placed on the dashboard (false).",
						},
						"markdown": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The markdown-formatted content of the tile.",
						},
						"exclude_maintenance_windows": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Include (`false') or exclude (`true`) maintenance windows from availability calculations.",
						},
						"custom_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the tile, set by user.",
						},
						"query": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The visualization of the tile.",
						},
						"timeframe_shift": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift.",
						},
						"visualization_config": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A filter applied to a tile, it overrides the dashboard's filter.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"has_axis_bucketing": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The axis bucketing when enabled groups similar series in the same virtual axis.",
									},
								},
							},
						},
						"limit": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The limit of the results, if not set will use the default value of the system.",
						},
					},
				},
			},
		},
	}
}

func filterConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the filter.",
			},
			"custom_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the tile, set by user.",
			},
			"default_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default name of the tile.",
			},
			"filters_per_entity_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A list of filters, applied to specific entity types.",
				Default:      "{}",
				ValidateFunc: validation.StringIsJSON,
			},
			"chart_config": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Configuration of a custom chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"legend_shown": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Defines if a legend should be shown.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of the chart.",
						},
						"result_metadata": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Additional information about charted metric.",
							Default:      "{}",
							ValidateFunc: validation.StringIsJSON,
						},
						"axis_limits": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The optional custom y-axis limits.",
							ValidateFunc: validation.StringIsJSON,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return true
							},
						},
						"left_axis_custom_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The custom unit for the left Y-axis.",
						},
						"right_axis_custom_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The custom unit for the right Y-axis.",
						},
						"series": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of charted metrics.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the charted metric.",
									},
									"aggregation": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The charted aggregation of the metric.",
									},
									"percentile": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The charted percentile.",
									},
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The visualization of the timeseries chart.",
									},
									"entity_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of the Dynatrace entity that delivered the charted metric.",
									},
									"sort_ascending": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Sort ascending (true) or descending (false).",
									},
									"sort_column": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Sort column (true) or (false).",
									},
									"aggregation_rate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The aggregation rate.",
									},
									"dimensions": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Configuration of the charted metric splitting.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The ID of the dimension by which the metric is split.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the dimension by which the metric is split.",
												},
												"values": &schema.Schema{
													Type:        schema.TypeList,
													Description: "The splitting value.",
													Optional:    true,
													Elem: &schema.Schema{
														Type:    schema.TypeString,
														Default: "",
													},
												},
												"entity_dimension": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "The name of the entity dimension by which the metric is split.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceDynatraceDashboardCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	dd, err := expandDashboard(d)
	if err != nil {
		return diag.FromErr(err)
	}

	dashboard, _, err := dynatraceConfigClientV1.DashboardsApi.CreateDashboard(authConfigV1).Dashboard(*dd).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace dashboard",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(dashboard.Id)

	resourceDynatraceDashboardRead(ctx, d, m)

	return diags
}

func resourceDynatraceDashboardRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	dashboardID := d.Id()

	dashboard, _, err := dynatraceConfigClientV1.DashboardsApi.GetDashboard(authConfigV1, dashboardID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace dashboard",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	dashboardMetadata := flattenDashboardMetadata(&dashboard.DashboardMetadata)
	if err := d.Set("dashboard_metadata", dashboardMetadata); err != nil {
		return diag.FromErr(err)
	}

	dashboardTiles := flattenDashboardTilesData(dashboard.Tiles)
	if err := d.Set("tile", dashboardTiles); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDynatraceDashboardUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	dashboardID := d.Id()

	if d.HasChange("dashboard_metadata") || d.HasChange("tile") {

		dd, err := expandExistingDashboard(d, dashboardID)
		if err != nil {
			return diag.FromErr(err)
		}

		_, _, err = dynatraceConfigClientV1.DashboardsApi.UpdateDashboard(authConfigV1, dashboardID).Dashboard(*dd).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update dynatrace dashboard",
				Detail:   getErrorMessage(err),
			})
			return diags
		}
	}

	return resourceDynatraceDashboardRead(ctx, d, m)
}

func resourceDynatraceDashboardDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	dashboardID := d.Id()

	_, err := dynatraceConfigClientV1.DashboardsApi.DeleteDashboard(authConfigV1, dashboardID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace dashbaord",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
