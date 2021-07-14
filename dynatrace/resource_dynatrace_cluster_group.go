package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDynatraceClusterUserGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceClusterUserGroupCreate,
		ReadContext:   resourceDynatraceClusterUserGroupRead,
		UpdateContext: resourceDynatraceClusterUserGroupUpdate,
		DeleteContext: resourceDynatraceClusterUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Group name",
				Required:    true,
			},
			"is_cluster_admin_group": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If true, then the group has the cluster administrator rights.",
				Optional:    true,
				Default:     false,
			},
			"has_access_account_role": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "User's last name",
				Optional:    true,
				Default:     false,
			},
			"has_manage_account_and_view_product_usage_role": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If true, then the group has the manage account rights.",
				Optional:    true,
			},
			"is_access_account": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "writeOnly: true",
				Optional:    true,
				Default:     false,
			},
			"is_manage_account": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "IwriteOnly: true",
				Optional:    true,
				Default:     false,
			},
			"ldap_group_names": &schema.Schema{
				Type:        schema.TypeList,
				Description: "LDAP group names",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sso_group_names": &schema.Schema{
				Type:        schema.TypeList,
				Description: "SSO group names. If defined it's used to map SSO group name to Dynatrace group name, otherwise mapping is done by group name",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"access_rights": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "Access Rights",
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
			},
		},
	}
}

func resourceDynatraceClusterUserGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cu, err := expandClusterUserGroup(d)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterUserGroup, _, err := dynatraceClusterV1.UserGroupsApi.CreateGroup(authClusterV1).GroupConfig(*cu).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create cluster user group",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(clusterUserGroup.Id)

	resourceDynatraceClusterUserGroupRead(ctx, d, m)

	return diags

}

func resourceDynatraceClusterUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterUserGroupID := d.Id()

	clusterUserGroup, _, err := dynatraceClusterV1.UserGroupsApi.GetGroup(authClusterV1, clusterUserGroupID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read cluster user group",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenClusterUserGroup(clusterUserGroup, d)

	return diags

}

func resourceDynatraceClusterUserGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cu, err := expandClusterUserGroup(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = dynatraceClusterV1.UserGroupsApi.UpdateGroup(authClusterV1).GroupConfig(*cu).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update cluster user group",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceClusterUserGroupRead(ctx, d, m)

}

func resourceDynatraceClusterUserGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterUserGroupID := d.Id()

	_, _, err := dynatraceClusterV1.UserGroupsApi.RemoveGroup(authClusterV1, clusterUserGroupID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete cluster user group",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
