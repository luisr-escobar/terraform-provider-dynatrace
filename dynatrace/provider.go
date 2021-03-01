package dynatrace

import (
	"context"
	"fmt"
	"net/url"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function for Dynatrace API
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dt_env_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_ENV_URL", nil),
			},
			"dt_api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_API_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dynatrace_alerting_profile":           resourceDynatraceAlertingProfile(),
			"dynatrace_management_zone":            resourceDynatraceManagementZone(),
			"dynatrace_maintenance_window":         resourceDynatraceMaintenanceWindow(),
			"dynatrace_dashboard":                  resourceDynatraceDashboard(),
			"dynatrace_auto_tag":                   resourceDynatraceAutoTag(),
			"dynatrace_notification":               resourceDynatraceNotification(),
			"dynatrace_web_application":            resourceDynatraceWebApplication(),
			"dynatrace_application_detection_rule": resourceDynatraceApplicationDetectionRule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dynatrace_alerting_profile": dataSourceDynatraceAlertingProfile(),
			"dynatrace_management_zone":  dataSourceDynatraceManagementZone(),
			"dynatrace_web_application":  dataSourceDynatraceWebApplication(),
		},
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return providerConfigure(ctx, d, p.TerraformVersion)
	}

	return p
}

// ProviderConfiguration stores the Dynatrace API client
type ProviderConfiguration struct {
	DynatraceConfigClientV1 *dynatraceConfigV1.APIClient
	AuthConfigV1            context.Context
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	dtEnvURL := d.Get("dt_env_url").(string)
	apiToken := d.Get("dt_api_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	parsedDTUrl, err := url.Parse(dtEnvURL)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid dynatrace URL",
			Detail:   err.Error(),
		})
	}

	// Initialize the Dynatrace Configuration V1 API client
	authConfigV1 := context.WithValue(
		context.Background(),
		dynatraceConfigV1.ContextAPIKeys,
		map[string]dynatraceConfigV1.APIKey{
			"Api-Token": {
				Key:    apiToken,
				Prefix: "Api-Token",
			},
		},
	)

	authConfigV1 = context.WithValue(authConfigV1, dynatraceConfigV1.ContextServerVariables, map[string]string{
		"name":     string(parsedDTUrl.Host),
		"protocol": string(parsedDTUrl.Scheme),
	})

	configV1 := dynatraceConfigV1.NewConfiguration()
	dynatraceConfigClientV1 := dynatraceConfigV1.NewAPIClient(configV1)

	return &ProviderConfiguration{
		DynatraceConfigClientV1: dynatraceConfigClientV1,
		AuthConfigV1:            authConfigV1,
	}, diags

}

func getErrorMessage(err error) string {
	var errorMessage string

	if apiErr, ok := err.(dynatraceConfigV1.GenericOpenAPIError); ok {
		return fmt.Sprintf("%v: %s", err, apiErr.Body())
	}

	if errURL, ok := err.(*url.Error); ok {
		return fmt.Sprintf("%v:%s", err, errURL)
	}

	return errorMessage
}
