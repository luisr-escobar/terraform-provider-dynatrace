package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceApplicationDetectionRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceApplicationDetectionRuleCreate,
		ReadContext:   resourceDynatraceApplicationDetectionRuleRead,
		UpdateContext: resourceDynatraceApplicationDetectionRuleUpdate,
		DeleteContext: resourceDynatraceApplicationDetectionRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"order": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The order of the rule in the rules list.",
				Optional:    true,
			},
			"application_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Dynatrace entity ID of the application, for example APPLICATION-4A3B43.",
				Required:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The unique name of the Application detection rule.",
				Optional:    true,
			},
			"filter_config": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The condition of an application detection rule.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pattern": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The value to look for.",
							Required:    true,
						},
						"application_match_type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The operator of the matching.",
							Required:    true,
						},
						"application_match_target": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Where to look for the the pattern value.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDynatraceApplicationDetectionRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	dr, err := expandApplicationDetectionRule(d)
	if err != nil {
		return diag.FromErr(err)
	}

	applicationDetectionRule, _, err := dynatraceConfigClientV1.RUMApplicationDetectionRulesApi.CreateApplicationDetectionConfig(authConfigV1).ApplicationDetectionRuleConfig(*dr).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace application detection rule",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(applicationDetectionRule.Id)
	resourceDynatraceApplicationDetectionRuleRead(ctx, d, m)

	return diags

}

func resourceDynatraceApplicationDetectionRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	applicationDetectionRuleID := d.Id()

	applicationDetectionRule, _, err := dynatraceConfigClientV1.RUMApplicationDetectionRulesApi.GetApplicationDetectionConfig(authConfigV1, applicationDetectionRuleID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace application detection rule",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenApplicationDetectionRule(applicationDetectionRule, d)

	return diags

}

func resourceDynatraceApplicationDetectionRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	applicationDetectionRuleID := d.Id()

	dr, err := expandApplicationDetectionRule(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = dynatraceConfigClientV1.RUMApplicationDetectionRulesApi.UpdateApplicationDetectionConfig(authConfigV1, applicationDetectionRuleID).ApplicationDetectionRuleConfig(*dr).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update dynatrace application detection rule",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceApplicationDetectionRuleRead(ctx, d, m)

}

func resourceDynatraceApplicationDetectionRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	applicationDetectionRuleID := d.Id()

	_, err := dynatraceConfigClientV1.RUMApplicationDetectionRulesApi.DeleteApplicationDetectionConfig(authConfigV1, applicationDetectionRuleID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace application detection rule",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
