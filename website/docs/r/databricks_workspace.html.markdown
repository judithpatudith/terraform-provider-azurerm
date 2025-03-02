---
subcategory: "Databricks"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_databricks_workspace"
description: |-
  Manages a Databricks Workspace
---

# azurerm_databricks_workspace

~> **NOTE:** Some Databricks Workspace features are in Private Preview(e.g. Private Link Endpoint, Customer Managed Keys for Managed Services, etc.) and potentially subject to breaking change without notice. If you would like to use these features please contact your Microsoft support representative on how to opt-in to the Databricks Workspace Private Preview feature program.

Manages a Databricks Workspace

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_databricks_workspace" "example" {
  name                = "databricks-test"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku                 = "standard"

  tags = {
    Environment = "Production"
  }
}
```

-> You can use [the Databricks Terraform Provider](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs) to manage resources within the Databricks Workspace.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.

* `location` - (Required) Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.

* `load_balancer_backend_address_pool_id` - (Optional) Resource ID of the Outbound Load balancer Backend Address Pool for Secure Cluster Connectivity (No Public IP) workspace. Changing this forces a new resource to be created.

* `sku` - (Required) The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.

~> **NOTE** Downgrading to a `trial sku` from a `standard` or `premium sku` will force a new resource to be created.

* `managed_services_cmk_key_vault_key_id` - (Optional) Customer managed encryption properties for the Databricks Workspace managed resources(e.g. Notebooks and Artifacts). Changing this forces a new resource to be created.

* `managed_resource_group_name` - (Optional) The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.

~> **NOTE** Make sure that this field is unique if you have multiple Databrick Workspaces deployed in your subscription and choose to not have the `managed_resource_group_name` auto generated by the Azure Resource Provider. Having multiple Databrick Workspaces deployed in the same subscription with the same `manage_resource_group_name` may result in some resources that cannot be deleted.

* `customer_managed_key_enabled` - (Optional) Is the workspace enabled for customer managed key encryption? If `true` this enables the Managed Identity for the managed storage account. Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.

* `infrastructure_encryption_enabled`- (Optional) Is the Databricks File System root file system enabled with a secondary layer of encryption with platform managed keys? Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.

* `public_network_access_enabled` - (Optional) Allow public access for accessing workspace. Set value to `false` to access workspace only via private link endpoint. Possible values include `true` or `false`. Defaults to `true`. Changing this forces a new resource to be created.

* `network_security_group_rules_required` - (Optional) Does the data plane (clusters) to control plane communication happen over private link endpoint only or publicly? Possible values `AllRules`, `NoAzureDatabricksRules` or `NoAzureServiceRules`. Required when `public_network_access_enabled` is set to `false`. Changing this forces a new resource to be created.

* `custom_parameters` - (Optional) A `custom_parameters` block as documented below.

* `tags` - (Optional) A mapping of tags to assign to the resource.

---

A `custom_parameters` block supports the following:

* `aml_workspace_id` - (Optional) The ID of a Azure Machine Learning workspace to link with Databricks workspace. Changing this forces a new resource to be created.

* `nat_gateway_name` - (Optional) Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets. Defaults to `nat-gateway`. Changing this forces a new resource to be created.

* `public_ip_name` - (Optional) Name of the Public IP for No Public IP workspace with managed vNet. Defaults to `nat-gw-public-ip`. Changing this forces a new resource to be created.

* `no_public_ip` - (Optional) Are public IP Addresses not allowed? Possible values are `true` or `false`. Defaults to `false`. Changing this forces a new resource to be created.

* `public_subnet_name` - (Optional) The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.

* `public_subnet_network_security_group_association_id` - (Optional) The resource ID of the `azurerm_subnet_network_security_group_association` resource which is referred to by the `public_subnet_name` field. Required if `virtual_network_id` is set.

* `private_subnet_name` - (Optional) The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.

* `private_subnet_network_security_group_association_id` - (Optional) The resource ID of the `azurerm_subnet_network_security_group_association` resource which is referred to by the `private_subnet_name` field. Required if `virtual_network_id` is set.

* `storage_account_name` - (Optional) Default Databricks File Storage account name. Defaults to a randomized name(e.g. `dbstoragel6mfeghoe5kxu`). Changing this forces a new resource to be created.

* `storage_account_sku_name` - (Optional) Storage account SKU name. Possible values inclued`Standard_LRS`, `Standard_GRS`, `Standard_RAGRS`, `Standard_GZRS`, `Standard_RAGZRS`, `Standard_ZRS`, `Premium_LRS` or `Premium_ZRS`. Defaults to `Standard_GRS`. Changing this forces a new resource to be created.

* `virtual_network_id` - (Optional) The ID of a Virtual Network where this Databricks Cluster should be created. Changing this forces a new resource to be created.

* `vnet_address_prefix` - (Optional) Address prefix for Managed virtual network. Defaults to `10.139`. Changing this forces a new resource to be created.

~> **NOTE** Databricks requires that a network security group is associated with the `public` and `private` subnets when a `virtual_network_id` has been defined. Both `public` and `private` subnets must be delegated to `Microsoft.Databricks/workspaces`. For more information about subnet delegation see the [product documentation](https://docs.microsoft.com/azure/virtual-network/subnet-delegation-overview).


## Example HCL Configurations

* [Databricks Workspace Secure Connectivity Cluster with Load Balancer](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/databricks/secure-connectivity-cluster/with-load-balancer)
* [Databricks Workspace Secure Connectivity Cluster without Load Balancer](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/databricks/secure-connectivity-cluster/without-load-balancer)
* [Databricks Workspace with Private Endpoint](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/private-endpoint/databricks/private-endpoint)
* [Databricks Workspace with Private Endpoint, Customer Managed Keys for Managed Services and Databricks File System Customer Managed Keys](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/private-endpoint/databricks/managed-services)
* [Databricks Workspace with Databricks File System Customer Managed Keys](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/databricks/customer-managed-key/dbfs)
* [Databricks Workspace with Customer Managed Keys for Managed Services](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/databricks/customer-managed-key/managed-services)


## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Databricks Workspace in the Azure management plane.

* `managed_resource_group_id` - The ID of the Managed Resource Group created by the Databricks Workspace.

* `workspace_url` - The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'

* `workspace_id` - The unique identifier of the databricks workspace in Databricks control plane.

* `storage_account_identity` - A `storage_account_identity` block as documented below.

---

A `storage_account_identity` block exports the following:

* `principal_id` - The principal UUID for the internal databricks storage account needed to provide access to the workspace for enabling Customer Managed Keys.

* `tenant_id` - The UUID of the tenant where the internal databricks storage account was created.

* `type` - The type of the internal databricks storage account.


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Databricks Workspace.
* `update` - (Defaults to 30 minutes) Used when updating the Databricks Workspace.
* `read` - (Defaults to 5 minutes) Used when retrieving the Databricks Workspace.
* `delete` - (Defaults to 30 minutes) Used when deleting the Databricks Workspace.

## Import

Databrick Workspaces can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_databricks_workspace.workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
```
