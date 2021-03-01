package dynatrace

import (
	"encoding/json"
	"log"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func conditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The key to identify the data we're matching.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The attribute to be used for comparision.",
							Required:    true,
						},
						"dynamic_key": {
							Type:         schema.TypeString,
							Description:  "Dynamic key generated based on selected type/attribute.",
							Optional:     true,
							ValidateFunc: validation.StringIsJSON,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Defines the actual set of fields depending on the value.",
							Optional:    true,
						},
					},
				},
			},
			"comparison_info": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines how the matching is actually performed: what and how are we comparing.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Operator of the comparison. You can reverse it by setting negate to true. Possible values depend on the type of the comparison. Find the list of actual models in the description of the type field and check the description of the model you need.",
							Required:    true,
						},
						"value": {
							Type:         schema.TypeString,
							Description:  "The value to compare to.",
							Optional:     true,
							ValidateFunc: validation.StringIsJSON,
						},
						"negate": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Reverses the comparison operator. For example it turns the begins with into does not begin with.",
							Required:    true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Defines the actual set of fields depending on the value.",
							Required:    true,
						},
						"case_sensitive": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Defines if value to compare to is case sensitive",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func expandConditions(conditions []interface{}) []dynatraceConfigV1.EntityRuleEngineCondition {
	if len(conditions) < 1 {
		return []dynatraceConfigV1.EntityRuleEngineCondition{}
	}

	mcs := make([]dynatraceConfigV1.EntityRuleEngineCondition, len(conditions))

	for i, condition := range conditions {

		m := condition.(map[string]interface{})

		var dtEntityRuleEngineCondition dynatraceConfigV1.EntityRuleEngineCondition

		if key, ok := m["key"].([]interface{}); ok {
			dtEntityRuleEngineCondition.SetKey(expandConditionKey(key))
		}

		if comparisonInfo, ok := m["comparison_info"].([]interface{}); ok {
			dtEntityRuleEngineCondition.SetComparisonInfo(expandConditionComparisonInfo(comparisonInfo))
		}

		mcs[i] = dtEntityRuleEngineCondition
	}

	return mcs
}

func expandConditionKey(conditionKey []interface{}) dynatraceConfigV1.ConditionKey {
	if len(conditionKey) == 0 || conditionKey[0] == nil {
		return dynatraceConfigV1.ConditionKey{}
	}

	dtConditionKey := dynatraceConfigV1.NewConditionKeyWithDefaults()

	m := conditionKey[0].(map[string]interface{})

	if attribute, ok := m["attribute"].(string); ok {
		dtConditionKey.SetAttribute(attribute)
	}

	if dynamicKey, ok := m["dynamic_key"].(string); ok && len(dynamicKey) != 0 {
		dtConditionKey.SetDynamicKey(expandDynamicKey(dynamicKey))
	}

	if ckType, ok := m["type"].(string); ok && len(ckType) != 0 {
		dtConditionKey.SetType(ckType)
	}

	return *dtConditionKey

}

func expandDynamicKey(key interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(key.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal dynamic key value %s: %v", key.(string), err)
		return nil
	}

	return val
}

func expandConditionComparisonInfo(comparisonInfo []interface{}) dynatraceConfigV1.ComparisonBasic {
	if len(comparisonInfo) == 0 || comparisonInfo[0] == nil {
		return dynatraceConfigV1.ComparisonBasic{}
	}

	dtComparsionInfo := dynatraceConfigV1.NewComparisonBasicWithDefaults()

	m := comparisonInfo[0].(map[string]interface{})

	if operator, ok := m["operator"].(string); ok {
		dtComparsionInfo.SetOperator(operator)
	}

	if value, ok := m["value"].(string); ok && len(value) != 0 {
		dtComparsionInfo.SetValue(expandComparisonInfoValue(value))
	}

	if negate, ok := m["negate"].(bool); ok {
		dtComparsionInfo.SetNegate(negate)
	}

	if ciType, ok := m["type"].(string); ok {
		dtComparsionInfo.SetType(ciType)
	}

	if caseSensitive, ok := m["case_sensitive"]; ok && m["operator"].(string) != "EXISTS" {
		dtComparsionInfo.SetCaseSensitive(caseSensitive.(bool))
	}

	return *dtComparsionInfo

}

func expandComparisonInfoValue(value interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(value.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal comparison info value %s: %v", value.(string), err)
		return nil
	}

	return val
}

func flattenConditionsData(conditions *[]dynatraceConfigV1.EntityRuleEngineCondition) []interface{} {
	if conditions != nil {
		mcs := make([]interface{}, len(*conditions))

		for i, conditions := range *conditions {
			mc := make(map[string]interface{})

			mc["key"] = flattenConditionKey(&conditions.Key)
			mc["comparison_info"] = flattenComparisonInfo(&conditions.ComparisonInfo)
			mcs[i] = mc
		}

		return mcs
	}

	return make([]interface{}, 0)

}

func flattenConditionKey(conditionKey *dynatraceConfigV1.ConditionKey) []interface{} {
	if conditionKey == nil {
		return []interface{}{conditionKey}
	}

	k := make(map[string]interface{})

	k["attribute"] = conditionKey.Attribute
	k["type"] = conditionKey.Type
	k["dynamic_key"] = flattenDynamicKey(&conditionKey.DynamicKey)

	return []interface{}{k}
}

func flattenDynamicKey(key interface{}) interface{} {
	json, err := json.Marshal(key)
	if err != nil {
		log.Printf("[ERROR] Could not marshal comparison info value %s: %v", key.(string), err)
		return nil
	}
	return string(json)

}

func flattenComparisonInfo(comparisonInfo *dynatraceConfigV1.ComparisonBasic) []interface{} {
	if comparisonInfo == nil {
		return []interface{}{comparisonInfo}
	}

	c := make(map[string]interface{})

	c["operator"] = comparisonInfo.Operator
	c["value"] = flattenComparisonInfoValue(&comparisonInfo.Value)
	c["negate"] = comparisonInfo.Negate
	c["type"] = comparisonInfo.Type

	return []interface{}{c}
}

func flattenComparisonInfoValue(value interface{}) interface{} {
	json, err := json.Marshal(value)
	if err != nil {
		log.Printf("[ERROR] Could not marshal comparison info value %s: %v", value.(string), err)
		return nil
	}
	return string(json)

}
