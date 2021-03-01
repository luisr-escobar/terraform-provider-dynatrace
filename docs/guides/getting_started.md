# Getting started with the dynatrace Terraform provider

## Before you begin

---

1. A dynatrace environment SaaS or Managed is required. If you don't have an environment yet you can request a [free trial].

1. [Terraform], a basic understanding of terraform is required.

1. A [dynatrace API token] with the following scope:
    * Read Configuration
    * Write Configuration

## Using the provider

Please see the [dynatrace provider configuration documentation] to get started with configuring the provider.

## Initialize your terraform configuration

After configuring the provider, you can initialize the terraform configuration:

```bash
$ terraform init
```

After initializing terraform, you can start creating your first configuration file:

`main.tf`

```hcl
terraform {
  required_version = "~> 0.13.0"
  required_providers {
    dynatrace = {
      version = "1.0.0"
      source = "dynatrace.com/com/dynatrace"
    }
  }
}

provider "dynatrace" {
  dt_env_url   = <your dynatrace environment URL>
  dt_api_token = <your dynatrace API token>
}
```

* `Managed` <https://{your-domain}/e/{your-environment-id}>
* `SaaS` <https://{your-environment-id}.live.dynatrace.com>

To verify that the provider is working as expected, run the following commmand:

```bash
$ terraform plan
```

The [terraform plan] command is used to create an execution plan. Terraform performs a refresh, unless explicitly disabled, and then determines what actions are necessary to achieve the desired state specified in the configuration files.

## Add a management zone to a new alerting profile

If you have an existing management zone that you would like to assign to a new alerting profile you can use both the `data` and `resource` blocks. The data block will store the management zone information for terraform to use. In the example below, terraform will query the dynatrace API for amanagement zone name that matches the value provided.

```hcl
data "dynatrace_management_zone" "management_zone_name"{
  name = "my management zone"
}

resource "dynatrace_alerting_profile" "alerting_profile_name" {
  display_name = "my alerting profile"
  mz_id = data.dynatrace_management_zone.management_zone_name.id
  rule{
    severity_level = "AVAILABILITY"
    tag_filters {
      include_mode = "INCLUDE_ALL"
      tag_filter {
        context = "ENVIRONMENT"
        key = "product"
        value = "sockshop"
      }
    }
    delay_in_minutes = 2
  }
}
```

The alerting profile that matches the provided rule and the management zone assigned to it will trigger after 2 minutes.

## Apply your terraform configuration

To apply the changes and provision the specified configurations on your dynatrace environment, run the following command:

```bash
$ terraform apply
```

Answer yes to the prompt to apply the changes.

You can confirm the resources created by navigating to the `settings` menu on dynatrace and then searching for the newly created alerting profile.

To make changes just update your terraform configuration file(s) with the new values and run `terraform apply` again.

To delete the resources created you can run `terraform destroy`.

[free trial]: https://www.dynatrace.com/trial/
[Terraform]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[dynatrace API token]: https://www.dynatrace.com/support/help/dynatrace-api/basics/dynatrace-api-authentication/
[dynatrace provider configuration documentation]: ./provider_configuration.md
[terraform plan]: https://www.terraform.io/docs/commands/plan.html