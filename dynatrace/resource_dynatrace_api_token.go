package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceApiToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceApiTokenCreate,
		ReadContext:   resourceDynatraceApiTokenRead,
		UpdateContext: resourceDynatraceApiTokenUpdate,
		DeleteContext: resourceDynatraceApiTokenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the token.",
				Required:    true,
			},
			"personal_access_token": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The token is a personal access token (true) or an API token (false).",
				Optional:    true,
				Default:     true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The token is enabled (true) or disabled (false).",
				Optional:    true,
				Default:     true,
			},
			"expiration_date": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The expiration date of the token.",
				Optional:    true,
			},
			"scopes": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of the scopes to be assigned to the token.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_used_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token last used IP address.",
				Optional:    true,
				Computed:    true,
			},
			"last_used_date": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token last used date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')",
				Optional:    true,
				Computed:    true,
			},
			"creation_date": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token creation date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')",
				Optional:    true,
				Computed:    true,
			},
			"owner": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The owner of the token",
				Optional:    true,
				Computed:    true,
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The secret of the token",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceDynatraceApiTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceEnvironmentClientV2 := providerConf.DynatraceEnvironmentClientV2
	authEnvironmentV2 := providerConf.AuthEnvironmentV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	at, err := expandApiTokenCreate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	apiToken, _, err := dynatraceEnvironmentClientV2.AccessTokensAPITokensApi.CreateApiToken(authEnvironmentV2).ApiTokenCreate(*at).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create api token",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(*apiToken.Id)
	d.Set("api_token", apiToken.Token)

	resourceDynatraceApiTokenRead(ctx, d, m)

	return diags

}

func resourceDynatraceApiTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceEnvironmentClientV2 := providerConf.DynatraceEnvironmentClientV2
	authEnvironmentV2 := providerConf.AuthEnvironmentV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	apiTokenID := d.Id()

	apiToken, _, err := dynatraceEnvironmentClientV2.AccessTokensAPITokensApi.GetApiToken(authEnvironmentV2, apiTokenID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read api token",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenApiToken(apiToken, d)

	return diags

}

func resourceDynatraceApiTokenUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceEnvironmentClientV2 := providerConf.DynatraceEnvironmentClientV2
	authEnvironmentV2 := providerConf.AuthEnvironmentV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	apiTokenID := d.Id()

	at, err := expandApiTokenUpdate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = dynatraceEnvironmentClientV2.AccessTokensAPITokensApi.UpdateApiToken(authEnvironmentV2, apiTokenID).ApiTokenUpdate(*at).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update api token",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceEnvironmentRead(ctx, d, m)

}

func resourceDynatraceApiTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceEnvironmentClientV2 := providerConf.DynatraceEnvironmentClientV2
	authEnvironmentV2 := providerConf.AuthEnvironmentV2

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	apiTokenID := d.Id()

	_, err := dynatraceEnvironmentClientV2.AccessTokensAPITokensApi.DeleteApiToken(authEnvironmentV2, apiTokenID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete api token",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
