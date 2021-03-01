---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_alerting_profile Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
Provides a dynatrace alerting profile resource. It allows to create, update, delete alerting profiles in a dynatrace environment. [Alerting profiles API]
---

# dynatrace_alerting_profile (Resource)

Provides a dynatrace alerting profile resource. It allows to create, update, delete alerting profiles in a dynatrace environment. [Alerting profiles API]

## Example Usage

```hcl
resource "dynatrace_alerting_profile" "sockshop_errors" {

  display_name = "sockshop_errors"
  mz_id = dynatrace_management_zone.sockshop_prod.id

  rule{
    severity_level = "AVAILABILITY"
    tag_filters {
      include_mode = "INCLUDE_ALL"
      tag_filter {
        context = "CONTEXTLESS"
        key = "app"
        value = "carts"
      }
      tag_filter {
        context = "CONTEXTLESS"
        key = "env"
        value = "prod"
      }
    }
    delay_in_minutes = 2
  }

  event_type_filter{
    predefined_event_filter{
      negate = true
      event_type = "EC2_HIGH_CPU"
    }
  }

  event_type_filter{
    custom_event_filter{
      custom_title_filter{
        enabled = true
        value = "sockshop"
        operator = "CONTAINS"
        negate = true
        case_insensitive = false
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **display_name** (String) The name of the alerting profile, displayed in the UI.

### Optional

- **event_type_filter** (Block List) Configuration of the event filter for the alerting profile. (see [below for nested schema](#nestedblock--event_type_filter))
- **id** (String) The ID of this resource.
- **mz_id** (String) The ID of the management zone to which the alerting profile applies.
- **rule** (Block List) A list of severity rules. The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops. If you specify both severity rule and event filter, the AND logic applies. (see [below for nested schema](#nestedblock--rule))

<a id="nestedblock--event_type_filter"></a>
### Nested Schema for `event_type_filter`

Optional:

- **custom_event_filter** (Block List) Configuration of a custom event filter. (see [below for nested schema](#nestedblock--event_type_filter--custom_event_filter))
- **predefined_event_filter** (Block List) Configuration of a predefined event filter. (see [below for nested schema](#nestedblock--event_type_filter--predefined_event_filter))

<a id="nestedblock--event_type_filter--custom_event_filter"></a>
### Nested Schema for `event_type_filter.custom_event_filter`

Optional:

- **custom_description_filter** (Block List) Configuration of a matching filter. (see [below for nested schema](#nestedblock--event_type_filter--custom_event_filter--custom_description_filter))
- **custom_title_filter** (Block List) Configuration of a matching filter. (see [below for nested schema](#nestedblock--event_type_filter--custom_event_filter--custom_title_filter))

<a id="nestedblock--event_type_filter--custom_event_filter--custom_description_filter"></a>
### Nested Schema for `event_type_filter.custom_event_filter.custom_description_filter`

Required:

- **enabled** (Boolean) The filter is enabled (true) or disabled (false).
- **negate** (Boolean) Reverses the comparison operator. For example it turns the begins with into does not begin with.
- **operator** (String) Operator of the comparison. You can reverse it by setting negate to true.
- **value** (String) The value to compare to.

Optional:

- **case_insensitive** (Boolean) The condition is case sensitive (false) or case insensitive (true). If not set, then false is used, making the condition case sensitive.


<a id="nestedblock--event_type_filter--custom_event_filter--custom_title_filter"></a>
### Nested Schema for `event_type_filter.custom_event_filter.custom_title_filter`

Required:

- **enabled** (Boolean) The filter is enabled (true) or disabled (false).
- **negate** (Boolean) Reverses the comparison operator. For example it turns the begins with into does not begin with.
- **operator** (String) Operator of the comparison. You can reverse it by setting negate to true.
- **value** (String) The value to compare to.

Optional:

- **case_insensitive** (Boolean) The condition is case sensitive (false) or case insensitive (true). If not set, then false is used, making the condition case sensitive.



<a id="nestedblock--event_type_filter--predefined_event_filter"></a>
### Nested Schema for `event_type_filter.predefined_event_filter`

Required:

- **event_type** (String) The type of the predefined event.
- **negate** (Boolean) The alert triggers when the problem of specified severity arises while the specified event is happening (false) or while the specified event is not happening (true).



<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Required:

- **delay_in_minutes** (Number) Send a notification if a problem remains open longer than X minutes.
- **severity_level** (String) The severity level to trigger the alert.
- **tag_filters** (Block List, Min: 1) Configuration of the tag filtering of the alerting profile. (see [below for nested schema](#nestedblock--rule--tag_filters))

<a id="nestedblock--rule--tag_filters"></a>
### Nested Schema for `rule.tag_filters`

Optional:

- **include_mode** (String) The filtering mode.
- **tag_filter** (Block List) A tag-based filter of monitored entities. (see [below for nested schema](#nestedblock--rule--tag_filters--tag_filter))

<a id="nestedblock--rule--tag_filters--tag_filter"></a>
### Nested Schema for `rule.tag_filters.tag_filter`

Required:

- **context** (String) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value.
- **key** (String) The key of the tag. Custom tags have the tag value here.
- **value** (String) The value of the tag. Not applicable to custom tags.

## Import

Dynatrace alerting profiles can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_alerting_profile.keptn dc228252-2b3d-43ec-b6c5-7bd231adeb6e
```

[Alerting profiles API]: (https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/alerting-profiles-api/post-profile/)