terraform {
  required_providers {
    azapi = {
      source = "Azure/azapi"
    }
  }
}

provider "azapi" {
}

data "azapi_client_config" "current" {}

output "subscription_id" {
  value = data.azapi_client_config.current.subscription_id
}

output "subscription_resource_id" {
  value = data.azapi_client_config.current.subscription_resource_id
}

output "tenant_id" {
  value = data.azapi_client_config.current.tenant_id
}
