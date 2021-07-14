# Dynatrace Terraform Provider

- Terraform Website: [https://www.terraform.io](https://www.terraform.io)
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## Requirements

- [Terraform] 0.12+

## Using the provider

If you want to run Terraform with the dynatrace provider plugin on your system, complete the following steps:

1. [Download the dynatrace provider plugin for Terraform].

1. Unzip the release archive to extract the plugin binary (`terraform-provider-dynatrace_vX.Y.Z`).

For Terraform version 0.12.x

1. Move the binary into the Terraform [plugins directory] for the platform.
    - Linux/Unix/macOS: `~/.terraform.d/plugins`
    - Windows: `%APPDATA%\terraform.d\plugins`

1. Add the plug-in provider to the Terraform configuration file.

    ```hcl
    terraform {
        required_providers {
            dynatrace = {
                version = "1.0.2"
            }
        }
    }
    ```

For Terraform version 0.13.x

1. Move the binary into the Terraform [plugins directory] for the platform.
    - Linux: `~/.terraform.d/plugins/dynatrace.com/com/dynatrace/1.0.2/linux_amd64/`
    - macOS: `~/.terraform.d/plugins/dynatrace.com/com/dynatrace/1.0.2/darwin_amd64/`
    - Windows: `%APPDATA%\terraform.d\plugins\dynatrace.com\com\dynatrace\1.0.2\windows_amd64\`

1. Add the plug-in provider to the Terraform configuration file.

    ```hcl
    terraform {
        required_version = "~> 0.13.0"
        required_providers {
            dynatrace = {
                version = "1.0.2"
                source = "dynatrace.com/com/dynatrace"
            }
        }
    }
    ```

See the [dynatrace provider documentation] to learn more about how to use it and configure it.

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/luisr-escobar/terraform-provider-dynatrace`

```sh
mkdir -p $GOPATH/src/github.com/luisr-escobar; cd $GOPATH/src/github.com/luisr-escobar
git clone https://github.com/luisr-escobar/terraform-provider-dynatrace.git
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/luisr-escobar/terraform-provider-dynatrace
make build
```

## Contributing to the provider

To contribute to the provider, [Go](http://www.golang.org) is required to be installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

### Development Environment

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
make build

$GOPATH/bin/terraform-provider-dynatrace
```

In order to test the provider, you can simply run `make test`.

```sh
make test
```

In order to run the full suite of Acceptance tests, run make testacc.

**Note:** Acceptance tests create real resources, and often cost money to run.

```sh
make testacc
```

[Terraform]: https://www.terraform.io/downloads.html
[Go]: https://golang.org/doc/install
[Download and install Terraform for your system]: https://www.terraform.io/intro/getting-started/install.html
[Download the dynatrace provider plugin for Terraform]: https://github.com/luisr-escobar/terraform-provider-dynatrace/releases
[plugins directory]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
[API token]: https://www.dynatrace.com/support/help/dynatrace-api/basics/dynatrace-api-authentication/
[dynatrace provider documentation]: ./docs/index.md