---
subcategory: "Microsoft.Logic - Logic Apps"
page_title: "integrationAccounts/agreements"
description: |-
  Manages a Logic App Integration Account Agreement.
---

# Microsoft.Logic/integrationAccounts/agreements - Logic App Integration Account Agreement

This article demonstrates how to use `azapi` provider to manage the Logic App Integration Account Agreement resource in Azure.

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
      name = "Standard"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

resource "azapi_resource" "partner" {
  type      = "Microsoft.Logic/integrationAccounts/partners@2019-05-01"
  parent_id = azapi_resource.integrationAccount.id
  name      = var.resource_name
  body = {
    properties = {
      content = {
        b2b = {
          businessIdentities = [
            {
              qualifier = "AS2Identity"
              value     = "FabrikamNY"
            },
          ]
        }
      }
      partnerType = "B2B"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

resource "azapi_resource" "partner2" {
  type      = "Microsoft.Logic/integrationAccounts/partners@2019-05-01"
  parent_id = azapi_resource.integrationAccount.id
  name      = "${var.resource_name}another"
  body = {
    properties = {
      content = {
        b2b = {
          businessIdentities = [
            {
              qualifier = "AS2Identity"
              value     = "FabrikamNY"
            },
          ]
        }
      }
      partnerType = "B2B"
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}

resource "azapi_resource" "agreement" {
  type      = "Microsoft.Logic/integrationAccounts/agreements@2019-05-01"
  parent_id = azapi_resource.integrationAccount.id
  name      = var.resource_name
  body = {
    properties = {
      agreementType = "AS2"
      content = {
        aS2 = {
          receiveAgreement = {
            protocolSettings = {
              acknowledgementConnectionSettings = {
                ignoreCertificateNameMismatch = false
                keepHttpConnectionAlive       = false
                supportHttpStatusCodeContinue = false
                unfoldHttpHeaders             = false
              }
              envelopeSettings = {
                autogenerateFileName                    = false
                fileNameTemplate                        = "%FILE().ReceivedFileName%"
                messageContentType                      = "text/plain"
                suspendMessageOnFileNameGenerationError = true
                transmitFileNameInMimeHeader            = false
              }
              errorSettings = {
                resendIfMDNNotReceived  = false
                suspendDuplicateMessage = false
              }
              mdnSettings = {
                dispositionNotificationTo  = "http://localhost"
                micHashingAlgorithm        = "SHA1"
                needMDN                    = false
                sendInboundMDNToMessageBox = true
                sendMDNAsynchronously      = false
                signMDN                    = false
                signOutboundMDNIfOptional  = false
              }
              messageConnectionSettings = {
                ignoreCertificateNameMismatch = false
                keepHttpConnectionAlive       = true
                supportHttpStatusCodeContinue = true
                unfoldHttpHeaders             = true
              }
              securitySettings = {
                enableNRRForInboundDecodedMessages  = false
                enableNRRForInboundEncodedMessages  = false
                enableNRRForInboundMDN              = false
                enableNRRForOutboundDecodedMessages = false
                enableNRRForOutboundEncodedMessages = false
                enableNRRForOutboundMDN             = false
                overrideGroupSigningCertificate     = false
              }
              validationSettings = {
                checkCertificateRevocationListOnReceive = false
                checkCertificateRevocationListOnSend    = false
                checkDuplicateMessage                   = false
                compressMessage                         = false
                encryptMessage                          = false
                encryptionAlgorithm                     = "DES3"
                interchangeDuplicatesValidityDays       = 5
                overrideMessageProperties               = false
                signMessage                             = false
                signingAlgorithm                        = "Default"
              }
            }
            receiverBusinessIdentity = {
              qualifier = "AS2Identity"
              value     = "FabrikamNY"
            }
            senderBusinessIdentity = {
              qualifier = "AS2Identity"
              value     = "FabrikamDC"
            }
          }
          sendAgreement = {
            protocolSettings = {
              acknowledgementConnectionSettings = {
                ignoreCertificateNameMismatch = false
                keepHttpConnectionAlive       = false
                supportHttpStatusCodeContinue = false
                unfoldHttpHeaders             = false
              }
              envelopeSettings = {
                autogenerateFileName                    = false
                fileNameTemplate                        = "%FILE().ReceivedFileName%"
                messageContentType                      = "text/plain"
                suspendMessageOnFileNameGenerationError = true
                transmitFileNameInMimeHeader            = false
              }
              errorSettings = {
                resendIfMDNNotReceived  = false
                suspendDuplicateMessage = false
              }
              mdnSettings = {
                dispositionNotificationTo  = "http://localhost"
                micHashingAlgorithm        = "SHA1"
                needMDN                    = false
                sendInboundMDNToMessageBox = true
                sendMDNAsynchronously      = false
                signMDN                    = false
                signOutboundMDNIfOptional  = false
              }
              messageConnectionSettings = {
                ignoreCertificateNameMismatch = false
                keepHttpConnectionAlive       = true
                supportHttpStatusCodeContinue = true
                unfoldHttpHeaders             = true
              }
              securitySettings = {
                enableNRRForInboundDecodedMessages  = false
                enableNRRForInboundEncodedMessages  = false
                enableNRRForInboundMDN              = false
                enableNRRForOutboundDecodedMessages = false
                enableNRRForOutboundEncodedMessages = false
                enableNRRForOutboundMDN             = false
                overrideGroupSigningCertificate     = false
              }
              validationSettings = {
                checkCertificateRevocationListOnReceive = false
                checkCertificateRevocationListOnSend    = false
                checkDuplicateMessage                   = false
                compressMessage                         = false
                encryptMessage                          = false
                encryptionAlgorithm                     = "DES3"
                interchangeDuplicatesValidityDays       = 5
                overrideMessageProperties               = false
                signMessage                             = false
                signingAlgorithm                        = "Default"
              }
            }
            receiverBusinessIdentity = {
              qualifier = "AS2Identity"
              value     = "FabrikamDC"
            }
            senderBusinessIdentity = {
              qualifier = "AS2Identity"
              value     = "FabrikamNY"
            }
          }
        }
      }
      guestIdentity = {
        qualifier = "AS2Identity"
        value     = "FabrikamDC"
      }
      guestPartner = azapi_resource.partner2.name
      hostIdentity = {
        qualifier = "AS2Identity"
        value     = "FabrikamNY"
      }
      hostPartner = azapi_resource.partner.name
    }
  }
  schema_validation_enabled = false
  response_export_values    = ["*"]
}


```



## Arguments Reference

The following arguments are supported:

* `type` - (Required) The type of the resource. This should be set to `Microsoft.Logic/integrationAccounts/agreements@api-version`. The available api-versions for this resource are: [`2015-08-01-preview`, `2016-06-01`, `2018-07-01-preview`, `2019-05-01`].

* `parent_id` - (Required) The ID of the azure resource in which this resource is created. The allowed values are:  
  `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}`

* `name` - (Required) Specifies the name of the azure resource. Changing this forces a new resource to be created.

* `body` - (Required) Specifies the configuration of the resource. More information about the arguments in `body` can be found in the [Microsoft documentation](https://learn.microsoft.com/en-us/azure/templates/Microsoft.Logic/integrationAccounts/agreements?pivots=deployment-language-terraform).

## Import

 ```shell
 # Azure resource can be imported using the resource id, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}/agreements/{resourceName}
 
 # It also supports specifying API version by using the resource id with api-version as a query parameter, e.g.
 terraform import azapi_resource.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{resourceName}/agreements/{resourceName}?api-version=2019-05-01
 ```
