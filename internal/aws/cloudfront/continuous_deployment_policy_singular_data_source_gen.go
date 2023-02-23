// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_continuous_deployment_policy", continuousDeploymentPolicyDataSource)
}

// continuousDeploymentPolicyDataSource returns the Terraform awscc_cloudfront_continuous_deployment_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::ContinuousDeploymentPolicy resource.
func continuousDeploymentPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContinuousDeploymentPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Enabled": {
		//	      "type": "boolean"
		//	    },
		//	    "StagingDistributionDnsNames": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TrafficConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SingleHeaderConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Header": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 1783,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Header",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "SingleWeightConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "SessionStickinessConfig": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "IdleTTL": {
		//	                  "maximum": 3600,
		//	                  "minimum": 300,
		//	                  "type": "integer"
		//	                },
		//	                "MaximumTTL": {
		//	                  "maximum": 3600,
		//	                  "minimum": 300,
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "IdleTTL",
		//	                "MaximumTTL"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "Weight": {
		//	              "maximum": 1,
		//	              "minimum": 0,
		//	              "type": "number"
		//	            }
		//	          },
		//	          "required": [
		//	            "Weight"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Type": {
		//	          "enum": [
		//	            "SingleWeight",
		//	            "SingleHeader"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Type"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Enabled",
		//	    "StagingDistributionDnsNames"
		//	  ],
		//	  "type": "object"
		//	}
		"continuous_deployment_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: StagingDistributionDnsNames
				"staging_distribution_dns_names": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TrafficConfig
				"traffic_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SingleHeaderConfig
						"single_header_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Header
								"header": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SingleWeightConfig
						"single_weight_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: SessionStickinessConfig
								"session_stickiness_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IdleTTL
										"idle_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: MaximumTTL
										"maximum_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Weight
								"weight": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Type
						"type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::ContinuousDeploymentPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::ContinuousDeploymentPolicy").WithTerraformTypeName("awscc_cloudfront_continuous_deployment_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"continuous_deployment_policy_config": "ContinuousDeploymentPolicyConfig",
		"enabled":                             "Enabled",
		"header":                              "Header",
		"id":                                  "Id",
		"idle_ttl":                            "IdleTTL",
		"last_modified_time":                  "LastModifiedTime",
		"maximum_ttl":                         "MaximumTTL",
		"session_stickiness_config":           "SessionStickinessConfig",
		"single_header_config":                "SingleHeaderConfig",
		"single_weight_config":                "SingleWeightConfig",
		"staging_distribution_dns_names":      "StagingDistributionDnsNames",
		"traffic_config":                      "TrafficConfig",
		"type":                                "Type",
		"value":                               "Value",
		"weight":                              "Weight",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}