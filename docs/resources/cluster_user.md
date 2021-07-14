---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_cluster_user Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# dynatrace_cluster_user (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **email** (String) User's email address
- **first_name** (String) User's first name
- **last_name** (String) User's last name
- **user_id** (String) User ID

### Optional

- **groups** (List of String) List of user's user group IDs.
- **id** (String) The ID of this resource.
- **password_clear_text** (String) User's password in a clear text; used only to set initial password

