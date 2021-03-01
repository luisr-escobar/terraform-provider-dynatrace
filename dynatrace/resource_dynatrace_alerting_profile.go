package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceAlertingProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceAlertingProfileCreate,
		ReadContext:   resourceDynatraceAlertingProfileRead,
		UpdateContext: resourceDynatraceAlertingProfileUpdate,
		DeleteContext: resourceDynatraceAlertingProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the alerting profile, displayed in the UI.",
			},
			"mz_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the management zone to which the alerting profile applies.",
			},
			"rule": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of severity rules. The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops. If you specify both severity rule and event filter, the AND logic applies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity_level": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The severity level to trigger the alert.",
						},
						"tag_filters": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Configuration of the tag filtering of the alerting profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_mode": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The filtering mode.",
									},
									"tag_filter": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "A tag-based filter of monitored entities.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value.",
												},
												"key": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the tag. Custom tags have the tag value here.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of the tag. Not applicable to custom tags.",
												},
											},
										},
									},
								},
							},
						},
						"delay_in_minutes": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Send a notification if a problem remains open longer than X minutes.",
						},
					},
				},
			},
			"event_type_filter": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Configuration of the event filter for the alerting profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"predefined_event_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of a predefined event filter.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"event_type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of the predefined event.",
									},
									"negate": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "The alert triggers when the problem of specified severity arises while the specified event is happening (false) or while the specified event is not happening (true).",
									},
								},
							},
						},
						"custom_event_filter": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of a custom event filter.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_title_filter": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Configuration of a matching filter.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "The filter is enabled (true) or disabled (false).",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value to compare to.",
												},
												"operator": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Operator of the comparison. You can reverse it by setting negate to true.",
												},
												"negate": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Reverses the comparison operator. For example it turns the begins with into does not begin with.",
												},
												"case_insensitive": &schema.Schema{
													Type:        schema.TypeBool,
													Default:     false,
													Optional:    true,
													Description: "The condition is case sensitive (false) or case insensitive (true). If not set, then false is used, making the condition case sensitive.",
												},
											},
										},
									},
									"custom_description_filter": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Configuration of a matching filter.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "The filter is enabled (true) or disabled (false).",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value to compare to.",
												},
												"operator": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Operator of the comparison. You can reverse it by setting negate to true.",
												},
												"negate": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Reverses the comparison operator. For example it turns the begins with into does not begin with.",
												},
												"case_insensitive": &schema.Schema{
													Type:        schema.TypeBool,
													Default:     false,
													Optional:    true,
													Description: "The condition is case sensitive (false) or case insensitive (true). If not set, then false is used, making the condition case sensitive.",
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

func resourceDynatraceAlertingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ap, err := expandAlertingProfile(d)
	if err != nil {
		return diag.FromErr(err)
	}

	alertingProfile, _, err := dynatraceConfigClientV1.AlertingProfilesApi.CreateAlertingProfile(authConfigV1).AlertingProfile(*ap).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace alerting profile",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(alertingProfile.Id)

	resourceDynatraceAlertingProfileRead(ctx, d, m)

	return diags
}

func resourceDynatraceAlertingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	alertingProfileID := d.Id()

	alertingProfile, _, err := dynatraceConfigClientV1.AlertingProfilesApi.GetAlertingProfile(authConfigV1, alertingProfileID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace alerting profile",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	alertingProfileRules := flattenAlertingProfileRulesData(alertingProfile.Rules)
	if err := d.Set("rule", alertingProfileRules); err != nil {
		return diag.FromErr(err)
	}

	alertingProfileEventTypeFilters := flattenAlertingProfileEventTypeFiltersData(alertingProfile.EventTypeFilters)
	if err := d.Set("event_type_filter", alertingProfileEventTypeFilters); err != nil {
		return diag.FromErr(err)
	}

	d.Set("display_name", &alertingProfile.DisplayName)
	d.Set("mz_id", &alertingProfile.MzId)

	return diags
}

func resourceDynatraceAlertingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	alertingProfileID := d.Id()

	if d.HasChange("display_name") || d.HasChange("rule") || d.HasChange("event_type_filter") {

		ap, err := expandAlertingProfile(d)
		if err != nil {
			return diag.FromErr(err)
		}

		_, _, err = dynatraceConfigClientV1.AlertingProfilesApi.CreateOrUpdateAlertingProfile(authConfigV1, alertingProfileID).AlertingProfile(*ap).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update dynatrace alerting profile",
				Detail:   getErrorMessage(err),
			})
			return diags
		}
	}

	return resourceDynatraceAlertingProfileRead(ctx, d, m)

}

func resourceDynatraceAlertingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	alertingProfileID := d.Id()

	_, err := dynatraceConfigClientV1.AlertingProfilesApi.DeleteAlertingProfile(authConfigV1, alertingProfileID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace alerting profile",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
