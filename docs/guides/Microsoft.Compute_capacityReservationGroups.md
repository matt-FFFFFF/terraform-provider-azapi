---
subcategory: "Microsoft.Compute - Virtual Machines, Virtual Machine Scale Sets"
page_title: "capacityReservationGroups"
description: |-
  Manages a Capacity Reservation Group.
---

# Microsoft.Compute/capacityReservationGroups - Capacity Reservation Group

This article demonstrates how to use `azapi` provider to manage the Capacity Reservation Group resource in Azure.

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

resource "azapi_resource" "capacityReservationGroup" {
  type                      = "Microsoft.Compute/capacityReservationGroups@2022-03-01"
  parent_id                 = azapi_resource.resourceGroup.id
  name                      = var.resource_name
  location                  = var.location
  schema_validation_enabled = false
  response_export_values    = ["*"]
}


```



## Arguments Reference

The following arguments are supported:

* `type` - (Required) The type of the resource. This should be set to `Microsoft.Compute/capacityReservationGroups@api-version`. The available api-versions for this resource are: [`2021-04-01`, `2021-07-01`, `2021-11-01`, `2022-03-01`, `2022-08-01`, `2022-11-01`, `2023-03-01`, `2023-07-01`, `2023-09-01`, `2024-03-01`, `2024-07-01`].

* `parent_id` - (Required) The ID of the azure resource in which this resource is created. The allowed values are:  
  `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}`

* `name` - (Required) Specifies the name of the azure resource. Changing this forces a new resource to be created.

* `body` - (Required) Specifies the configuration of the resource. More information about the arguments in `body` can be found in the [Microsoft documentation](https://learn.microsoft.com/en-us/azure/templates/Microsoft.Compute/capacityReservationGroups?pivots=deployment-language-terraform).

## Import

 ```shell
 # Azure resource can be imported using the resource id, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{resourceName}
 
 # It also supports specifying API version by using the resource id with api-version as a query parameter, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{resourceName}?api-version=2024-07-01
 ```
