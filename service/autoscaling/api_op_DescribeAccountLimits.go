// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAccountLimitsInput
type DescribeAccountLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAccountLimitsAnswer
type DescribeAccountLimitsOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of groups allowed for your AWS account. The default limit
	// is 200 per region.
	MaxNumberOfAutoScalingGroups *int64 `type:"integer"`

	// The maximum number of launch configurations allowed for your AWS account.
	// The default limit is 200 per region.
	MaxNumberOfLaunchConfigurations *int64 `type:"integer"`

	// The current number of groups for your AWS account.
	NumberOfAutoScalingGroups *int64 `type:"integer"`

	// The current number of launch configurations for your AWS account.
	NumberOfLaunchConfigurations *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeAccountLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountLimits = "DescribeAccountLimits"

// DescribeAccountLimitsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the current Amazon EC2 Auto Scaling resource limits for your AWS
// account.
//
// For information about requesting an increase in these limits, see Amazon
// EC2 Auto Scaling Limits (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using DescribeAccountLimitsRequest.
//    req := client.DescribeAccountLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAccountLimits
func (c *Client) DescribeAccountLimitsRequest(input *DescribeAccountLimitsInput) DescribeAccountLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountLimitsOutput{})
	return DescribeAccountLimitsRequest{Request: req, Input: input, Copy: c.DescribeAccountLimitsRequest}
}

// DescribeAccountLimitsRequest is the request type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsRequest struct {
	*aws.Request
	Input *DescribeAccountLimitsInput
	Copy  func(*DescribeAccountLimitsInput) DescribeAccountLimitsRequest
}

// Send marshals and sends the DescribeAccountLimits API request.
func (r DescribeAccountLimitsRequest) Send(ctx context.Context) (*DescribeAccountLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountLimitsResponse{
		DescribeAccountLimitsOutput: r.Request.Data.(*DescribeAccountLimitsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountLimitsResponse is the response type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsResponse struct {
	*DescribeAccountLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountLimits request.
func (r *DescribeAccountLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
