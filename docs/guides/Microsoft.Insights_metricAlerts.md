---
subcategory: "Microsoft.Insights - Azure Monitor"
page_title: "metricAlerts"
description: |-
  Manages a Metric Alert within Azure Monitor.
---

# Microsoft.Insights/metricAlerts - Metric Alert within Azure Monitor

This article demonstrates how to use `azapi` provider to manage the Metric Alert within Azure Monitor resource in Azure.

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

resource "azapi_resource" "storageAccount" {
  type      = "Microsoft.Storage/storageAccounts@2021-09-01"
  parent_id = azapi_resource.resourceGroup.id
  name      = var.resource_name
  location  = var.location
  body = {
    kind = "StorageV2"
    properties = {
      accessTier                   = "Hot"
      allowBlobPublicAccess        = true
      allowCrossTenantReplication  = true
      allowSharedKeyAccess         = true
      defaultToOAuthAuthentication = false
      encryption = {
        keySource = "Microsoft.Storage"
        services = {
          queue = {
            keyType = "Service"
          }
          table = {
            keyType = "Service"
          }
        }
      }
      isHnsEnabled      = false
      isNfsV3Enabled    = false
      isSftpEnabled     = false
      minimumTlsVersion = "TLS1_2"
      networkAcls = {
        defaultAction = "Allow"
      }
      publicNetworkAccess      = "Enabled"
      supportsHttpsTrafficOnly = true
    }
    sku = {
      name = "Standard_LRS"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

resource "azapi_resource" "metricAlert" {
  type      = "Microsoft.Insights/metricAlerts@2018-03-01"
  parent_id = azapi_resource.resourceGroup.id
  name      = var.resource_name
  location  = "global"
  body = {
    properties = {
      actions = [
      ]
      autoMitigate = true
      criteria = {
        allOf = [
          {
            criterionType = "StaticThresholdCriterion"
            dimensions = [
            ]
            metricName           = "UsedCapacity"
            metricNamespace      = "Microsoft.Storage/storageAccounts"
            name                 = "Metric1"
            operator             = "GreaterThan"
            skipMetricValidation = false
            threshold            = 55.5
            timeAggregation      = "Average"
          },
        ]
        "odata.type" = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria"
      }
      description         = ""
      enabled             = true
      evaluationFrequency = "PT1M"
      scopes = [
        azapi_resource.storageAccount.id,
      ]
      severity             = 3
      targetResourceRegion = ""
      targetResourceType   = ""
      windowSize           = "PT1H"
    }
    tags = {
      CUSTOMER  = "CUSTOMERx"
      Example   = "Example123"
      terraform = "Coolllll"
      test      = "123"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}


```



## Arguments Reference

The following arguments are supported:

* `type` - (Required) The type of the resource. This should be set to `Microsoft.Insights/metricAlerts@api-version`. The available api-versions for this resource are: [`2018-03-01`].

* `parent_id` - (Required) The ID of the azure resource in which this resource is created. The allowed values are:  
  `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}`

* `name` - (Required) Specifies the name of the azure resource. Changing this forces a new resource to be created.

* `body` - (Required) Specifies the configuration of the resource. More information about the arguments in `body` can be found in the [Microsoft documentation](https://learn.microsoft.com/en-us/azure/templates/Microsoft.Insights/metricAlerts?pivots=deployment-language-terraform).

## Import

 ```shell
 # Azure resource can be imported using the resource id, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{resourceName}
 
 # It also supports specifying API version by using the resource id with api-version as a query parameter, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{resourceName}?api-version=2018-03-01
 ```
