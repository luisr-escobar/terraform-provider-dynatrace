package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceEnvironmentCreate,
		ReadContext:   resourceDynatraceEnvironmentRead,
		UpdateContext: resourceDynatraceEnvironmentUpdate,
		DeleteContext: resourceDynatraceEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The display name of the environment.",
				Required:    true,
			},
			"create_token": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If true, a token management token with the scopes 'apiTokens.read' and 'apiTokens.write' and 'TenantTokenManagement'is created when creating a new environment.",
				Optional:    true,
				Default:     true,
			},
			"trial": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that.",
				Optional:    true,
				Default:     false,
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Indicates whether the environment is enabled or disabled. The default value is ENABLED.",
				Optional:    true,
				Default:     "ENABLED",
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Environment Token",
				Computed:    true,
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"quotas": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Environment level consumption and quotas information. Only returned if includeConsumptionInfo or includeUncachedConsumptionInfo param is true. If skipped when editing via PUT method then already set quotas will remain.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_units": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Host units consumption and quota information on environment level. If skipped when editing via PUT method then already set quota will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"current_usage": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current environment usage.",
										Computed:    true,
									},
								},
							},
						},
						"dem_units": &schema.Schema{
							Type:        schema.TypeList,
							Description: "DEM units consumption and quota information on environment level. Not set (and not editable) if DEM units is not enabled. If skipped when editing via PUT method then already set quotas will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"monthly_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"annual_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
								},
							},
						},
						"user_sessions": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User sessions consumption and quota information on environment level. If skipped when editing via PUT method then already set quotas will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_nobile_sessions_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly Mobile user sessions environment consumption. Resets each calendar month.",
										Optional:    true,
										Computed:    true,
									},
									"consumed_user_sessions_with_web_session_replay_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly Web user sessions with replay environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"consumed_user_sessions_with_mobile_session_replay_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly Mobile user sessions with replay environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_user_sessions_with_mobile_session_replay_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly Mobile user sessions with replay environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"total_consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly total User sessions environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"total_consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly total User sessions environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"consumed_mobile_sessions_this_year	": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly Mobile user sessions environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"consumed_user_sessions_with_web_session_replay_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly Web user sessions with replay environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"total_monthly_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"total_annual_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly Web user sessions with replay environment consumption. Resets each year on license creation date anniversary.",
										Optional:    true,
									},
								},
							},
						},
						"session_properties": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User session properties consumption information on environment level.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
								},
							},
						},
						"synthetic_monitors": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Synthetic monitors consumption and quota information on environment level. Not set (and not editable) if neither Synthetic nor DEM units is enabled. If skipped when editing via PUT method then already set quotas will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"monthly_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"annual_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
								},
							},
						},
						"custom_metrics": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Custom metrics consumption and quota information on environment level. Not set (and not editable) if Custom metrics is not enabled. Not set (and not editable) if Davis data units is enabled. If skipped when editing via PUT method then already set quota will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"current_usage": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current environment usage.",
										Computed:    true,
									},
								},
							},
						},
						"davis_data_units": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Davis data units consumption and quota information on environment level. Not set (and not editable) if Davis data units is not enabled. If skipped when editing via PUT method then already set quotas will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"monthly_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"annual_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
								},
							},
						},
						"log_monitoring": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Log monitoring consumption and quota information on environment level. Not set (and not editable) if Log monitoring is not enabled. Not set (and not editable) if Log monitoring is migrated to Davis data on license level. If skipped when editing via PUT method then already set quotas will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_this_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment consumption. Resets each calendar month.",
										Computed:    true,
									},
									"consumed_this_year": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Yearly environment consumption. Resets each year on license creation date anniversary.",
										Computed:    true,
									},
									"monthly_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
									"annual_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"storage": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Environment level storage usage and limit information. Not returned if includeStorageInfo param is not true. If skipped when editing via PUT method then already set limits will remain.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transaction_storage": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retention_reduction_percentage": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Percentage of truncation for new data.",
										Computed:    true,
									},
									"retention_reduction_reason": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Reason of truncation.",
										Computed:    true,
									},
									"currently_used": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Currently used storage [bytes]",
										Computed:    true,
									},
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [bytes].",
										Optional:    true,
									},
								},
							},
						},
						"session_replay_storage": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Session replay storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retention_reduction_percentage": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Percentage of truncation for new data.",
										Computed:    true,
									},
									"retention_reduction_reason": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Reason of truncation.",
										Computed:    true,
									},
									"currently_used": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Currently used storage [bytes]",
										Computed:    true,
									},
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [bytes].",
										Optional:    true,
									},
								},
							},
						},
						"symbol_files_from_mobile_apps": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Symbol files from mobile apps storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Currently used storage [bytes]",
										Computed:    true,
									},
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [bytes].",
										Optional:    true,
									},
								},
							},
						},
						"log_monitoring_storage": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Log monitoring storage usage and limit information on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Currently used storage [bytes]",
										Computed:    true,
									},
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [bytes].",
										Optional:    true,
									},
								},
							},
						},
						"service_request_level_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Service request level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"service_code_level_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Service code level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"real_user_monitoring_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Real user monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"synthetic_monitoring_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Synthetic monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"session_replay_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Session replay retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"log_monitoring_retention": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Session replay retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currently_used_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [milliseconds]",
										Computed:    true,
									},
									"currently_used_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Current data age [days]",
										Computed:    true,
									},
									"max_limit_in_days": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum storage limit [days].",
										Optional:    true,
									},
								},
							},
						},
						"user_actions_per_minute": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Maximum number of user actions generated per minute on environment level. Can be set to any value from 1 to 2147483646 or left unlimited. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum traffic [units per minute]",
										Optional:    true,
									},
								},
							},
						},
						"transaction_traffic_quota": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Maximum number of newly monitored entry point PurePaths captured per process/minute on environment level. Can be set to any value from 100 to 100000. If skipped when editing via PUT method then already set limit will remain.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum traffic [units per minute]",
										Optional:    true,
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

func resourceDynatraceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	en, err := expandEnvironment(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createToken := d.Get("create_token").(bool)

	environment, _, err := dynatraceClusterV2.EnvironmentsApi.CreateEnvironment(authClusterV2).Environment(*en).CreateToken(createToken).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(environment.Id)
	d.Set("api_token", environment.TokenManagementToken)

	resourceDynatraceEnvironmentRead(ctx, d, m)

	return diags

}

func resourceDynatraceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	environment, _, err := dynatraceClusterV2.EnvironmentsApi.GetSingleEnvironment(authClusterV2, environmentID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenEnvironment(environment, d)

	return diags

}

func resourceDynatraceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	en, err := expandEnvironment(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createToken := d.Get("create_token").(bool)

	_, _, err = dynatraceClusterV2.EnvironmentsApi.CreateOrUpdateEnvironment(authClusterV2, environmentID).Environment(*en).CreateToken(createToken).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceEnvironmentRead(ctx, d, m)

}

func resourceDynatraceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	_, err := dynatraceClusterV2.EnvironmentsApi.DeleteEnvironment(authClusterV2, environmentID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
