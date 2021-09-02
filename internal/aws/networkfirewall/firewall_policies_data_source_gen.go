// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_networkfirewall_firewall_policies", firewallPoliciesDataSourceType)
}

// firewallPoliciesDataSourceType returns the Terraform awscc_networkfirewall_firewall_policies data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NetworkFirewall::FirewallPolicy resource type.
func firewallPoliciesDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	// Required for acceptance testing.
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::NetworkFirewall::FirewallPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.FromCloudFormationAndTerraform("AWS::NetworkFirewall::FirewallPolicy", "awscc_networkfirewall_firewall_policies", schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkfirewall_firewall_policies", "schema", hclog.Fmt("%v", schema))

	return pluralDataSourceType, nil
}
