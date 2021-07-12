package dynatrace

import (
	"encoding/json"
	"log"

	dynatraceClusterV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/cluster/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandClusterUserGroup(d *schema.ResourceData) (*dynatraceClusterV1.GroupConfig, error) {

	var dtGroup dynatraceClusterV1.GroupConfig

	clusterUserGroupID := d.Id()

	if clusterUserGroupID != "" {
		dtGroup.SetId(clusterUserGroupID)
	}

	if name, ok := d.GetOk("name"); ok {
		dtGroup.SetName(name.(string))
	}

	if isClusterAdminGroup, ok := d.GetOk("is_cluster_admin_group"); ok {
		dtGroup.SetIsClusterAdminGroup(isClusterAdminGroup.(bool))
	}

	if hasAccessAccountRole, ok := d.GetOk("has_access_account_role"); ok {
		dtGroup.SetHasAccessAccountRole(hasAccessAccountRole.(bool))
	}

	if hasManageAccountAndViewProductUsageRole, ok := d.GetOk("has_manage_account_and_view_product_usage_role"); ok {
		dtGroup.SetHasManageAccountAndViewProductUsageRole(hasManageAccountAndViewProductUsageRole.(bool))
	}

	if isAccessAccount, ok := d.GetOk("is_access_account"); ok {
		dtGroup.SetIsAccessAccount(isAccessAccount.(bool))
	}

	if isManageAccount, ok := d.GetOk("is_manage_account"); ok {
		dtGroup.SetIsManageAccount(isManageAccount.(bool))
	}

	if ldapGroupNames, ok := d.GetOk("ldap_group_names"); ok {
		dtGroup.SetLdapGroupNames(expandGroupNames(ldapGroupNames.([]interface{})))
	}

	if ssoGroupNames, ok := d.GetOk("sso_group_names"); ok {
		dtGroup.SetSsoGroupNames(expandGroupNames(ssoGroupNames.([]interface{})))
	}

	if accessRight, ok := d.GetOk("access_rights"); ok {
		dtGroup.SetAccessRight(expandAccessRights(accessRight))
	}

	return &dtGroup, nil

}

func expandGroupNames(groups []interface{}) []string {
	pts := make([]string, len(groups))

	for i, v := range groups {
		pts[i] = v.(string)
	}

	return pts

}

func expandAccessRights(key interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(key.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal access rights value %s: %v", key.(string), err)
		return nil
	}

	return val
}

func flattenClusterUserGroup(clusterUserGroup dynatraceClusterV1.GroupConfig, d *schema.ResourceData) diag.Diagnostics {
	d.Set("group_id", clusterUserGroup.Id)
	d.Set("name", clusterUserGroup.Name)
	d.Set("is_cluster_admin_group", clusterUserGroup.IsClusterAdminGroup)
	d.Set("has_access_account_role", clusterUserGroup.HasAccessAccountRole)
	d.Set("has_manage_account_and_view_product_usage_role", clusterUserGroup.HasManageAccountAndViewProductUsageRole)
	d.Set("is_access_account", clusterUserGroup.IsAccessAccount)
	d.Set("is_manage_account", clusterUserGroup.IsManageAccount)
	d.Set("ldap_group_names", flattenGroupNames(*clusterUserGroup.LdapGroupNames))
	d.Set("sso_group_names", flattenGroupNames(*clusterUserGroup.SsoGroupNames))
	d.Set("access_rights", flattenAccessRights(&clusterUserGroup.AccessRight))

	return nil

}

func flattenGroupNames(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}

func flattenAccessRights(value interface{}) interface{} {
	json, err := json.Marshal(value)
	if err != nil {
		log.Printf("[ERROR] Could not marshal access right info value %s: %v", value.(string), err)
		return nil
	}
	return string(json)

}
