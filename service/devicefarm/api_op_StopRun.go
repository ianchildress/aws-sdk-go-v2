// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to stop a specific run.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRunRequest
type StopRunInput struct {
	_ struct{} `type:"structure"`

	// Represents the Amazon Resource Name (ARN) of the Device Farm run you wish
	// to stop.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s StopRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopRunInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the results of your stop run attempt.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRunResult
type StopRunOutput struct {
	_ struct{} `type:"structure"`

	// The run that was stopped.
	Run *Run `locationName:"run" type:"structure"`
}

// String returns the string representation
func (s StopRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopRun = "StopRun"

// StopRunRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Initiates a stop request for the current test run. AWS Device Farm will immediately
// stop the run on devices where tests have not started executing, and you will
// not be billed for these devices. On devices where tests have started executing,
// Setup Suite and Teardown Suite tests will run to completion before stopping
// execution on those devices. You will be billed for Setup, Teardown, and any
// tests that were in progress or already completed.
//
//    // Example sending a request using StopRunRequest.
//    req := client.StopRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRun
func (c *Client) StopRunRequest(input *StopRunInput) StopRunRequest {
	op := &aws.Operation{
		Name:       opStopRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopRunInput{}
	}

	req := c.newRequest(op, input, &StopRunOutput{})
	return StopRunRequest{Request: req, Input: input, Copy: c.StopRunRequest}
}

// StopRunRequest is the request type for the
// StopRun API operation.
type StopRunRequest struct {
	*aws.Request
	Input *StopRunInput
	Copy  func(*StopRunInput) StopRunRequest
}

// Send marshals and sends the StopRun API request.
func (r StopRunRequest) Send(ctx context.Context) (*StopRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopRunResponse{
		StopRunOutput: r.Request.Data.(*StopRunOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopRunResponse is the response type for the
// StopRun API operation.
type StopRunResponse struct {
	*StopRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopRun request.
func (r *StopRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
