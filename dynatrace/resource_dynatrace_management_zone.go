package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceManagementZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceManagementZoneCreate,
		ReadContext:   resourceDynatraceManagementZoneRead,
		UpdateContext: resourceDynatraceManagementZoneUpdate,
		DeleteContext: resourceDynatraceManagementZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the management zone.",
				Required:    true,
			},
			"rule": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of rules for management zone usage. Each rule is evaluated independently of all other rules.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The type of Dynatrace entities the management zone can be applied to.",
							Required:    true,
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "The rule is enabled (true) or disabled (false).",
							Required:    true,
						},
						"propagation_types": &schema.Schema{
							Type:        schema.TypeList,
							Description: "How to apply the management zone to underlying entities.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"condition": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of matching rules for the management zone. The management zone applies only if all conditions are fulfilled.",
							Required:    true,
							Elem:        conditionSchema(),
						},
					},
				},
			},
			"dimensional_rule": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of dimensional data rules for management zone usage. If several rules are specified, the OR logic applies.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "The rule is enabled (true) or disabled (false).",
							Required:    true,
						},
						"applies_to": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The target of the rule.",
							Required:    true,
						},
						"condition": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of conditions for the management zone. The management zone applies only if all conditions are fulfilled.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition_type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The type of the condition",
										Required:    true,
									},
									"rule_matcher": &schema.Schema{
										Type:        schema.TypeString,
										Description: "How we compare the values",
										Required:    true,
									},
									"key": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The reference value for comparison. For conditions of the DIMENSION type, specify the key here.",
										Required:    true,
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The value of the dimension. Only applicable when the conditionType is set to DIMENSION.",
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

func resourceDynatraceManagementZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	mz, err := expandManagementZone(d)
	if err != nil {
		return diag.FromErr(err)
	}

	managementZone, _, err := dynatraceConfigClientV1.ManagementZonesApi.CreateManagementZone(authConfigV1).ManagementZone(*mz).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace management zone",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(managementZone.Id)

	resourceDynatraceManagementZoneRead(ctx, d, m)

	return diags
}

func resourceDynatraceManagementZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	managementZone, _, err := dynatraceConfigClientV1.ManagementZonesApi.GetManagementZone(authConfigV1, managementZoneID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace management zone",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	managementZoneRules := flattenManagementZoneRulesData(managementZone.Rules)
	if err := d.Set("rule", managementZoneRules); err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", &managementZone.Name)

	return diags
}

func resourceDynatraceManagementZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	if d.HasChange("name") || d.HasChange("rule") {

		mz, err := expandManagementZone(d)
		if err != nil {
			return diag.FromErr(err)
		}

		_, _, err = dynatraceConfigClientV1.ManagementZonesApi.UpdateManagementZone(authConfigV1, managementZoneID).ManagementZone(*mz).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update dynatrace management zone",
				Detail:   getErrorMessage(err),
			})
			return diags
		}

	}

	return resourceDynatraceManagementZoneRead(ctx, d, m)
}

func resourceDynatraceManagementZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	_, err := dynatraceConfigClientV1.ManagementZonesApi.DeleteManagementZone(authConfigV1, managementZoneID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace management zone",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
