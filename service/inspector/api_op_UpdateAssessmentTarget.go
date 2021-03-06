// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/UpdateAssessmentTargetRequest
type UpdateAssessmentTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment target that you want to update.
	//
	// AssessmentTargetArn is a required field
	AssessmentTargetArn *string `locationName:"assessmentTargetArn" min:"1" type:"string" required:"true"`

	// The name of the assessment target that you want to update.
	//
	// AssessmentTargetName is a required field
	AssessmentTargetName *string `locationName:"assessmentTargetName" min:"1" type:"string" required:"true"`

	// The ARN of the resource group that is used to specify the new resource group
	// to associate with the assessment target.
	ResourceGroupArn *string `locationName:"resourceGroupArn" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAssessmentTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAssessmentTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAssessmentTargetInput"}

	if s.AssessmentTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetArn"))
	}
	if s.AssessmentTargetArn != nil && len(*s.AssessmentTargetArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetArn", 1))
	}

	if s.AssessmentTargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetName"))
	}
	if s.AssessmentTargetName != nil && len(*s.AssessmentTargetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetName", 1))
	}
	if s.ResourceGroupArn != nil && len(*s.ResourceGroupArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/UpdateAssessmentTargetOutput
type UpdateAssessmentTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAssessmentTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAssessmentTarget = "UpdateAssessmentTarget"

// UpdateAssessmentTargetRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Updates the assessment target that is specified by the ARN of the assessment
// target.
//
// If resourceGroupArn is not specified, all EC2 instances in the current AWS
// account and region are included in the assessment target.
//
//    // Example sending a request using UpdateAssessmentTargetRequest.
//    req := client.UpdateAssessmentTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/UpdateAssessmentTarget
func (c *Client) UpdateAssessmentTargetRequest(input *UpdateAssessmentTargetInput) UpdateAssessmentTargetRequest {
	op := &aws.Operation{
		Name:       opUpdateAssessmentTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAssessmentTargetInput{}
	}

	req := c.newRequest(op, input, &UpdateAssessmentTargetOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAssessmentTargetRequest{Request: req, Input: input, Copy: c.UpdateAssessmentTargetRequest}
}

// UpdateAssessmentTargetRequest is the request type for the
// UpdateAssessmentTarget API operation.
type UpdateAssessmentTargetRequest struct {
	*aws.Request
	Input *UpdateAssessmentTargetInput
	Copy  func(*UpdateAssessmentTargetInput) UpdateAssessmentTargetRequest
}

// Send marshals and sends the UpdateAssessmentTarget API request.
func (r UpdateAssessmentTargetRequest) Send(ctx context.Context) (*UpdateAssessmentTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAssessmentTargetResponse{
		UpdateAssessmentTargetOutput: r.Request.Data.(*UpdateAssessmentTargetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAssessmentTargetResponse is the response type for the
// UpdateAssessmentTarget API operation.
type UpdateAssessmentTargetResponse struct {
	*UpdateAssessmentTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAssessmentTarget request.
func (r *UpdateAssessmentTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
