---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_web_application Data Source - terraform-provider-dynatrace"
subcategory: ""
description: |-
Use this data source to get information about a specific web application in dynatrace that already exists
---

# dynatrace_web_application (Data Source)

Use this data source to get information about a specific web application in dynatrace that already exists

```hcl
data "dynatrace_web_application" "carts_prod"{
  name = "Carts Production"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of the Dynatrace entity.
- **name** (String) The name of the Dynatrace entity.