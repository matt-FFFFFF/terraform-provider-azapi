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
