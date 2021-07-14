package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceClusterUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceClusterUserCreate,
		ReadContext:   resourceDynatraceClusterUserRead,
		UpdateContext: resourceDynatraceClusterUserUpdate,
		DeleteContext: resourceDynatraceClusterUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "User ID",
				Required:    true,
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Description: "User's email address",
				Required:    true,
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "User's first name",
				Required:    true,
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "User's last name",
				Required:    true,
			},
			"password_clear_text": &schema.Schema{
				Type:        schema.TypeString,
				Description: "User's password in a clear text; used only to set initial password",
				Optional:    true,
			},
			"groups": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of user's user group IDs.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDynatraceClusterUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cu, err := expandClusterUser(d)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterUser, _, err := dynatraceClusterV1.UsersApi.CreateUser(authClusterV1).UserConfig(*cu).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create cluster user",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(clusterUser.Id)

	resourceDynatraceClusterUserRead(ctx, d, m)

	return diags

}

func resourceDynatraceClusterUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterUserID := d.Id()

	clusterUser, _, err := dynatraceClusterV1.UsersApi.GetUser(authClusterV1, clusterUserID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read cluster user",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenClusterUser(clusterUser, d)

	return diags

}

func resourceDynatraceClusterUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cu, err := expandClusterUser(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = dynatraceClusterV1.UsersApi.UpdateUser(authClusterV1).UserConfig(*cu).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update cluster user",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceClusterUserRead(ctx, d, m)

}

func resourceDynatraceClusterUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV1 := providerConf.DynatraceClusterClientV1
	authClusterV1 := providerConf.AuthClusterV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterUserID := d.Id()

	_, _, err := dynatraceClusterV1.UsersApi.RemoveUser(authClusterV1, clusterUserID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete cluster user",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
