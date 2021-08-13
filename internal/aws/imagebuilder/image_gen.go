// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_imagebuilder_image", imageResourceType)
}

// imageResourceType returns the Terraform aws_imagebuilder_image resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::Image resource type.
func imageResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"container_recipe_arn": {
			// Property: ContainerRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			Type:        types.StringType,
			Optional:    true,
		},
		"distribution_configuration_arn": {
			// Property: DistributionConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the distribution configuration.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// DistributionConfigurationArn is a force-new attribute.
		},
		"enhanced_image_metadata_enabled": {
			// Property: EnhancedImageMetadataEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			//   "type": "boolean"
			// }
			Description: "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			// EnhancedImageMetadataEnabled is a force-new attribute.
		},
		"image_id": {
			// Property: ImageId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AMI ID of the EC2 AMI in current region.",
			//   "type": "string"
			// }
			Description: "The AMI ID of the EC2 AMI in current region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_recipe_arn": {
			// Property: ImageRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ImageRecipeArn is a force-new attribute.
		},
		"image_tests_configuration": {
			// Property: ImageTestsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "The image tests configuration used when creating this image.",
			//   "properties": {
			//     "ImageTestsEnabled": {
			//       "description": "ImageTestsEnabled",
			//       "type": "boolean"
			//     },
			//     "TimeoutMinutes": {
			//       "description": "TimeoutMinutes",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The image tests configuration used when creating this image.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"image_tests_enabled": {
						// Property: ImageTestsEnabled
						Description: "ImageTestsEnabled",
						Type:        types.BoolType,
						Optional:    true,
					},
					"timeout_minutes": {
						// Property: TimeoutMinutes
						Description: "TimeoutMinutes",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// ImageTestsConfiguration is a force-new attribute.
		},
		"infrastructure_configuration_arn": {
			// Property: InfrastructureConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// InfrastructureConfigurationArn is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image.",
			//   "type": "string"
			// }
			Description: "The name of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags associated with the image.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags associated with the image.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::Image",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::Image").WithTerraformTypeName("aws_imagebuilder_image").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(720).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_imagebuilder_image", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}