---
subcategory: "Microsoft.Logic - Logic Apps"
page_title: "integrationAccounts/schemas"
description: |-
  Manages a Logic App Integration Account Schema.
---

# Microsoft.Logic/integrationAccounts/schemas - Logic App Integration Account Schema

This article demonstrates how to use `azapi` provider to manage the Logic App Integration Account Schema resource in Azure.

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

resource "azapi_resource" "integrationAccount" {
  type      = "Microsoft.Logic/integrationAccounts@2019-05-01"
  parent_id = azapi_resource.resourceGroup.id
  name      = var.resource_name
  location  = var.location
  body = {
    properties = {
    }
    sku = {
      name = "Basic"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

resource "azapi_resource" "schema" {
  type      = "Microsoft.Logic/integrationAccounts/schemas@2019-05-01"
  parent_id = azapi_resource.integrationAccount.id
  name      = var.resource_name
  body = {
    properties = {
      content     = "<xs:schema xmlns:b=\"http://schemas.microsoft.com/BizTalk/2003\"\n           xmlns=\"http://Inbound_EDI.OrderFile\"\n           targetNamespace=\"http://Inbound_EDI.OrderFile\"\n           xmlns:xs=\"http://www.w3.org/2001/XMLSchema\">\n<xs:annotation>\n<xs:appinfo>\n<b:schemaInfo default_pad_char=\" \"\n              count_positions_by_byte=\"false\"\n              parser_optimization=\"speed\"\n              lookahead_depth=\"3\"\n              suppress_empty_nodes=\"false\"\n              generate_empty_nodes=\"true\"\n              allow_early_termination=\"false\"\n              early_terminate_optional_fields=\"false\"\n              allow_message_breakup_of_infix_root=\"false\"\n              compile_parse_tables=\"false\"\n              standard=\"Flat File\"\n              root_reference=\"OrderFile\" />\n<schemaEditorExtension:schemaInfo namespaceAlias=\"b\"\n                                  extensionClass=\"Microsoft.BizTalk.FlatFileExtension.FlatFileExtension\"\n                                  standardName=\"Flat File\"\n                                  xmlns:schemaEditorExtension=\"http://schemas.microsoft.com/BizTalk/2003/SchemaEditorExtensions\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:element name=\"OrderFile\">\n<xs:annotation>\n<xs:appinfo>\n<b:recordInfo structure=\"delimited\"\n              preserve_delimiter_for_empty_data=\"true\"\n              suppress_trailing_delimiters=\"false\"\n              sequence_number=\"1\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:complexType>\n<xs:sequence>\n<xs:annotation>\n<xs:appinfo>\n<b:groupInfo sequence_number=\"0\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:element name=\"Order\">\n<xs:annotation>\n<xs:appinfo>\n<b:recordInfo sequence_number=\"1\"\n              structure=\"delimited\"\n              preserve_delimiter_for_empty_data=\"true\"\n              suppress_trailing_delimiters=\"false\"\n              child_delimiter_type=\"hex\"\n              child_delimiter=\"0x0D 0x0A\"\n              child_order=\"infix\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:complexType>\n<xs:sequence>\n<xs:annotation>\n<xs:appinfo>\n<b:groupInfo sequence_number=\"0\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:element name=\"Header\">\n<xs:annotation>\n<xs:appinfo>\n<b:recordInfo sequence_number=\"1\"\n              structure=\"delimited\"\n              preserve_delimiter_for_empty_data=\"true\"\n              suppress_trailing_delimiters=\"false\"\n              child_delimiter_type=\"char\"\n              child_delimiter=\"|\"\n              child_order=\"infix\"\n              tag_name=\"HDR|\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:complexType>\n<xs:sequence>\n<xs:annotation>\n<xs:appinfo>\n<b:groupInfo sequence_number=\"0\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:element name=\"PODate\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"1\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"PONumber\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo justification=\"left\"\n             sequence_number=\"2\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"CustomerID\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"3\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"CustomerContactName\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"5\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"CustomerContactPhone\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"5\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n</xs:sequence>\n</xs:complexType>\n</xs:element>\n<xs:element minOccurs=\"1\"\n            maxOccurs=\"unbounded\"\n            name=\"LineItems\">\n<xs:annotation>\n<xs:appinfo>\n<b:recordInfo sequence_number=\"2\"\n              structure=\"delimited\"\n              preserve_delimiter_for_empty_data=\"true\"\n              suppress_trailing_delimiters=\"false\"\n              child_delimiter_type=\"char\"\n              child_delimiter=\"|\"\n              child_order=\"infix\"\n              tag_name=\"DTL|\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:complexType>\n<xs:sequence>\n<xs:annotation>\n<xs:appinfo>\n<b:groupInfo sequence_number=\"0\" />\n</xs:appinfo>\n</xs:annotation>\n<xs:element name=\"PONumber\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"1\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"ItemOrdered\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"2\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"Quantity\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"3\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"UOM\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"4\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"Price\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"5\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"ExtendedPrice\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"6\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n<xs:element name=\"Description\"\n            type=\"xs:string\">\n<xs:annotation>\n<xs:appinfo>\n<b:fieldInfo sequence_number=\"7\"\n             justification=\"left\" />\n</xs:appinfo>\n</xs:annotation>\n</xs:element>\n</xs:sequence>\n</xs:complexType>\n</xs:element>\n</xs:sequence>\n</xs:complexType>\n</xs:element>\n</xs:sequence>\n</xs:complexType>\n</xs:element>\n</xs:schema>\n"
      contentType = "application/xml"
      schemaType  = "Xml"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}


```



## Arguments Reference

The following arguments are supported:

* `type` - (Required) The type of the resource. This should be set to `Microsoft.Logic/integrationAccounts/schemas@api-version`. The available api-versions for this resource are: [`2015-08-01-preview`, `2016-06-01`, `2018-07-01-preview`, `2019-05-01`].

* `parent_id` - (Required) The ID of the azure resource in which this resource is created. The allowed values are:  
  `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}`

* `name` - (Required) Specifies the name of the azure resource. Changing this forces a new resource to be created.

* `body` - (Required) Specifies the configuration of the resource. More information about the arguments in `body` can be found in the [Microsoft documentation](https://learn.microsoft.com/en-us/azure/templates/Microsoft.Logic/integrationAccounts/schemas?pivots=deployment-language-terraform).

## Import

 ```shell
 # Azure resource can be imported using the resource id, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}/schemas/{resourceName}
 
 # It also supports specifying API version by using the resource id with api-version as a query parameter, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}/schemas/{resourceName}?api-version=2019-05-01
 ```
