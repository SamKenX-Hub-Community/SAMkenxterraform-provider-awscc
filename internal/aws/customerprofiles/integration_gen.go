// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_customerprofiles_integration", integrationResourceType)
}

// integrationResourceType returns the Terraform awscc_customerprofiles_integration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CustomerProfiles::Integration resource type.
func integrationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got created",
			//   "type": "string"
			// }
			Description: "The time of this integration got created",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the domain.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			// DomainName is a force-new attribute.
		},
		"flow_definition": {
			// Property: FlowDefinition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Description": {
			//       "maxLength": 2048,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "FlowName": {
			//       "maxLength": 256,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "KmsArn": {
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SourceFlowConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ConnectorProfileName": {
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "ConnectorType": {
			//           "enum": [
			//             "Salesforce",
			//             "Marketo",
			//             "ServiceNow",
			//             "Zendesk",
			//             "S3"
			//           ],
			//           "type": "string"
			//         },
			//         "IncrementalPullConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "DatetimeTypeFieldName": {
			//               "maxLength": 256,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SourceConnectorProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Marketo": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "S3": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "BucketName": {
			//                   "maxLength": 63,
			//                   "minLength": 3,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "BucketPrefix": {
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "BucketName"
			//               ],
			//               "type": "object"
			//             },
			//             "Salesforce": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "EnableDynamicFieldUpdate": {
			//                   "type": "boolean"
			//                 },
			//                 "IncludeDeletedRecords": {
			//                   "type": "boolean"
			//                 },
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "ServiceNow": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "Zendesk": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "ConnectorType",
			//         "SourceConnectorProperties"
			//       ],
			//       "type": "object"
			//     },
			//     "Tasks": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ConnectorOperator": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Marketo": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "S3": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Salesforce": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "ServiceNow": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Zendesk": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "GREATER_THAN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "DestinationField": {
			//             "maxLength": 256,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "SourceFields": {
			//             "items": {
			//               "maxLength": 2048,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "TaskProperties": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "OperatorPropertyKey": {
			//                   "enum": [
			//                     "VALUE",
			//                     "VALUES",
			//                     "DATA_TYPE",
			//                     "UPPER_BOUND",
			//                     "LOWER_BOUND",
			//                     "SOURCE_DATA_TYPE",
			//                     "DESTINATION_DATA_TYPE",
			//                     "VALIDATION_ACTION",
			//                     "MASK_VALUE",
			//                     "MASK_LENGTH",
			//                     "TRUNCATE_LENGTH",
			//                     "MATH_OPERATION_FIELDS_ORDER",
			//                     "CONCAT_FORMAT",
			//                     "SUBFIELD_CATEGORY_MAP"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Property": {
			//                   "maxLength": 2048,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "OperatorPropertyKey",
			//                 "Property"
			//               ],
			//               "type": "object"
			//             },
			//             "type": "array"
			//           },
			//           "TaskType": {
			//             "enum": [
			//               "Arithmetic",
			//               "Filter",
			//               "Map",
			//               "Mask",
			//               "Merge",
			//               "Truncate",
			//               "Validate"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "SourceFields",
			//           "TaskType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "TriggerConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "TriggerProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Scheduled": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "DataPullMode": {
			//                   "enum": [
			//                     "Incremental",
			//                     "Complete"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "FirstExecutionFrom": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleEndTime": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleExpression": {
			//                   "maxLength": 256,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "ScheduleOffset": {
			//                   "maximum": 36000,
			//                   "minimum": 0,
			//                   "type": "integer"
			//                 },
			//                 "ScheduleStartTime": {
			//                   "type": "number"
			//                 },
			//                 "Timezone": {
			//                   "maxLength": 256,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "ScheduleExpression"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "TriggerType": {
			//           "enum": [
			//             "Scheduled",
			//             "Event",
			//             "OnDemand"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TriggerType"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "FlowName",
			//     "KmsArn",
			//     "Tasks",
			//     "TriggerConfig",
			//     "SourceFlowConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 2048),
						},
					},
					"flow_name": {
						// Property: FlowName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
					"kms_arn": {
						// Property: KmsArn
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(20, 2048),
						},
					},
					"source_flow_config": {
						// Property: SourceFlowConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"connector_profile_name": {
									// Property: ConnectorProfileName
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 256),
									},
								},
								"connector_type": {
									// Property: ConnectorType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"Salesforce",
											"Marketo",
											"ServiceNow",
											"Zendesk",
											"S3",
										}),
									},
								},
								"incremental_pull_config": {
									// Property: IncrementalPullConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"datetime_type_field_name": {
												// Property: DatetimeTypeFieldName
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
										},
									),
									Optional: true,
								},
								"source_connector_properties": {
									// Property: SourceConnectorProperties
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"marketo": {
												// Property: Marketo
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 512),
															},
														},
													},
												),
												Optional: true,
											},
											"s3": {
												// Property: S3
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"bucket_name": {
															// Property: BucketName
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(3, 63),
															},
														},
														"bucket_prefix": {
															// Property: BucketPrefix
															Type:     types.StringType,
															Optional: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 512),
															},
														},
													},
												),
												Optional: true,
											},
											"salesforce": {
												// Property: Salesforce
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"enable_dynamic_field_update": {
															// Property: EnableDynamicFieldUpdate
															Type:     types.BoolType,
															Optional: true,
														},
														"include_deleted_records": {
															// Property: IncludeDeletedRecords
															Type:     types.BoolType,
															Optional: true,
														},
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 512),
															},
														},
													},
												),
												Optional: true,
											},
											"service_now": {
												// Property: ServiceNow
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 512),
															},
														},
													},
												),
												Optional: true,
											},
											"zendesk": {
												// Property: Zendesk
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 512),
															},
														},
													},
												),
												Optional: true,
											},
										},
									),
									Required: true,
								},
							},
						),
						Required: true,
					},
					"tasks": {
						// Property: Tasks
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"connector_operator": {
									// Property: ConnectorOperator
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"marketo": {
												// Property: Marketo
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PROJECTION",
														"LESS_THAN",
														"GREATER_THAN",
														"BETWEEN",
														"ADDITION",
														"MULTIPLICATION",
														"DIVISION",
														"SUBTRACTION",
														"MASK_ALL",
														"MASK_FIRST_N",
														"MASK_LAST_N",
														"VALIDATE_NON_NULL",
														"VALIDATE_NON_ZERO",
														"VALIDATE_NON_NEGATIVE",
														"VALIDATE_NUMERIC",
														"NO_OP",
													}),
												},
											},
											"s3": {
												// Property: S3
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PROJECTION",
														"LESS_THAN",
														"GREATER_THAN",
														"BETWEEN",
														"LESS_THAN_OR_EQUAL_TO",
														"GREATER_THAN_OR_EQUAL_TO",
														"EQUAL_TO",
														"NOT_EQUAL_TO",
														"ADDITION",
														"MULTIPLICATION",
														"DIVISION",
														"SUBTRACTION",
														"MASK_ALL",
														"MASK_FIRST_N",
														"MASK_LAST_N",
														"VALIDATE_NON_NULL",
														"VALIDATE_NON_ZERO",
														"VALIDATE_NON_NEGATIVE",
														"VALIDATE_NUMERIC",
														"NO_OP",
													}),
												},
											},
											"salesforce": {
												// Property: Salesforce
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PROJECTION",
														"LESS_THAN",
														"GREATER_THAN",
														"CONTAINS",
														"BETWEEN",
														"LESS_THAN_OR_EQUAL_TO",
														"GREATER_THAN_OR_EQUAL_TO",
														"EQUAL_TO",
														"NOT_EQUAL_TO",
														"ADDITION",
														"MULTIPLICATION",
														"DIVISION",
														"SUBTRACTION",
														"MASK_ALL",
														"MASK_FIRST_N",
														"MASK_LAST_N",
														"VALIDATE_NON_NULL",
														"VALIDATE_NON_ZERO",
														"VALIDATE_NON_NEGATIVE",
														"VALIDATE_NUMERIC",
														"NO_OP",
													}),
												},
											},
											"service_now": {
												// Property: ServiceNow
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PROJECTION",
														"LESS_THAN",
														"GREATER_THAN",
														"CONTAINS",
														"BETWEEN",
														"LESS_THAN_OR_EQUAL_TO",
														"GREATER_THAN_OR_EQUAL_TO",
														"EQUAL_TO",
														"NOT_EQUAL_TO",
														"ADDITION",
														"MULTIPLICATION",
														"DIVISION",
														"SUBTRACTION",
														"MASK_ALL",
														"MASK_FIRST_N",
														"MASK_LAST_N",
														"VALIDATE_NON_NULL",
														"VALIDATE_NON_ZERO",
														"VALIDATE_NON_NEGATIVE",
														"VALIDATE_NUMERIC",
														"NO_OP",
													}),
												},
											},
											"zendesk": {
												// Property: Zendesk
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PROJECTION",
														"GREATER_THAN",
														"ADDITION",
														"MULTIPLICATION",
														"DIVISION",
														"SUBTRACTION",
														"MASK_ALL",
														"MASK_FIRST_N",
														"MASK_LAST_N",
														"VALIDATE_NON_NULL",
														"VALIDATE_NON_ZERO",
														"VALIDATE_NON_NEGATIVE",
														"VALIDATE_NUMERIC",
														"NO_OP",
													}),
												},
											},
										},
									),
									Optional: true,
								},
								"destination_field": {
									// Property: DestinationField
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 256),
									},
								},
								"source_fields": {
									// Property: SourceFields
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
								},
								"task_properties": {
									// Property: TaskProperties
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"operator_property_key": {
												// Property: OperatorPropertyKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"VALUE",
														"VALUES",
														"DATA_TYPE",
														"UPPER_BOUND",
														"LOWER_BOUND",
														"SOURCE_DATA_TYPE",
														"DESTINATION_DATA_TYPE",
														"VALIDATION_ACTION",
														"MASK_VALUE",
														"MASK_LENGTH",
														"TRUNCATE_LENGTH",
														"MATH_OPERATION_FIELDS_ORDER",
														"CONCAT_FORMAT",
														"SUBFIELD_CATEGORY_MAP",
													}),
												},
											},
											"property": {
												// Property: Property
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 2048),
												},
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"task_type": {
									// Property: TaskType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"Arithmetic",
											"Filter",
											"Map",
											"Mask",
											"Merge",
											"Truncate",
											"Validate",
										}),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
					},
					"trigger_config": {
						// Property: TriggerConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"trigger_properties": {
									// Property: TriggerProperties
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"scheduled": {
												// Property: Scheduled
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"data_pull_mode": {
															// Property: DataPullMode
															Type:     types.StringType,
															Optional: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringInSlice([]string{
																	"Incremental",
																	"Complete",
																}),
															},
														},
														"first_execution_from": {
															// Property: FirstExecutionFrom
															Type:     types.NumberType,
															Optional: true,
														},
														"schedule_end_time": {
															// Property: ScheduleEndTime
															Type:     types.NumberType,
															Optional: true,
														},
														"schedule_expression": {
															// Property: ScheduleExpression
															Type:     types.StringType,
															Required: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 256),
															},
														},
														"schedule_offset": {
															// Property: ScheduleOffset
															Type:     types.NumberType,
															Optional: true,
															Validators: []tfsdk.AttributeValidator{
																validate.IntBetween(0, 36000),
															},
														},
														"schedule_start_time": {
															// Property: ScheduleStartTime
															Type:     types.NumberType,
															Optional: true,
														},
														"timezone": {
															// Property: Timezone
															Type:     types.StringType,
															Optional: true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringLenBetween(0, 256),
															},
														},
													},
												),
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"trigger_type": {
									// Property: TriggerType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"Scheduled",
											"Event",
											"OnDemand",
										}),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			// FlowDefinition is a write-only attribute.
		},
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got last updated at",
			//   "type": "string"
			// }
			Description: "The time of this integration got last updated at",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_type_name": {
			// Property: ObjectTypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the ObjectType defined for the 3rd party data in Profile Service",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the ObjectType defined for the 3rd party data in Profile Service",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the integration",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the integration",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"uri": {
			// Property: Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URI of the S3 bucket or any other type of data source.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The URI of the S3 bucket or any other type of data source.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			// Uri is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The resource schema for creating an Amazon Connect Customer Profiles Integration.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Integration").WithTerraformTypeName("awscc_customerprofiles_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":                 "BucketName",
		"bucket_prefix":               "BucketPrefix",
		"connector_operator":          "ConnectorOperator",
		"connector_profile_name":      "ConnectorProfileName",
		"connector_type":              "ConnectorType",
		"created_at":                  "CreatedAt",
		"data_pull_mode":              "DataPullMode",
		"datetime_type_field_name":    "DatetimeTypeFieldName",
		"description":                 "Description",
		"destination_field":           "DestinationField",
		"domain_name":                 "DomainName",
		"enable_dynamic_field_update": "EnableDynamicFieldUpdate",
		"first_execution_from":        "FirstExecutionFrom",
		"flow_definition":             "FlowDefinition",
		"flow_name":                   "FlowName",
		"include_deleted_records":     "IncludeDeletedRecords",
		"incremental_pull_config":     "IncrementalPullConfig",
		"key":                         "Key",
		"kms_arn":                     "KmsArn",
		"last_updated_at":             "LastUpdatedAt",
		"marketo":                     "Marketo",
		"object":                      "Object",
		"object_type_name":            "ObjectTypeName",
		"operator_property_key":       "OperatorPropertyKey",
		"property":                    "Property",
		"s3":                          "S3",
		"salesforce":                  "Salesforce",
		"schedule_end_time":           "ScheduleEndTime",
		"schedule_expression":         "ScheduleExpression",
		"schedule_offset":             "ScheduleOffset",
		"schedule_start_time":         "ScheduleStartTime",
		"scheduled":                   "Scheduled",
		"service_now":                 "ServiceNow",
		"source_connector_properties": "SourceConnectorProperties",
		"source_fields":               "SourceFields",
		"source_flow_config":          "SourceFlowConfig",
		"tags":                        "Tags",
		"task_properties":             "TaskProperties",
		"task_type":                   "TaskType",
		"tasks":                       "Tasks",
		"timezone":                    "Timezone",
		"trigger_config":              "TriggerConfig",
		"trigger_properties":          "TriggerProperties",
		"trigger_type":                "TriggerType",
		"uri":                         "Uri",
		"value":                       "Value",
		"zendesk":                     "Zendesk",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/FlowDefinition",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_customerprofiles_integration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
