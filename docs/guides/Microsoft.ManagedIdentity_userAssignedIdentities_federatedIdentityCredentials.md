---
subcategory: "Microsoft.ManagedIdentity - Managed identities for Azure resources"
page_title: "userAssignedIdentities/federatedIdentityCredentials"
description: |-
  Manages a Federated Identity Credential.
---

# Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials - Federated Identity Credential

This article demonstrates how to use `azapi` provider to manage the Federated Identity Credential resource in Azure.

## Example Usage

### default

```hcl
terraform {
  required_providers {
    azapi = {
      source = "Azure/azapi"
    }
  }
}

provider "azapi" {
  skip_provider_registration = false
}

variable "resource_name" {
  type    = string
  default = "acctest0001"
}

variable "location" {
  type    = string
  default = "westeurope"
}

resource "azapi_resource" "resourceGroup" {
  type     = "Microsoft.Resources/resourceGroups@2020-06-01"
  name     = var.resource_name
  location = var.location
}

resource "azapi_resource" "userAssignedIdentity" {
  type      = "Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31"
  parent_id = azapi_resource.resourceGroup.id
  name      = var.resource_name
  location  = var.location
}

resource "azapi_resource" "federatedIdentityCredential" {
  type      = "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials@2022-01-31-preview"
  parent_id = azapi_resource.userAssignedIdentity.id
  name      = var.resource_name
  location  = var.location
  body = {
    properties = {
      issuer    = "https://foo"
      subject   = "foo"
      audiences = ["foo"]
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

```



## Arguments Reference

The following arguments are supported:

* `type` - (Required) The type of the resource. This should be set to `Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials@api-version`. The available api-versions for this resource are: [`2022-01-31-preview`, `2023-01-31`, `2023-07-31-preview`, `2024-11-30`].

* `parent_id` - (Required) The ID of the azure resource in which this resource is created. The allowed values are:  
  `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}`

* `name` - (Required) Specifies the name of the azure resource. Changing this forces a new resource to be created.

* `body` - (Required) Specifies the configuration of the resource. More information about the arguments in `body` can be found in the [Microsoft documentation](https://learn.microsoft.com/en-us/azure/templates/Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials?pivots=deployment-language-terraform).

## Import

 ```shell
 # Azure resource can be imported using the resource id, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{resourceName}
 
 # It also supports specifying API version by using the resource id with api-version as a query parameter, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{resourceName}?api-version=2024-11-30
 ```
