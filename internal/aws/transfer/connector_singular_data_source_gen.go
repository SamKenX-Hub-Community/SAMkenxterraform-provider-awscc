// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_transfer_connector", connectorDataSource)
}

// connectorDataSource returns the Terraform awscc_transfer_connector data source.
// This Terraform data source corresponds to the CloudFormation AWS::Transfer::Connector resource.
func connectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the access role for the connector.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:.*role/.*",
		//	  "type": "string"
		//	}
		"access_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the access role for the connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
		//	  "maxLength": 1600,
		//	  "minLength": 20,
		//	  "pattern": "arn:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: As2Config
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for an AS2 connector.",
		//	  "properties": {
		//	    "Compression": {
		//	      "description": "Compression setting for this AS2 connector configuration.",
		//	      "enum": [
		//	        "ZLIB",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "EncryptionAlgorithm": {
		//	      "description": "Encryption algorithm for this AS2 connector configuration.",
		//	      "enum": [
		//	        "AES128_CBC",
		//	        "AES192_CBC",
		//	        "AES256_CBC",
		//	        "NONE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "LocalProfileId": {
		//	      "description": "A unique identifier for the local profile.",
		//	      "maxLength": 19,
		//	      "minLength": 19,
		//	      "pattern": "^p-([0-9a-f]{17})$",
		//	      "type": "string"
		//	    },
		//	    "MdnResponse": {
		//	      "description": "MDN Response setting for this AS2 connector configuration.",
		//	      "enum": [
		//	        "SYNC",
		//	        "NONE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "MdnSigningAlgorithm": {
		//	      "description": "MDN Signing algorithm for this AS2 connector configuration.",
		//	      "enum": [
		//	        "SHA256",
		//	        "SHA384",
		//	        "SHA512",
		//	        "SHA1",
		//	        "NONE",
		//	        "DEFAULT"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "MessageSubject": {
		//	      "description": "The message subject for this AS2 connector configuration.",
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "PartnerProfileId": {
		//	      "description": "A unique identifier for the partner profile.",
		//	      "maxLength": 19,
		//	      "minLength": 19,
		//	      "pattern": "^p-([0-9a-f]{17})$",
		//	      "type": "string"
		//	    },
		//	    "SigningAlgorithm": {
		//	      "description": "Signing algorithm for this AS2 connector configuration.",
		//	      "enum": [
		//	        "SHA256",
		//	        "SHA384",
		//	        "SHA512",
		//	        "SHA1",
		//	        "NONE"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"as_2_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Compression
				"compression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Compression setting for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EncryptionAlgorithm
				"encryption_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Encryption algorithm for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: LocalProfileId
				"local_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A unique identifier for the local profile.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MdnResponse
				"mdn_response": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "MDN Response setting for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MdnSigningAlgorithm
				"mdn_signing_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "MDN Signing algorithm for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MessageSubject
				"message_subject": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The message subject for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PartnerProfileId
				"partner_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A unique identifier for the partner profile.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SigningAlgorithm
				"signing_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Signing algorithm for this AS2 connector configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for an AS2 connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectorId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the connector.",
		//	  "maxLength": 19,
		//	  "minLength": 19,
		//	  "pattern": "^c-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"connector_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the logging role for the connector.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:.*role/.*",
		//	  "type": "string"
		//	}
		"logging_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the logging role for the connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Creates a key-value pair for a specific resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The name assigned to the tag that you create.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Contains one or more values that you assigned to the key name you create.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name assigned to the tag that you create.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Contains one or more values that you assigned to the key name you create.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Url
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "URL for Connector",
		//	  "maxLength": 255,
		//	  "type": "string"
		//	}
		"url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "URL for Connector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Transfer::Connector",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Connector").WithTerraformTypeName("awscc_transfer_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role":           "AccessRole",
		"arn":                   "Arn",
		"as_2_config":           "As2Config",
		"compression":           "Compression",
		"connector_id":          "ConnectorId",
		"encryption_algorithm":  "EncryptionAlgorithm",
		"key":                   "Key",
		"local_profile_id":      "LocalProfileId",
		"logging_role":          "LoggingRole",
		"mdn_response":          "MdnResponse",
		"mdn_signing_algorithm": "MdnSigningAlgorithm",
		"message_subject":       "MessageSubject",
		"partner_profile_id":    "PartnerProfileId",
		"signing_algorithm":     "SigningAlgorithm",
		"tags":                  "Tags",
		"url":                   "Url",
		"value":                 "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
