// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateGatewayGroupRequest
type UpdateGatewayGroupInput struct {
	_ struct{} `type:"structure"`

	// The updated description of the gateway group.
	Description *string `type:"string"`

	// The ARN of the gateway group to update.
	//
	// GatewayGroupArn is a required field
	GatewayGroupArn *string `type:"string" required:"true"`

	// The updated name of the gateway group.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateGatewayGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewayGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGatewayGroupInput"}

	if s.GatewayGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayGroupArn"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateGatewayGroupResponse
type UpdateGatewayGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateGatewayGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateGatewayGroup = "UpdateGatewayGroup"

// UpdateGatewayGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates the details of a gateway group. If any optional field is not provided,
// the existing corresponding value is left unmodified.
//
//    // Example sending a request using UpdateGatewayGroupRequest.
//    req := client.UpdateGatewayGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateGatewayGroup
func (c *Client) UpdateGatewayGroupRequest(input *UpdateGatewayGroupInput) UpdateGatewayGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateGatewayGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGatewayGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateGatewayGroupOutput{})
	return UpdateGatewayGroupRequest{Request: req, Input: input, Copy: c.UpdateGatewayGroupRequest}
}

// UpdateGatewayGroupRequest is the request type for the
// UpdateGatewayGroup API operation.
type UpdateGatewayGroupRequest struct {
	*aws.Request
	Input *UpdateGatewayGroupInput
	Copy  func(*UpdateGatewayGroupInput) UpdateGatewayGroupRequest
}

// Send marshals and sends the UpdateGatewayGroup API request.
func (r UpdateGatewayGroupRequest) Send(ctx context.Context) (*UpdateGatewayGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGatewayGroupResponse{
		UpdateGatewayGroupOutput: r.Request.Data.(*UpdateGatewayGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGatewayGroupResponse is the response type for the
// UpdateGatewayGroup API operation.
type UpdateGatewayGroupResponse struct {
	*UpdateGatewayGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGatewayGroup request.
func (r *UpdateGatewayGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
