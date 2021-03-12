# Dynatrace Terraform Provider

The [dynatrace] provider is used to interact with the resources supported by dynatrace. The provider needs to be configured with the proper credentials before it can be used.

## Quick Links

---

* [Getting started]
* [Configure the provider]

## Argument Reference

* `dt_env_url` - (Required) Dynatrace environment URL. SAAS `https://{your-environment-id}.live.dynatrace.com` Managed `https://{your-domain}/e/{your-environment-id}`
* `dt_api_token` - (Required) Dynatrace API Token.

## Example Usage

```hcl
    terraform {
        required_version = ">= 0.13.0"
        required_providers {
            dynatrace = {
                version = "1.0.1"
                source = "dynatrace.com/com/dynatrace"
            }
        }
    }

    provider "dynatrace" {
        dt_env_url   = <your dynatrace environment URL>
        dt_api_token = <your dynatrace API token>
    }

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

## Support

This is provided as an open source project. It does not incude WARRANTY OR SUPPORT, issues can be reported on [GitHub].

[dynatrace]: https://www.dynatrace.com/
[Getting started]: ./guides/getting_started.md
[Configure the provider]: ./guides/provider_configuration.md
[GitHub]: https://github.com/luisr-escobar/terraform-provider-dynatrace/issues