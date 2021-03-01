package dynatrace

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceAutoTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceAutoTagCreate,
		ReadContext:   resourceDynatraceAutoTagRead,
		UpdateContext: resourceDynatraceAutoTagUpdate,
		DeleteContext: resourceDynatraceAutoTagDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the auto-tag, which is applied to entities. Additionally you can specify a valueFormat in the tag rule. In that case the tag is used in the name:valueFormat format. For example you can extend the Infrastructure tag to Infrastructure:Windows and Infrastructure:Linux.",
				Required:    true,
			},
			"rule": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The list of rules for tag usage. When there are multiple rules, the OR logic applies.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Type of entities to which the rule applies.",
							Required:    true,
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Tag rule is enabled (true) or disabled (false).",
							Required:    true,
						},
						"value_format": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The value of the auto-tag. If specified, the tag is used in the name:valueFormat format.",
							Optional:    true,
						},
						"propagation_types": &schema.Schema{
							Type:        schema.TypeList,
							Description: "How to apply the tag to underlying entities.",
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
		},
	}
}

func resourceDynatraceAutoTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	at, err := expandAutoTag(d)
	if err != nil {
		return diag.FromErr(err)
	}

	autoTagBody, err := json.Marshal(at)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("[DEBUG] AUTOTAG JSON BODY IS \n %s \n \n", autoTagBody)

	autoTag, _, err := dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.CreateAutoTag(authConfigV1).AutoTag(*at).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create auto tag",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(autoTag.Id)

	resourceDynatraceAutoTagRead(ctx, d, m)

	return diags

}

func resourceDynatraceAutoTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	autoTagID := d.Id()

	autoTag, _, err := dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.GetAutoTag(authConfigV1, autoTagID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace auto tag",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	autoTagRules := flattenAutoTagRulesData(autoTag.Rules)
	if err := d.Set("rule", autoTagRules); err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", &autoTag.Name)

	return diags

}

func resourceDynatraceAutoTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	autoTagID := d.Id()

	if d.HasChange("name") || d.HasChange("rule") {

		at, err := expandAutoTag(d)
		if err != nil {
			return diag.FromErr(err)
		}

		_, _, err = dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.UpdateAutoTag(authConfigV1, autoTagID).AutoTag(*at).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update dynatrace auto tag",
				Detail:   getErrorMessage(err),
			})
			return diags
		}

	}

	return resourceDynatraceAutoTagRead(ctx, d, m)

}

func resourceDynatraceAutoTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	autoTagID := d.Id()

	_, err := dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.DeleteAutoTag(authConfigV1, autoTagID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete auto tag",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
