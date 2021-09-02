// Code generated by generators/resource/main.go; DO NOT EDIT.

package chatbot

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
	registry.AddResourceTypeFactory("awscc_chatbot_slack_channel_configuration", slackChannelConfigurationResourceType)
}

// slackChannelConfigurationResourceType returns the Terraform awscc_chatbot_slack_channel_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Chatbot::SlackChannelConfiguration resource type.
func slackChannelConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the configuration",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the configuration",
			Type:        types.StringType,
			Computed:    true,
		},
		"configuration_name": {
			// Property: ConfigurationName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the configuration",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the configuration",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			// ConfigurationName is a force-new attribute.
		},
		"iam_role_arn": {
			// Property: IamRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the IAM role that defines the permissions for AWS Chatbot",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the IAM role that defines the permissions for AWS Chatbot",
			Type:        types.StringType,
			Required:    true,
		},
		"logging_level": {
			// Property: LoggingLevel
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
			Type:        types.StringType,
			Optional:    true,
		},
		"slack_channel_id": {
			// Property: SlackChannelId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the Slack channel",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The id of the Slack channel",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"slack_workspace_id": {
			// Property: SlackWorkspaceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the Slack workspace",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The id of the Slack workspace",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			// SlackWorkspaceId is a force-new attribute.
		},
		"sns_topic_arns": {
			// Property: SnsTopicArns
			// CloudFormation resource type schema:
			// {
			//   "description": "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
			//   "items": {
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Chatbot::SlackChannelConfiguration.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Chatbot::SlackChannelConfiguration").WithTerraformTypeName("awscc_chatbot_slack_channel_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"configuration_name": "ConfigurationName",
		"iam_role_arn":       "IamRoleArn",
		"logging_level":      "LoggingLevel",
		"slack_channel_id":   "SlackChannelId",
		"slack_workspace_id": "SlackWorkspaceId",
		"sns_topic_arns":     "SnsTopicArns",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_chatbot_slack_channel_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
