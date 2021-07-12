package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceEnvironmentCreate,
		ReadContext:   resourceDynatraceEnvironmentRead,
		UpdateContext: resourceDynatraceEnvironmentUpdate,
		DeleteContext: resourceDynatraceEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The display name of the environment.",
				Required:    true,
			},
			"create_token": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If true, a token management token with the scopes 'apiTokens.read' and 'apiTokens.write' and 'TenantTokenManagement'is created when creating a new environment.",
				Optional:    true,
				Default:     true,
			},
			"trial": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that.",
				Optional:    true,
				Default:     false,
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Indicates whether the environment is enabled or disabled. The default value is ENABLED.",
				Optional:    true,
				Default:     "ENABLED",
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Environment Token",
				Optional:    true,
				Computed:    true,
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDynatraceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	en, err := expandEnvironment(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createToken := d.Get("create_token").(bool)

	environment, _, err := dynatraceClusterV2.EnvironmentsApi.CreateEnvironment(authClusterV2).Environment(*en).CreateToken(createToken).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(environment.Id)
	d.Set("api_token", environment.TokenManagementToken)

	resourceDynatraceEnvironmentRead(ctx, d, m)

	return diags

}

func resourceDynatraceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	environment, _, err := dynatraceClusterV2.EnvironmentsApi.GetSingleEnvironment(authClusterV2, environmentID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenEnvironment(environment, d)

	return diags

}

func resourceDynatraceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	en, err := expandEnvironment(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createToken := d.Get("create_token").(bool)

	_, _, err = dynatraceClusterV2.EnvironmentsApi.CreateOrUpdateEnvironment(authClusterV2, environmentID).Environment(*en).CreateToken(createToken).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceEnvironmentRead(ctx, d, m)

}

func resourceDynatraceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceClusterV2 := providerConf.DynatraceClusterClientV2
	authClusterV2 := providerConf.AuthClusterV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Id()

	_, err := dynatraceClusterV2.EnvironmentsApi.DeleteEnvironment(authClusterV2, environmentID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete Dynatrace environment",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags

}
