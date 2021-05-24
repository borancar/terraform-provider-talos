# Terraform provider for Talos

The Terraform Talos provider is a plugin for Terraform that allows creating
Talos configs to assign to machines.

## Install

Run:

```
make
make install
```

This will put the `terraform-provider-talos` into the [Implied Local Mirror
Directory](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)
so that the next `terraform init` can pick it up. This is until the plugin is
finalized and submitted to the Terraform registry.
