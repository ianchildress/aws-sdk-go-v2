// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/StopFleetRequest
type StopFleetInput struct {
	_ struct{} `type:"structure"`

	// The name of the fleet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopFleetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/StopFleetResult
type StopFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopFleet = "StopFleet"

// StopFleetRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Stops the specified fleet.
//
//    // Example sending a request using StopFleetRequest.
//    req := client.StopFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/StopFleet
func (c *Client) StopFleetRequest(input *StopFleetInput) StopFleetRequest {
	op := &aws.Operation{
		Name:       opStopFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopFleetInput{}
	}

	req := c.newRequest(op, input, &StopFleetOutput{})
	return StopFleetRequest{Request: req, Input: input, Copy: c.StopFleetRequest}
}

// StopFleetRequest is the request type for the
// StopFleet API operation.
type StopFleetRequest struct {
	*aws.Request
	Input *StopFleetInput
	Copy  func(*StopFleetInput) StopFleetRequest
}

// Send marshals and sends the StopFleet API request.
func (r StopFleetRequest) Send(ctx context.Context) (*StopFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopFleetResponse{
		StopFleetOutput: r.Request.Data.(*StopFleetOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopFleetResponse is the response type for the
// StopFleet API operation.
type StopFleetResponse struct {
	*StopFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopFleet request.
func (r *StopFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
