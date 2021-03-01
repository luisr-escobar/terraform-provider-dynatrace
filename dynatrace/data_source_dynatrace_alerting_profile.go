package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

func dataSourceDynatraceAlertingProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDynatraceAlertingProfilesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the Dynatrace entity.",
			},
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Dynatrace entity.",
			},
		},
	}
}

func dataSourceDynatraceAlertingProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("display_name")

	alertingProfiles, _, err := dynatraceConfigClientV1.AlertingProfilesApi.GetAlertingProfiles(authConfigV1).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get dynatrace alerting profiles",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	var ap *dynatraceConfigV1.EntityShortRepresentation

	if nameOk {
		for _, a := range alertingProfiles.Values {
			if *a.Name == name.(string) {
				ap = &a
				break
			}
		}

		if ap == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Alerting profile not found",
				Detail:   "The value given does not match with any dynatrace alerting profile name",
			})
			return diags
		}
	}

	if idOk {
		for _, a := range alertingProfiles.Values {
			if a.Id == id.(string) {
				ap = &a
				break
			}
		}

		if ap == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Alerting profile not found",
				Detail:   "The value given does not match with any dynatrace alerting profile id",
			})
			return diags
		}
	}

	flattenAlertingProfileData(ap, d)

	return diags

}

func flattenAlertingProfileData(a *dynatraceConfigV1.EntityShortRepresentation, d *schema.ResourceData) error {
	d.SetId(a.Id)
	var err error

	err = d.Set("display_name", a.Name)
	if err != nil {
		return err
	}

	err = d.Set("id", a.Id)
	if err != nil {
		return err
	}

	return nil
}
