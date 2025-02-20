// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_security_key", securityKeyDataSource)
}

// securityKeyDataSource returns the Terraform awscc_connect_security_key data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::SecurityKey resource.
func securityKeyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An associationID is automatically generated when a storage config is associated with an instance",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[-a-z0-9]*$",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An associationID is automatically generated when a storage config is associated with an instance",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Connect instance identifier",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Connect instance identifier",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Key
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A valid security key in PEM format.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A valid security key in PEM format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::SecurityKey",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::SecurityKey").WithTerraformTypeName("awscc_connect_security_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"instance_id":    "InstanceId",
		"key":            "Key",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
