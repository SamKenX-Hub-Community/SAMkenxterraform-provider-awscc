// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudwatch_metric_stream", metricStreamDataSource)
}

// metricStreamDataSource returns the Terraform awscc_cloudwatch_metric_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudWatch::MetricStream resource.
func metricStreamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name of the metric stream.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name of the metric stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "anyOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "description": "The date of creation of the metric stream.",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date of creation of the metric stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExcludeFilters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "This structure defines the metrics that will be streamed.",
		//	    "properties": {
		//	      "Namespace": {
		//	        "description": "Only metrics with Namespace matching this value will be streamed.",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Namespace"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1000,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"exclude_filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Namespace
					"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FirehoseArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Kinesis Firehose where to stream the data.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"firehose_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Kinesis Firehose where to stream the data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IncludeFilters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "This structure defines the metrics that will be streamed.",
		//	    "properties": {
		//	      "Namespace": {
		//	        "description": "Only metrics with Namespace matching this value will be streamed.",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Namespace"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1000,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"include_filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Namespace
					"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IncludeLinkedAccountsMetrics
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false.",
		//	  "type": "boolean"
		//	}
		"include_linked_accounts_metrics": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdateDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "anyOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "description": "The date of the last update of the metric stream.",
		//	  "type": "string"
		//	}
		"last_update_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date of the last update of the metric stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the metric stream.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the metric stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutputFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The output format of the data streamed to the Kinesis Firehose.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"output_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The output format of the data streamed to the Kinesis Firehose.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the role that provides access to the Kinesis Firehose.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the role that provides access to the Kinesis Firehose.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Displays the state of the Metric Stream.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Displays the state of the Metric Stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatisticsConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "This structure specifies a list of additional statistics to stream, and the metrics to stream those additional statistics for. All metrics that match the combination of metric name and namespace will be streamed with the extended statistics, no matter their dimensions.",
		//	    "properties": {
		//	      "AdditionalStatistics": {
		//	        "description": "The additional statistics to stream for the metrics listed in IncludeMetrics.",
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "maxItems": 20,
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "IncludeMetrics": {
		//	        "description": "An array that defines the metrics that are to have additional statistics streamed.",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "A structure that specifies the metric name and namespace for one metric that is going to have additional statistics included in the stream.",
		//	          "properties": {
		//	            "MetricName": {
		//	              "description": "The name of the metric.",
		//	              "maxLength": 255,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Namespace": {
		//	              "description": "The namespace of the metric.",
		//	              "maxLength": 255,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "MetricName",
		//	            "Namespace"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 100,
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "AdditionalStatistics",
		//	      "IncludeMetrics"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 100,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"statistics_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AdditionalStatistics
					"additional_statistics": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The additional statistics to stream for the metrics listed in IncludeMetrics.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: IncludeMetrics
					"include_metrics": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MetricName
								"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the metric.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Namespace
								"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The namespace of the metric.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "An array that defines the metrics that are to have additional statistics streamed.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of tags to assign to the delivery stream.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata that you can assign to a Metric Stream, consisting of a key-value pair.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A unique identifier for the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "An optional string, which you can use to describe or define the tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique identifier for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "An optional string, which you can use to describe or define the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A set of tags to assign to the delivery stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudWatch::MetricStream",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::MetricStream").WithTerraformTypeName("awscc_cloudwatch_metric_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_statistics":           "AdditionalStatistics",
		"arn":                             "Arn",
		"creation_date":                   "CreationDate",
		"exclude_filters":                 "ExcludeFilters",
		"firehose_arn":                    "FirehoseArn",
		"include_filters":                 "IncludeFilters",
		"include_linked_accounts_metrics": "IncludeLinkedAccountsMetrics",
		"include_metrics":                 "IncludeMetrics",
		"key":                             "Key",
		"last_update_date":                "LastUpdateDate",
		"metric_name":                     "MetricName",
		"name":                            "Name",
		"namespace":                       "Namespace",
		"output_format":                   "OutputFormat",
		"role_arn":                        "RoleArn",
		"state":                           "State",
		"statistics_configurations":       "StatisticsConfigurations",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
