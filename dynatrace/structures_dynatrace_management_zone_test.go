package dynatrace

import (
	"reflect"
	"testing"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

var keyType = "PROCESS_CUSTOM_METADATA_KEY"
var dynamicKey interface{} = "ENVIRONMENT"

func TestFlattenConditionKey(t *testing.T) {
	cases := []struct {
		Input          *dynatraceConfigV1.ConditionKey
		ExpectedOutput []interface{}
	}{
		{
			&dynatraceConfigV1.ConditionKey{
				Attribute:  "SERVICE_TAGS",
				DynamicKey: &dynamicKey,
				Type:       &keyType,
			},
			[]interface{}{
				map[string]interface{}{
					"attribute":   "SERVICE_TAGS",
					"dynamic_key": "\"ENVIRONMENT\"",
					"type":        &keyType,
				},
			},
		},
	}
	for _, tc := range cases {
		output := flattenConditionKey(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandConditionKey(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput dynatraceConfigV1.ConditionKey
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"attribute":   "SERVICE_TAGS",
					"dynamic_key": "\"ENVIRONMENT\"",
					"type":        "PROCESS_CUSTOM_METADATA_KEY",
				},
			},
			dynatraceConfigV1.ConditionKey{
				Attribute:  "SERVICE_TAGS",
				DynamicKey: &dynamicKey,
				Type:       &keyType,
			},
		},
	}
	for _, tc := range cases {
		output := expandConditionKey(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
