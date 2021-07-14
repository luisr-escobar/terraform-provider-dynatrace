package dynatrace

import (
	dynatraceClusterV2 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v2/cluster/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandEnvironment(d *schema.ResourceData) (*dynatraceClusterV2.Environment, error) {

	var dtEnvironment dynatraceClusterV2.Environment

	if name, ok := d.GetOk("name"); ok {
		dtEnvironment.SetName(name.(string))
	}

	if trial, ok := d.GetOk("trial"); ok {
		dtEnvironment.SetTrial(trial.(bool))
	}

	if state, ok := d.GetOk("state"); ok {
		dtEnvironment.SetState(state.(string))
	}

	if tags, ok := d.GetOk("tags"); ok {
		dtEnvironment.SetTags(expandEnvironmentTags(tags.([]interface{})))
	}

	return &dtEnvironment, nil

}

func expandEnvironmentTags(tags []interface{}) []string {
	pts := make([]string, len(tags))

	for i, v := range tags {
		pts[i] = v.(string)
	}

	return pts

}

func flattenEnvironment(environment dynatraceClusterV2.Environment, d *schema.ResourceData) diag.Diagnostics {

	d.Set("name", environment.Name)
	d.Set("trial", environment.Trial)
	d.Set("state", environment.State)
	d.Set("tags", flattenEnvironmentTags(*environment.Tags))

	return nil

}

func flattenEnvironmentTags(values []string) []string {
	if values == nil {
		return nil
	}

	dvs := make([]string, len(values))

	for i, e := range values {
		dvs[i] = e
	}

	return dvs
}
