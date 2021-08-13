// Code generated by generators/resource/main.go; DO NOT EDIT.

package autoscaling_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSAutoScalingWarmPool_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AutoScaling::WarmPool", "aws_autoscaling_warm_pool", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}