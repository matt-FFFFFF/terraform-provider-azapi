package functions_test

import (
	"context"
	"testing"

	"github.com/Azure/terraform-provider-azapi/internal/services/functions"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestFlattenMapFunction(t *testing.T) {
	testCases := []struct {
		name          string
		request       function.RunRequest
		expected      attr.Value
		expectedError *function.FuncError
	}{
		{
			name: "simple map - no nesting",
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"key1": types.DynamicValue(types.StringValue("value1")),
							"key2": types.DynamicValue(types.StringValue("value2")),
						},
					),
					types.StringValue("-"),
				}),
			},
			expected: types.MapValueMust(
				types.DynamicType,
				map[string]attr.Value{
					"key1": types.DynamicValue(types.StringValue("value1")),
					"key2": types.DynamicValue(types.StringValue("value2")),
				},
			),
		},
		{
			name: "nested map - one level",
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"parent1": types.DynamicValue(types.ObjectValueMust(
								map[string]attr.Type{
									"name":     types.StringType,
									"children": types.MapType{ElemType: types.DynamicType},
								},
								map[string]attr.Value{
									"name": types.StringValue("Parent 1"),
									"children": types.MapValueMust(
										types.DynamicType,
										map[string]attr.Value{
											"child1": types.DynamicValue(types.StringValue("Child 1 Value")),
											"child2": types.DynamicValue(types.StringValue("Child 2 Value")),
										},
									),
								},
							)),
							"parent2": types.DynamicValue(types.ObjectValueMust(
								map[string]attr.Type{
									"name":     types.StringType,
									"children": types.MapType{ElemType: types.DynamicType},
								},
								map[string]attr.Value{
									"name": types.StringValue("Parent 2"),
									"children": types.MapValueMust(
										types.DynamicType,
										map[string]attr.Value{
											"child3": types.DynamicValue(types.StringValue("Child 3 Value")),
										},
									),
								},
							)),
						},
					),
					types.StringValue("-"),
					types.StringValue("children"),
				}),
			},
			expected: types.MapValueMust(
				types.DynamicType,
				map[string]attr.Value{
					"parent1-child1": types.DynamicValue(types.StringValue("Child 1 Value")),
					"parent1-child2": types.DynamicValue(types.StringValue("Child 2 Value")),
					"parent2-child3": types.DynamicValue(types.StringValue("Child 3 Value")),
				},
			),
		},
		{
			name: "nested map - two levels",
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"region1": types.DynamicValue(types.ObjectValueMust(
								map[string]attr.Type{
									"name": types.StringType,
									"zones": types.MapType{ElemType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"name": types.StringType,
											"instances": types.MapType{ElemType: types.DynamicType},
										},
									}},
								},
								map[string]attr.Value{
									"name": types.StringValue("East US"),
									"zones": types.MapValueMust(
										types.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name": types.StringType,
												"instances": types.MapType{ElemType: types.DynamicType},
											},
										},
										map[string]attr.Value{
											"zone1": types.ObjectValueMust(
												map[string]attr.Type{
													"name": types.StringType,
													"instances": types.MapType{ElemType: types.DynamicType},
												},
												map[string]attr.Value{
													"name": types.StringValue("Zone 1"),
													"instances": types.MapValueMust(
														types.DynamicType,
														map[string]attr.Value{
															"vm1": types.DynamicValue(types.StringValue("VM 1 Details")),
															"vm2": types.DynamicValue(types.StringValue("VM 2 Details")),
														},
													),
												},
											),
											"zone2": types.ObjectValueMust(
												map[string]attr.Type{
													"name": types.StringType,
													"instances": types.MapType{ElemType: types.DynamicType},
												},
												map[string]attr.Value{
													"name": types.StringValue("Zone 2"),
													"instances": types.MapValueMust(
														types.DynamicType,
														map[string]attr.Value{
															"vm3": types.DynamicValue(types.StringValue("VM 3 Details")),
														},
													),
												},
											),
										},
									),
								},
							)),
						},
					),
					types.StringValue("/"),
					types.StringValue("zones"),
					types.StringValue("instances"),
				}),
			},
			expected: types.MapValueMust(
				types.DynamicType,
				map[string]attr.Value{
					"region1/zone1/vm1": types.DynamicValue(types.StringValue("VM 1 Details")),
					"region1/zone1/vm2": types.DynamicValue(types.StringValue("VM 2 Details")),
					"region1/zone2/vm3": types.DynamicValue(types.StringValue("VM 3 Details")),
				},
			),
		},
		{
			name: "custom separator",
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"root": types.DynamicValue(types.ObjectValueMust(
								map[string]attr.Type{
									"nested": types.MapType{ElemType: types.DynamicType},
								},
								map[string]attr.Value{
									"nested": types.MapValueMust(
										types.DynamicType,
										map[string]attr.Value{
											"leaf": types.DynamicValue(types.StringValue("value")),
										},
									),
								},
							)),
						},
					),
					types.StringValue("::"),
					types.StringValue("nested"),
				}),
			},
			expected: types.MapValueMust(
				types.DynamicType,
				map[string]attr.Value{
					"root::leaf": types.DynamicValue(types.StringValue("value")),
				},
			),
		},
		{
			name: "complex values preserved",
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"parent": types.DynamicValue(types.ObjectValueMust(
								map[string]attr.Type{
									"items": types.MapType{ElemType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"id":     types.NumberType,
											"tags":   types.ListType{ElemType: types.StringType},
											"config": types.MapType{ElemType: types.StringType},
										},
									}},
								},
								map[string]attr.Value{
									"items": types.MapValueMust(
										types.ObjectType{
											AttrTypes: map[string]attr.Type{
												"id":     types.NumberType,
												"tags":   types.ListType{ElemType: types.StringType},
												"config": types.MapType{ElemType: types.StringType},
											},
										},
										map[string]attr.Value{
											"item1": types.ObjectValueMust(
												map[string]attr.Type{
													"id":     types.NumberType,
													"tags":   types.ListType{ElemType: types.StringType},
													"config": types.MapType{ElemType: types.StringType},
												},
												map[string]attr.Value{
													"id": types.NumberValue(nil),
													"tags": types.ListValueMust(
														types.StringType,
														[]attr.Value{
															types.StringValue("tag1"),
															types.StringValue("tag2"),
														},
													),
													"config": types.MapValueMust(
														types.StringType,
														map[string]attr.Value{
															"key": types.StringValue("value"),
														},
													),
												},
											),
										},
									),
								},
							)),
						},
					),
					types.StringValue("-"),
					types.StringValue("items"),
				}),
			},
			expected: types.MapValueMust(
				types.DynamicType,
				map[string]attr.Value{
					"parent-item1": types.DynamicValue(types.ObjectValueMust(
						map[string]attr.Type{
							"id":     types.NumberType,
							"tags":   types.ListType{ElemType: types.StringType},
							"config": types.MapType{ElemType: types.StringType},
						},
						map[string]attr.Value{
							"id": types.NumberValue(nil),
							"tags": types.ListValueMust(
								types.StringType,
								[]attr.Value{
									types.StringValue("tag1"),
									types.StringValue("tag2"),
								},
							),
							"config": types.MapValueMust(
								types.StringType,
								map[string]attr.Value{
									"key": types.StringValue("value"),
								},
							),
						},
					)),
				},
			),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := function.RunResponse{
				Result: function.NewResultData(types.MapUnknown(types.DynamicType)),
			}

			flattenMapFunction := functions.FlattenMapFunction{}
			flattenMapFunction.Run(context.Background(), testCase.request, &got)

			if testCase.expectedError != nil {
				if got.Error == nil {
					t.Fatal("expected error, got none")
				}
				if diff := cmp.Diff(got.Error, testCase.expectedError); diff != "" {
					t.Errorf("unexpected error difference: %s", diff)
				}
				return
			}

			if got.Error != nil {
				t.Fatalf("unexpected error: %s", got.Error)
			}

			result := got.Result.Value()

			if !testCase.expected.Equal(result) {
				t.Errorf("expected %v, got %v", testCase.expected, result)
			}
		})
	}
}
