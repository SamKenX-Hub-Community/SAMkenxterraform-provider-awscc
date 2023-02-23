// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigatewayv2_api", apiDataSource)
}

// apiDataSource returns the Terraform awscc_apigatewayv2_api data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGatewayV2::Api resource.
func apiDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ApiKeySelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_key_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BasePath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"base_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Body
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"body": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BodyS3Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Bucket": {
		//	      "type": "string"
		//	    },
		//	    "Etag": {
		//	      "type": "string"
		//	    },
		//	    "Key": {
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"body_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Etag
				"etag": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Key
				"key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CorsConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AllowCredentials": {
		//	      "type": "boolean"
		//	    },
		//	    "AllowHeaders": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "AllowMethods": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "AllowOrigins": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "ExposeHeaders": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "MaxAge": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cors_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCredentials
				"allow_credentials": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AllowHeaders
				"allow_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowMethods
				"allow_methods": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowOrigins
				"allow_origins": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExposeHeaders
				"expose_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxAge
				"max_age": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CredentialsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"credentials_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisableExecuteApiEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"disable_execute_api_endpoint": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisableSchemaValidation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"disable_schema_validation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FailOnWarnings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"fail_on_warnings": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProtocolType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"protocol_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RouteKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"route_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RouteSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"route_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "This resource type use map for Tags, suggest to use List of Tag",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "This resource type use map for Tags, suggest to use List of Tag",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Target
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"target": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGatewayV2::Api",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Api").WithTerraformTypeName("awscc_apigatewayv2_api")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_credentials":            "AllowCredentials",
		"allow_headers":                "AllowHeaders",
		"allow_methods":                "AllowMethods",
		"allow_origins":                "AllowOrigins",
		"api_endpoint":                 "ApiEndpoint",
		"api_id":                       "ApiId",
		"api_key_selection_expression": "ApiKeySelectionExpression",
		"base_path":                    "BasePath",
		"body":                         "Body",
		"body_s3_location":             "BodyS3Location",
		"bucket":                       "Bucket",
		"cors_configuration":           "CorsConfiguration",
		"credentials_arn":              "CredentialsArn",
		"description":                  "Description",
		"disable_execute_api_endpoint": "DisableExecuteApiEndpoint",
		"disable_schema_validation":    "DisableSchemaValidation",
		"etag":                         "Etag",
		"expose_headers":               "ExposeHeaders",
		"fail_on_warnings":             "FailOnWarnings",
		"key":                          "Key",
		"max_age":                      "MaxAge",
		"name":                         "Name",
		"protocol_type":                "ProtocolType",
		"route_key":                    "RouteKey",
		"route_selection_expression":   "RouteSelectionExpression",
		"tags":                         "Tags",
		"target":                       "Target",
		"version":                      "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}