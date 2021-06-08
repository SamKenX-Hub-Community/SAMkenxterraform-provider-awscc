package provider

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsbase "github.com/hashicorp/aws-sdk-go-base"

	tfsdk "github.com/hashicorp/terraform-plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func New(version string) tfsdk.Provider {
	return &awsProvider{}
}

type awsClient struct {
	cfconn *cloudformation.CloudFormation
}

type awsProvider struct {
	Client *awsClient
}

func (p *awsProvider) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"region": {
				Type:        types.StringType,
				Description: "The region where AWS operations will take place.",
				Required:    true,
			},

			"role_arn": {
				Type:        types.StringType,
				Description: "Amazon Resource Name of an IAM Role that is used to do the actual provisioning.",
				Optional:    true,
			},
		},
	}, nil
}

func (p *awsProvider) Configure(_ context.Context, input *tfsdk.ConfigureProviderRequest, output *tfsdk.ConfigureProviderResponse) {
	client, err := newAWSClient()

	if err != nil {
		//return nil, appendDiagnostic(nil, fmt.Errorf("error configuring Terraform AWS Provider: %w", err))
	}

	p.Client = client
}

func (p *awsProvider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	return resources, nil
}

func (p *awsProvider) GetDataSources(context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	return nil, nil
}

/*
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"region": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"AWS_REGION",
						"AWS_DEFAULT_REGION",
					}, nil),
					Description:  "The region where AWS operations will take place.",
					InputDefault: "us-east-1",
				},

				"role_arn": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Amazon Resource Name of an IAM Role that is used to do the actual provisioning.",
				},
			},

			ResourcesMap: resources,
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}
*/
/*
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		config := Config{
			Region: d.Get("region").(string),
		}

		return config.Client()
	}
}
*/

var resources map[string]tfsdk.ResourceType
var resourcesMu sync.Mutex

func registerResourceType(name string, r tfsdk.ResourceType) {
	resourcesMu.Lock()
	defer resourcesMu.Unlock()

	if resources == nil {
		resources = make(map[string]tfsdk.ResourceType)
	}
	resources[name] = r
}

// newAWSClient configures and returns a fully initialized AWS client.
func newAWSClient() (*awsClient, error) {
	awsbaseConfig := &awsbase.Config{
		//DebugLogging: logging.IsDebugOrHigher(),
		Region: "us-west-2",
	}

	sess, _, _, err := awsbase.GetSessionWithAccountIDAndPartition(awsbaseConfig)
	if err != nil {
		return nil, fmt.Errorf("error getting AWS SDK session: %w", err)
	}

	client := &awsClient{
		cfconn: cloudformation.New(sess.Copy(&aws.Config{})),
	}

	return client, nil
}

// appendDiagnostic appends an error or warning message to a response's diagnostics.
func appendDiagnostic(diags []*tfprotov6.Diagnostic, d interface{}) []*tfprotov6.Diagnostic {
	switch d := d.(type) {
	case error:
		diags = append(diags, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  d.Error(),
		})
	case string:
		diags = append(diags, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  d,
		})
	}

	return diags
}
