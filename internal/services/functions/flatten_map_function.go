package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlattenMapFunction struct{}

func (f *FlattenMapFunction) Metadata(ctx context.Context, request function.MetadataRequest, response *function.MetadataResponse) {
	response.Name = "flatten_map"
}

func (f *FlattenMapFunction) Definition(ctx context.Context, request function.DefinitionRequest, response *function.DefinitionResponse) {
	response.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType:         types.DynamicType,
				AllowNullValue:      false,
				AllowUnknownValues:  true,
				Name:                "input",
				Description:         "The input map to flatten.",
				MarkdownDescription: "The input map to flatten.",
			},
			function.StringParameter{
				AllowNullValue:      false,
				AllowUnknownValues:  false,
				Name:                "separator",
				Description:         "The separator character to use when combining keys.",
				MarkdownDescription: "The separator character to use when combining keys.",
			},
		},
		VariadicParameter: function.StringParameter{
			AllowNullValue:      false,
			AllowUnknownValues:  false,
			Name:                "attribute_names",
			Description:         "The attribute names to traverse in order (index 0 is top-level attribute, index 1 is next level, etc).",
			MarkdownDescription: "The attribute names to traverse in order (index 0 is top-level attribute, index 1 is next level, etc).",
		},
		Return:      function.MapReturn{ElementType: types.DynamicType},
		Summary:     "Flattens a nested map structure.",
		Description: "Flattens a nested map by traversing specified attributes in order and combining keys with a separator.",
		MarkdownDescription: `Flattens a nested map by traversing specified attributes in order and combining keys with a separator.

## Example Usage

` + "```hcl" + `
variable "infrastructure" {
  type = map(object({
    name    = string
    regions = map(object({
      location = string
      zones    = map(object({
        availability_zone = string
        capacity          = number
      }))
    }))
  }))
  default = {
    "env1" = {
      name = "Production"
      regions = {
        "eastus" = {
          location = "East US"
          zones = {
            "zone1" = { availability_zone = "1", capacity = 100 }
            "zone2" = { availability_zone = "2", capacity = 150 }
          }
        }
        "westus" = {
          location = "West US"
          zones = {
            "zone1" = { availability_zone = "1", capacity = 200 }
          }
        }
      }
    }
  }
}

# Flatten the nested structure traversing 'regions' then 'zones'
output "flattened_zones" {
  value = provider::azapi::flatten_map(var.infrastructure, "/", "regions", "zones")
}

# Output:
# {
#   "env1/eastus/zone1" = { availability_zone = "1", capacity = 100 }
#   "env1/eastus/zone2" = { availability_zone = "2", capacity = 150 }
#   "env1/westus/zone1" = { availability_zone = "1", capacity = 200 }
# }
` + "```" + `

The function traverses the input map and at each level:
1. Takes the current map key
2. Looks for an attribute with the name specified at the current depth in the variadic arguments
3. If found and it's a map, recursively processes each key-value pair with the next attribute name
4. Combines all keys along the path with the separator
5. When all attribute names are exhausted or no matching attribute is found, stores the value`,
	}
}

func (f *FlattenMapFunction) Run(ctx context.Context, request function.RunRequest, response *function.RunResponse) {
	var inputMap types.Map
	var separator types.String
	var attributeNames types.Tuple

	response.Error = function.ConcatFuncErrors(
		request.Arguments.Get(ctx, &inputMap, &separator, &attributeNames),
	)
	if response.Error != nil {
		return
	}

	// Convert tuple to string slice
	attrNameSlice := make([]string, 0)
	for _, elem := range attributeNames.Elements() {
		if strVal, ok := elem.(types.String); ok {
			attrNameSlice = append(attrNameSlice, strVal.ValueString())
		}
	}

	if inputMap.IsUnknown() {
		response.Error = response.Result.Set(ctx, types.MapUnknown(types.DynamicType))
		return
	}

	result := flattenMap(ctx, inputMap, separator.ValueString(), attrNameSlice)
	response.Error = response.Result.Set(ctx, result)
}

var _ function.Function = &FlattenMapFunction{}

func flattenMap(ctx context.Context, input types.Map, separator string, attributeNames []string) types.Map {
	result := make(map[string]attr.Value)

	for key, value := range input.Elements() {
		flattenValue(ctx, key, value, separator, attributeNames, 0, result)
	}

	return types.MapValueMust(types.DynamicType, result)
}

func flattenValue(ctx context.Context, prefix string, value attr.Value, separator string, attributeNames []string, depth int, result map[string]attr.Value) {
	// Unwrap Dynamic values
	if dynVal, ok := value.(types.Dynamic); ok {
		if dynVal.IsNull() || dynVal.IsUnknown() {
			result[prefix] = value
			return
		}
		value = dynVal.UnderlyingValue()
	}

	// If we've exhausted all attribute names, store the value
	if depth >= len(attributeNames) {
		result[prefix] = types.DynamicValue(value)
		return
	}

	// Check if value is an object
	if objVal, ok := value.(types.Object); ok {
		if objVal.IsNull() || objVal.IsUnknown() {
			result[prefix] = value
			return
		}

		attrs := objVal.Attributes()
		attrName := attributeNames[depth]

		// Look for the attribute at this depth
		if attrValue, exists := attrs[attrName]; exists {
			// Unwrap if Dynamic
			actualValue := attrValue
			if dynVal, ok := attrValue.(types.Dynamic); ok {
				if dynVal.IsNull() || dynVal.IsUnknown() {
					result[prefix] = value
					return
				}
				actualValue = dynVal.UnderlyingValue()
			}

			// Check if it's a map
			if mapVal, ok := actualValue.(types.Map); ok {
				if mapVal.IsNull() || mapVal.IsUnknown() {
					result[prefix] = value
					return
				}

				// Traverse into the map with the next attribute name
				for childKey, childValue := range mapVal.Elements() {
					newPrefix := prefix + separator + childKey
					flattenValue(ctx, newPrefix, childValue, separator, attributeNames, depth+1, result)
				}
				return
			}
		}

		// If attribute doesn't exist or isn't a map, store the object as-is
		result[prefix] = types.DynamicValue(objVal)
		return
	}

	// For all other types, store as-is
	result[prefix] = types.DynamicValue(value)
}
