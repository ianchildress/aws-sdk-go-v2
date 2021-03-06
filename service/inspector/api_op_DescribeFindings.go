// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeFindingsRequest
type DescribeFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the finding that you want to describe.
	//
	// FindingArns is a required field
	FindingArns []string `locationName:"findingArns" min:"1" type:"list" required:"true"`

	// The locale into which you want to translate a finding description, recommendation,
	// and the short description that identifies the finding.
	Locale Locale `locationName:"locale" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFindingsInput"}

	if s.FindingArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingArns"))
	}
	if s.FindingArns != nil && len(s.FindingArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FindingArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeFindingsResponse
type DescribeFindingsOutput struct {
	_ struct{} `type:"structure"`

	// Finding details that cannot be described. An error code is provided for each
	// failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`

	// Information about the finding.
	//
	// Findings is a required field
	Findings []Finding `locationName:"findings" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFindings = "DescribeFindings"

// DescribeFindingsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Describes the findings that are specified by the ARNs of the findings.
//
//    // Example sending a request using DescribeFindingsRequest.
//    req := client.DescribeFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeFindings
func (c *Client) DescribeFindingsRequest(input *DescribeFindingsInput) DescribeFindingsRequest {
	op := &aws.Operation{
		Name:       opDescribeFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFindingsInput{}
	}

	req := c.newRequest(op, input, &DescribeFindingsOutput{})
	return DescribeFindingsRequest{Request: req, Input: input, Copy: c.DescribeFindingsRequest}
}

// DescribeFindingsRequest is the request type for the
// DescribeFindings API operation.
type DescribeFindingsRequest struct {
	*aws.Request
	Input *DescribeFindingsInput
	Copy  func(*DescribeFindingsInput) DescribeFindingsRequest
}

// Send marshals and sends the DescribeFindings API request.
func (r DescribeFindingsRequest) Send(ctx context.Context) (*DescribeFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFindingsResponse{
		DescribeFindingsOutput: r.Request.Data.(*DescribeFindingsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFindingsResponse is the response type for the
// DescribeFindings API operation.
type DescribeFindingsResponse struct {
	*DescribeFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFindings request.
func (r *DescribeFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
