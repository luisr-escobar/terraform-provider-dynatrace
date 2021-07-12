package dynatrace

import (
	"context"
	"fmt"
	"net/url"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	dynatraceClusterV2 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v2/cluster/dynatrace"
	dynatraceEnvironmentV2 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v2/environment/dynatrace"
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
			"dt_cluster_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_CLUSTER_URL", nil),
			},
			"dt_cluster_api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_CLUSTER_API_TOKEN", nil),
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
			"dynatrace_environment":                resourceDynatraceEnvironment(),
			"dynatrace_api_token":                  resourceDynatraceApiToken(),
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
	DynatraceConfigClientV1      *dynatraceConfigV1.APIClient
	DynatraceClusterClientV2     *dynatraceClusterV2.APIClient
	DynatraceEnvironmentClientV2 *dynatraceEnvironmentV2.APIClient
	AuthConfigV1                 context.Context
	AuthClusterV2                context.Context
	AuthEnvironmentV2            context.Context
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	dtEnvURL := d.Get("dt_env_url").(string)
	dtClusterURL := d.Get("dt_cluster_url").(string)
	apiToken := d.Get("dt_api_token").(string)
	clusterApiToken := d.Get("dt_cluster_api_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	parsedDTUrl, err := url.Parse(dtEnvURL)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid Dynatrace Environment URL",
			Detail:   err.Error(),
		})
	}

	parsedDTClusterUrl, err := url.Parse(dtClusterURL)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid Dynatrace Cluster URL",
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
		"name":     string(parsedDTUrl.Host + parsedDTUrl.Path),
		"protocol": string(parsedDTUrl.Scheme),
	})

	configV1 := dynatraceConfigV1.NewConfiguration()
	dynatraceConfigClientV1 := dynatraceConfigV1.NewAPIClient(configV1)

	// Initialize the Dynatrace Cluster V2 API client
	authClusterV2 := context.WithValue(
		context.Background(),
		dynatraceClusterV2.ContextAPIKeys,
		map[string]dynatraceClusterV2.APIKey{
			"Api-Token": {
				Key:    clusterApiToken,
				Prefix: "Api-Token",
			},
		},
	)

	authClusterV2 = context.WithValue(authClusterV2, dynatraceClusterV2.ContextServerVariables, map[string]string{
		"name":     string(parsedDTClusterUrl.Host),
		"protocol": string(parsedDTClusterUrl.Scheme),
	})

	clusterV2 := dynatraceClusterV2.NewConfiguration()
	dynatraceClusterClientV2 := dynatraceClusterV2.NewAPIClient(clusterV2)

	// Initialize the Dynatrace Environment V2 API client
	authEnvironmentV2 := context.WithValue(
		context.Background(),
		dynatraceEnvironmentV2.ContextAPIKeys,
		map[string]dynatraceEnvironmentV2.APIKey{
			"Api-Token": {
				Key:    apiToken,
				Prefix: "Api-Token",
			},
		},
	)

	authEnvironmentV2 = context.WithValue(authEnvironmentV2, dynatraceEnvironmentV2.ContextServerVariables, map[string]string{
		"name":     string(parsedDTUrl.Host + parsedDTUrl.Path),
		"protocol": string(parsedDTUrl.Scheme),
	})

	environmentV2 := dynatraceEnvironmentV2.NewConfiguration()
	dynatraceEnvironmentClientV2 := dynatraceEnvironmentV2.NewAPIClient(environmentV2)

	return &ProviderConfiguration{
		DynatraceConfigClientV1:      dynatraceConfigClientV1,
		DynatraceClusterClientV2:     dynatraceClusterClientV2,
		DynatraceEnvironmentClientV2: dynatraceEnvironmentClientV2,
		AuthConfigV1:                 authConfigV1,
		AuthClusterV2:                authClusterV2,
		AuthEnvironmentV2:            authEnvironmentV2,
	}, diags

}

func getErrorMessage(err error) string {
	var errorMessage string

	if apiErr, ok := err.(dynatraceConfigV1.GenericOpenAPIError); ok {
		return fmt.Sprintf("%v: %s", err, apiErr.Body())
	}

	if apiErr, ok := err.(dynatraceClusterV2.GenericOpenAPIError); ok {
		return fmt.Sprintf("%v: %s", err, apiErr.Body())
	}

	if apiErr, ok := err.(dynatraceEnvironmentV2.GenericOpenAPIError); ok {
		return fmt.Sprintf("%v: %s", err, apiErr.Body())
	}

	if errURL, ok := err.(*url.Error); ok {
		return fmt.Sprintf("%v:%s", err, errURL)
	}

	return errorMessage
}
