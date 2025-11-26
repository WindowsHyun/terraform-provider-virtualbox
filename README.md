[![Build Status](https://github.com/terra-farm/terraform-provider-virtualbox/workflows/CI/badge.svg)](https://github.com/terra-farm/terraform-provider-virtualbox/actions?query=branch%3Amaster)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fterra-farm%2Fterraform-provider-virtualbox.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fterra-farm%2Fterraform-provider-virtualbox?ref=badge_shield)
[![Gitter](https://badges.gitter.im/terra-farm/terraform-provider-virtualbox.svg)](https://gitter.im/terra-farm/terraform-provider-virtualbox?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

# VirtualBox provider for Terraform

Published documentation is located on the [Terraform Registry](https://registry.terraform.io/providers/terra-farm/virtualbox/latest/docs)

## Maintainers Needed

[__We are looking for additional maintainers.__](https://github.com/terra-farm/terraform-provider-virtualbox/discussions/117)

## Usage

```tf
terraform {
  required_providers {
    virtualbox = {
      source = "terra-farm/virtualbox"
      version = "<latest-tag>"
    }
  }
}

provider "virtualbox" {
  # Custom configuration options (added in this fork)
  machine_folder = "/exthdd/terraform/machine"  # Optional: Path where VMs will be stored
  gold_folder    = "/exthdd/terraform/gold"     # Optional: Path where gold images will be stored
}

resource "virtualbox_vm" "vm" {
  name   = "example-vm"
  image  = "ubuntu-2204.box"
  cpus   = 2
  memory = "2048mib"
  
  network_adapter {
    type           = "bridged"
    host_interface = "enp3s0"
    device         = "IntelPro1000MTDesktop"
    mac_address    = "080027F3971A"  # Optional: Fixed MAC address
  }
}
```

## Example

You can find a practical example in the [`/examples` directory](/examples)

If you want to contribute documentation changes, see the [Contribution guide](CONTRIBUTING.md).

## Custom Features (This Fork)

This fork adds the following custom features to the original provider:

### Provider Configuration Options

- **`machine_folder`** (Optional): Specifies the path where VirtualBox machines will be stored.
  - Default: `~/.terraform/virtualbox/machine`
  - Allows Terraform VMs to be stored in a custom location without affecting other VirtualBox VMs

- **`gold_folder`** (Optional): Specifies the path where gold images will be stored.
  - Default: `~/.terraform/virtualbox/gold`
  - Allows gold images to be stored in a custom location

### Resource Configuration Options

- **`mac_address`** (Optional): Fixed MAC address for network adapters.
  - Format: 12-digit hexadecimal (e.g., `080027F3971A`)
  - If not specified, VirtualBox will generate one automatically
  - Useful for maintaining consistent network configurations

### Example Usage

```hcl
terraform {
  required_providers {
    virtualbox = {
      source  = "terra-farm/virtualbox"
      version = "0.2.2-alpha.1"
    }
  }
}

provider "virtualbox" {
  machine_folder = "/exthdd/terraform/machine"
  gold_folder    = "/exthdd/terraform/gold"
}

resource "virtualbox_vm" "k3s_master" {
  name   = "k3s-master-node"
  image  = "ubuntu-2204.box"
  cpus   = 4
  memory = "4096mib"

  network_adapter {
    type           = "bridged"
    host_interface = "enp3s0"
    device         = "IntelPro1000MTDesktop"
    mac_address    = "080027F3971A"  # Fixed MAC address
  }
}
```

### Building and Installation

See [CUSTOM_BUILD.md](CUSTOM_BUILD.md) for detailed build and installation instructions.

## Limitations

- __Experimental provider!__
- We only officially support the latest version of Go, Virtualbox and Terraform. The provider might be compatible and work with other versions
  but we do not provide any level of support for this due to lack of time.
- The defaults here are only tested with the [vagrant insecure (packer) keys](https://github.com/hashicorp/vagrant/tree/master/keys) as the login.
- **Note on existing VMs**: Changing `machine_folder` or `gold_folder` settings does not automatically move existing VMs. Only newly created VMs will use the new paths. To move existing VMs, use `terraform destroy` and recreate them, or manually move the VM files.

## Contributors

Special thanks to all contributors, and [@ccll](https://github.com/ccll) for donating the original project to the terra-farm group!

Inspired by [terraform-provider-vix](https://github.com/hooklift/terraform-provider-vix)

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fterra-farm%2Fterraform-provider-virtualbox.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fterra-farm%2Fterraform-provider-virtualbox?ref=badge_large)
