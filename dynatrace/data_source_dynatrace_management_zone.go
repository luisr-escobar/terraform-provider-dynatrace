package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

func dataSourceDynatraceManagementZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDynatraceManagementZoneRead,
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

func dataSourceDynatraceManagementZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")

	managementZones, _, err := dynatraceConfigClientV1.ManagementZonesApi.ListManagementZones(authConfigV1).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get dynatrace management zones",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	var mz *dynatraceConfigV1.EntityShortRepresentation

	if nameOk {
		for _, a := range managementZones.Values {
			if *a.Name == name.(string) {
				mz = &a
				break
			}
		}

		if mz == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Management Zone not found",
				Detail:   "The value given does not match with any dynatrace management zone name",
			})
			return diags
		}
	}

	if idOk {
		for _, a := range managementZones.Values {
			if a.Id == id.(string) {
				mz = &a
				break
			}
		}

		if mz == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Management Zone not found",
				Detail:   "The value given does not match with any dynatrace management zone id",
			})
			return diags
		}
	}

	flattenManagementZoneData(mz, d)

	return diags
}

func flattenManagementZoneData(a *dynatraceConfigV1.EntityShortRepresentation, d *schema.ResourceData) error {
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
