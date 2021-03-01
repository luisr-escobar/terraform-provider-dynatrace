package dynatrace

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceMaintenanceWindowCreate,
		ReadContext:   resourceDynatraceMaintenanceWindowRead,
		UpdateContext: resourceDynatraceMaintenanceWindowUpdate,
		DeleteContext: resourceDynatraceMaintenanceWindowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the maintenance window, displayed in the UI.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A short description of the maintenance purpose.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the maintenance: planned or unplanned.",
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
			},
			"suppression": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of suppression of alerting and problem detection during the maintenance.",
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
			},
			"scope": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The scope of the maintenance window.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entities": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of Dynatrace entities (for example, hosts or services) to be included in the scope.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"match": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A matching rule for Dynatrace entities.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching.",
										Required:    true,
									},
									"mz_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The ID of a management zone to which the matched entities must belong.",
									},
									"tag_combination": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "OR",
										Description: "The logic that applies when several tags are specified: AND/OR.",
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The tag you want to use for matching.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The origin of the tag, such as AWS or Cloud Foundry.",
													Required:    true,
												},
												"key": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The key of the tag.",
													Required:    true,
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The value of the tag.",
													Optional:    true,
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
			"schedule": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The schedule of the maintenance window.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recurrence_type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The type of the schedule recurrence.",
							Required:    true,
						},
						"start": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.",
							Required:    true,
						},
						"end": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.",
							Required:    true,
						},
						"zone_id": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The time zone of the start and end time. Default time zone is UTC.",
							Required:    true,
						},
						"recurrence": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The recurrence of the maintenance window.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day_of_week": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The day of the week for weekly maintenance.",
										Optional:    true,
									},
									"day_of_month": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The day of the month for monthly maintenance.",
										Optional:    true,
									},
									"start_time": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The start time of the maintenance window in HH:mm format.",
										Required:    true,
									},
									"duration_minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The duration of the maintenance window in minutes.",
										Required:    true,
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

func resourceDynatraceMaintenanceWindowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	mw, err := expandMaintenanceWindow(d)
	if err != nil {
		return diag.FromErr(err)
	}

	maintenanceWindow, _, err := dynatraceConfigClientV1.MaintenanceWindowsApi.CreateMaintenanceWindow(authConfigV1).MaintenanceWindow(*mw).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace maintenance window",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(maintenanceWindow.Id)

	resourceDynatraceMaintenanceWindowRead(ctx, d, m)

	return diags

}

func resourceDynatraceMaintenanceWindowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	maintenanceWindowID := d.Id()

	maintenaceWindow, _, err := dynatraceConfigClientV1.MaintenanceWindowsApi.GetMaintenanceWindow(authConfigV1, maintenanceWindowID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace maintenance window",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	maintenaceWindowScope := flattenMaintenanceWindowScopeData(maintenaceWindow.Scope)
	if err := d.Set("scope", maintenaceWindowScope); err != nil {
		return diag.FromErr(err)
	}

	maintenaceWindowSchedule := flattenMaintenanceWindowScheduleData(&maintenaceWindow.Schedule)
	if err := d.Set("schedule", maintenaceWindowSchedule); err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", &maintenaceWindow.Name)
	d.Set("description", &maintenaceWindow.Description)
	d.Set("type", &maintenaceWindow.Type)
	d.Set("suppression", &maintenaceWindow.Suppression)

	return diags

}

func resourceDynatraceMaintenanceWindowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	maintenanceWindowID := d.Id()

	if d.HasChange("name") || d.HasChange("description") || d.HasChange("type") || d.HasChange("suppression") || d.HasChange("scope") || d.HasChange("schedule") {

		mw, err := expandMaintenanceWindow(d)
		if err != nil {
			return diag.FromErr(err)
		}

		_, _, err = dynatraceConfigClientV1.MaintenanceWindowsApi.UpdateMaintenanceWindow(authConfigV1, maintenanceWindowID).MaintenanceWindow(*mw).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update dynatrace maintenance window",
				Detail:   getErrorMessage(err),
			})
			return diags
		}

	}

	return resourceDynatraceMaintenanceWindowRead(ctx, d, m)

}

func resourceDynatraceMaintenanceWindowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	maintenanceWindowID := d.Id()

	_, err := dynatraceConfigClientV1.MaintenanceWindowsApi.DeleteMaintenanceWindow(authConfigV1, maintenanceWindowID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete maintenance window",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
