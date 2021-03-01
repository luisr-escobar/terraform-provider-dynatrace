---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_management_zone Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  Provides a dynatrace management zone resource. It allows to create, update, delete management zones in a dynatrace environment. [Management Zones API]
---

# dynatrace_management_zone (Resource)

Provides a dynatrace management zone resource. It allows to create, update, delete management zones in a dynatrace environment. [Management Zones API]

## Example Usage

```hcl
resource "dynatrace_management_zone" "sockshop_prod" {
  name = "sockshop_prod"
  rule{
    type = "SERVICE"
    enabled = true
    propagation_types = [
      "SERVICE_TO_HOST_LIKE",
      "SERVICE_TO_PROCESS_GROUP_LIKE"
    ]
    condition {
      key {
        attribute = "SERVICE_TAGS"
      }
      comparison_info {
        type = "TAG"
        operator = "EQUALS"
        value = jsonencode(
            {
                context = "CONTEXTLESS"
                key     = "env"
                value   = "prod"
            }
        )
        negate = false
      }
    }

    condition {
      key {
        attribute = "HOST_GROUP_NAME"
      }
      comparison_info {
        type = "STRING"
        operator = "BEGINS_WITH"
        value = jsonencode("simpleapp")
        negate = false
        case_sensitive = false
      }
    }
  }
  dimensional_rule {
    enabled = true
    applies_to = "ANY"
    condition {
      condition_type = "DIMENSION"
      rule_matcher = "EQUALS"
      key = "responsetime"
      value = "test"
    }
  }
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the management zone.

### Optional

- **dimensional_rule** (Block List) A list of dimensional data rules for management zone usage. If several rules are specified, the OR logic applies. (see [below for nested schema](#nestedblock--dimensional_rule))
- **id** (String) The ID of this resource.
- **rule** (Block List) A list of rules for management zone usage. Each rule is evaluated independently of all other rules. (see [below for nested schema](#nestedblock--rule))

<a id="nestedblock--dimensional_rule"></a>
### Nested Schema for `dimensional_rule`

Required:

- **applies_to** (String) The target of the rule.
- **condition** (Block List, Min: 1) A list of conditions for the management zone. The management zone applies only if all conditions are fulfilled. (see [below for nested schema](#nestedblock--dimensional_rule--condition))
- **enabled** (Boolean) The rule is enabled (true) or disabled (false).

<a id="nestedblock--dimensional_rule--condition"></a>
### Nested Schema for `dimensional_rule.condition`

Required:

- **condition_type** (String) The type of the condition
- **key** (String) The reference value for comparison. For conditions of the DIMENSION type, specify the key here.
- **rule_matcher** (String) How we compare the values

Optional:

- **value** (String) The value of the dimension. Only applicable when the conditionType is set to DIMENSION.



<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Required:

- **condition** (Block List, Min: 1) A list of matching rules for the management zone. The management zone applies only if all conditions are fulfilled. (see [below for nested schema](#nestedblock--rule--condition))
- **enabled** (Boolean) The rule is enabled (true) or disabled (false).
- **type** (String) The type of Dynatrace entities the management zone can be applied to.

Optional:

- **propagation_types** (Set of String) How to apply the management zone to underlying entities.

<a id="nestedblock--rule--condition"></a>
### Nested Schema for `rule.condition`

Required:

- **comparison_info** (Block List, Min: 1) Defines how the matching is actually performed: what and how are we comparing. (see [below for nested schema](#nestedblock--rule--condition--comparison_info))
- **key** (Block List, Min: 1) The key to identify the data we're matching. (see [below for nested schema](#nestedblock--rule--condition--key))

<a id="nestedblock--rule--condition--comparison_info"></a>
### Nested Schema for `rule.condition.comparison_info`

Required:

- **negate** (Boolean) Reverses the comparison operator. For example it turns the begins with into does not begin with.
- **operator** (String) Operator of the comparison. You can reverse it by setting negate to true. Possible values depend on the type of the comparison. Find the list of actual models in the description of the type field and check the description of the model you need.
- **type** (String) Defines the actual set of fields depending on the value.

Optional:

- **case_sensitive** (Boolean) Defines if value to compare to is case sensitive
- **value** (String) The value to compare to.


<a id="nestedblock--rule--condition--key"></a>
### Nested Schema for `rule.condition.key`

Required:

- **attribute** (String) The attribute to be used for comparision.

Optional:

- **dynamic_key** (String) Dynamic key generated based on selected type/attribute.
- **type** (String) Defines the actual set of fields depending on the value.

## Import

Dynatrace management zones can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_management_zone.keptn_carts -4638826838889583423
```

[Management Zones API]: (https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/management-zones-api/)