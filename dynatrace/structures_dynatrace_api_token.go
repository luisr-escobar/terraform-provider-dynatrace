package dynatrace

import (
	dynatraceEnvironmentV2 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v2/environment/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandApiTokenCreate(d *schema.ResourceData) (*dynatraceEnvironmentV2.ApiTokenCreate, error) {

	var dtApiToken dynatraceEnvironmentV2.ApiTokenCreate

	if name, ok := d.GetOk("name"); ok {
		dtApiToken.SetName(name.(string))
	}

	if personalAccessToken, ok := d.GetOk("personal_access_token"); ok {
		dtApiToken.SetPersonalAccessToken(personalAccessToken.(bool))
	}

	if expirationDate, ok := d.GetOk("expiration_date"); ok {
		dtApiToken.SetExpirationDate(expirationDate.(string))
	}

	if scopes, ok := d.GetOk("scopes"); ok {
		dtApiToken.SetScopes(expandApiTokenScopes(scopes.([]interface{})))
	}

	return &dtApiToken, nil
}

func expandApiTokenUpdate(d *schema.ResourceData) (*dynatraceEnvironmentV2.ApiTokenUpdate, error) {

	var dtApiToken dynatraceEnvironmentV2.ApiTokenUpdate

	if name, ok := d.GetOk("name"); ok {
		dtApiToken.SetName(name.(string))
	}

	if enabled, ok := d.GetOk("enabled"); ok {
		dtApiToken.SetEnabled(enabled.(bool))
	}

	return &dtApiToken, nil
}

func expandApiTokenScopes(scopes []interface{}) []string {
	pts := make([]string, len(scopes))

	for i, v := range scopes {
		pts[i] = v.(string)
	}

	return pts

}

func flattenApiToken(apiToken dynatraceEnvironmentV2.ApiToken, d *schema.ResourceData) diag.Diagnostics {

	d.Set("name", apiToken.Name)
	d.Set("personal_access_token", apiToken.PersonalAccessToken)
	d.Set("enabled", apiToken.Enabled)
	d.Set("expiration_date", apiToken.ExpirationDate)
	d.Set("scopes", flattenApiTokenScopes(*apiToken.Scopes))
	d.Set("last_used_ip_address", apiToken.LastUsedIpAddress)
	d.Set("last_used_date", apiToken.LastUsedDate)
	d.Set("creation_date", apiToken.CreationDate)
	d.Set("owner", apiToken.Owner)

	return nil

}

func flattenApiTokenScopes(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}
