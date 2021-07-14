# Configuring the dynatrace terraform provider

The following methods are supported for configuring the dynatrace terraform provider:

1. [Using the `provider` block]
1. [Setting environment variables] (recommended)

## Configuring using the provider block

[Setting environment variables] is the recommended method, however configuring the provider within your HCL is a quick start option, below is the minimal recommended configuration:

```hcl
terraform {
  required_version = ">= 0.13.0"
  required_providers {
    dynatrace = {
      version = "1.0.2"
      source = "dynatrace.com/com/dynatrace"
    }
  }
}

provider "dynatrace" {
  dt_env_url   = <your dynatrace environment URL>
  dt_api_token = <your dynatrace API token>
}
```

## Configuration by setting environment variables

A set of environment variables will be automatically accessed by the dynatrace terraform provider.

If you're using Terraform locally, you can set the environment variables in your system's startup file, such as your .bash_profile or .bashrc file on UNIX machines.

.bash_profile

```bash
# Managed dynatrace environment
export DYNATRACE_ENV_URL="https://{your-domain}/e/{your-environment-id}"
export DYNATRACE_API_TOKEN="{your-dynatrace-api-token}"
```

```bash
# SaaS dynatrace environment
export DYNATRACE_ENV_URL=" https://{your-environment-id}.live.dynatrace.com"
export DYNATRACE_API_TOKEN="{your-dynatrace-api-token}"
```

Once the environment variables are set, the provider can be started as follows:

```bash
provider "dynatrace" {}
```

## Environment variables reference

Available environment variables and their mapping to the provider's schema attributes:

| <small>Schema Attribute</small> | <small>Equivalent Env Variable</small> | <small>Required?</small> | <small>Default</small> | <small>Description</small>                                                                   |
| ------------------------------- | -------------------------------------- | ------------------------ | ---------------------- | -------------------------------------------------------------------------------------------- |
| `dt_env_url`                    | `DYNATRACE_ENV_URL`                 | required                 | `null`                 | Your dynatrace environment URL.                                                                 |
| `dt_api_token`                       | `DYNATRACE_API_TOKEN`                    | required                 | `null`                 | [Your dynatrace API token.]                                     |

[Setting environment variables]: #configuration-by-setting-environment-variables
[Setting environment variables]: #configuration-by-setting-environment-variables
[Using the `provider` block]: #configuring-using-the-provider-block
[Your dynatrace API token.]: https://www.dynatrace.com/support/help/dynatrace-api/basics/dynatrace-api-authentication/