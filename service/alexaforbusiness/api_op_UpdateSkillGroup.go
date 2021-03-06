// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateSkillGroupRequest
type UpdateSkillGroupInput struct {
	_ struct{} `type:"structure"`

	// The updated description for the skill group.
	Description *string `min:"1" type:"string"`

	// The ARN of the skill group to update.
	SkillGroupArn *string `type:"string"`

	// The updated name for the skill group.
	SkillGroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSkillGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSkillGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSkillGroupInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.SkillGroupName != nil && len(*s.SkillGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SkillGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateSkillGroupResponse
type UpdateSkillGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSkillGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSkillGroup = "UpdateSkillGroup"

// UpdateSkillGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates skill group details by skill group ARN.
//
//    // Example sending a request using UpdateSkillGroupRequest.
//    req := client.UpdateSkillGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateSkillGroup
func (c *Client) UpdateSkillGroupRequest(input *UpdateSkillGroupInput) UpdateSkillGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateSkillGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSkillGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateSkillGroupOutput{})
	return UpdateSkillGroupRequest{Request: req, Input: input, Copy: c.UpdateSkillGroupRequest}
}

// UpdateSkillGroupRequest is the request type for the
// UpdateSkillGroup API operation.
type UpdateSkillGroupRequest struct {
	*aws.Request
	Input *UpdateSkillGroupInput
	Copy  func(*UpdateSkillGroupInput) UpdateSkillGroupRequest
}

// Send marshals and sends the UpdateSkillGroup API request.
func (r UpdateSkillGroupRequest) Send(ctx context.Context) (*UpdateSkillGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSkillGroupResponse{
		UpdateSkillGroupOutput: r.Request.Data.(*UpdateSkillGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSkillGroupResponse is the response type for the
// UpdateSkillGroup API operation.
type UpdateSkillGroupResponse struct {
	*UpdateSkillGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSkillGroup request.
func (r *UpdateSkillGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
