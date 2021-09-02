// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package customerprofiles

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
	registry.AddDataSourceTypeFactory("awscc_customerprofiles_domains", domainsDataSourceType)
}

// domainsDataSourceType returns the Terraform awscc_customerprofiles_domains data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CustomerProfiles::Domain resource type.
func domainsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
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
		Description: "Plural Data Source schema for AWS::CustomerProfiles::Domain",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.FromCloudFormationAndTerraform("AWS::CustomerProfiles::Domain", "awscc_customerprofiles_domains", schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_customerprofiles_domains", "schema", hclog.Fmt("%v", schema))

	return pluralDataSourceType, nil
}
