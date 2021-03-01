package dynatrace

import (
	"context"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDynatraceWebApplication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDynatraceWebApplicationRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the Dynatrace entity.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Dynatrace entity.",
			},
		},
	}
}

func dataSourceDynatraceWebApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")

	webApplications, _, err := dynatraceConfigClientV1.RUMWebApplicationConfigurationApi.ListWebApplicationConfigs(authConfigV1).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get dynatrace web applications",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	var wa *dynatraceConfigV1.EntityShortRepresentation

	if nameOk {
		for _, a := range webApplications.Values {
			if *a.Name == name.(string) {
				wa = &a
				break
			}
		}

		if wa == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Web Application not found",
				Detail:   "The value given does not match with any Dynatrace Web Application name",
			})
			return diags
		}
	}

	if idOk {
		for _, a := range webApplications.Values {
			if a.Id == id.(string) {
				wa = &a
				break
			}
		}

		if wa == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Web Application not found",
				Detail:   "The value given does not match with any Dynatrace Web Application id",
			})
			return diags
		}
	}

	flattenWebApplicationData(wa, d)

	return diags
}

func flattenWebApplicationData(a *dynatraceConfigV1.EntityShortRepresentation, d *schema.ResourceData) error {
	d.SetId(a.Id)
	var err error

	err = d.Set("name", a.Name)
	if err != nil {
		return err
	}

	err = d.Set("id", a.Id)
	if err != nil {
		return err
	}

	return nil
}
