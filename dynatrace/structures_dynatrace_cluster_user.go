package dynatrace

import (
	dynatraceClusterV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/cluster/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandClusterUser(d *schema.ResourceData) (*dynatraceClusterV1.UserConfig, error) {

	var dtUser dynatraceClusterV1.UserConfig

	if id, ok := d.GetOk("user_id"); ok {
		dtUser.SetId(id.(string))
	}

	if email, ok := d.GetOk("email"); ok {
		dtUser.SetEmail(email.(string))
	}

	if firstName, ok := d.GetOk("first_name"); ok {
		dtUser.SetFirstName(firstName.(string))
	}

	if lastName, ok := d.GetOk("last_name"); ok {
		dtUser.SetLastName(lastName.(string))
	}

	if password, ok := d.GetOk("password_clear_text"); ok {
		dtUser.SetPasswordClearText(password.(string))
	}

	if groups, ok := d.GetOk("groups"); ok {
		dtUser.SetGroups(expandClusterUserGroups(groups.([]interface{})))
	}

	return &dtUser, nil

}

func expandClusterUserGroups(groups []interface{}) []string {
	pts := make([]string, len(groups))

	for i, v := range groups {
		pts[i] = v.(string)
	}

	return pts

}

func flattenClusterUser(clusterUser dynatraceClusterV1.UserConfig, d *schema.ResourceData) diag.Diagnostics {
	d.Set("user_id", clusterUser.Id)
	d.Set("email", clusterUser.Email)
	d.Set("first_name", clusterUser.FirstName)
	d.Set("last_name", clusterUser.LastName)
	d.Set("password_clear_text", clusterUser.PasswordClearText)
	d.Set("groups", flattenClusterUserGroups(*clusterUser.Groups))

	return nil

}

func flattenClusterUserGroups(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}
