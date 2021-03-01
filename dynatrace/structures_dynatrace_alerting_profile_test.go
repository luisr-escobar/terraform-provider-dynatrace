package dynatrace

import (
	"reflect"
	"testing"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

func TestFlattenAlertingCustomTextFilter(t *testing.T) {
	cases := []struct {
		Input          *dynatraceConfigV1.AlertingCustomTextFilter
		ExpectedOutput []interface{}
	}{
		{
			&dynatraceConfigV1.AlertingCustomTextFilter{
				Enabled:         true,
				Value:           "sockshop",
				Operator:        "CONTAINS",
				Negate:          true,
				CaseInsensitive: false,
			},
			[]interface{}{
				map[string]interface{}{
					"enabled":          true,
					"value":            "sockshop",
					"operator":         "CONTAINS",
					"negate":           true,
					"case_insensitive": false,
				},
			},
		},
	}
	for _, tc := range cases {
		output := flattenCustomTextFilter(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandAlertingCustomTextFilter(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *dynatraceConfigV1.AlertingCustomTextFilter
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"enabled":          true,
					"value":            "sockshop",
					"operator":         "CONTAINS",
					"negate":           true,
					"case_insensitive": false,
				},
			},
			&dynatraceConfigV1.AlertingCustomTextFilter{
				Enabled:         true,
				Value:           "sockshop",
				Operator:        "CONTAINS",
				Negate:          true,
				CaseInsensitive: false,
			},
		},
	}
	for _, tc := range cases {
		output := expandCustomTextFilter(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
