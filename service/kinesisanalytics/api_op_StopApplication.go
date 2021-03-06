// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/StopApplicationRequest
type StopApplicationInput struct {
	_ struct{} `type:"structure"`

	// Name of the running application to stop.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/StopApplicationResponse
type StopApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopApplication = "StopApplication"

// StopApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Stops the application from processing input data. You can stop an application
// only if it is in the running state. You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
// operation to find the application state. After the application is stopped,
// Amazon Kinesis Analytics stops reading data from the input, the application
// stops processing data, and there is no output written to the destination.
//
// This operation requires permissions to perform the kinesisanalytics:StopApplication
// action.
//
//    // Example sending a request using StopApplicationRequest.
//    req := client.StopApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/StopApplication
func (c *Client) StopApplicationRequest(input *StopApplicationInput) StopApplicationRequest {
	op := &aws.Operation{
		Name:       opStopApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopApplicationInput{}
	}

	req := c.newRequest(op, input, &StopApplicationOutput{})
	return StopApplicationRequest{Request: req, Input: input, Copy: c.StopApplicationRequest}
}

// StopApplicationRequest is the request type for the
// StopApplication API operation.
type StopApplicationRequest struct {
	*aws.Request
	Input *StopApplicationInput
	Copy  func(*StopApplicationInput) StopApplicationRequest
}

// Send marshals and sends the StopApplication API request.
func (r StopApplicationRequest) Send(ctx context.Context) (*StopApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopApplicationResponse{
		StopApplicationOutput: r.Request.Data.(*StopApplicationOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopApplicationResponse is the response type for the
// StopApplication API operation.
type StopApplicationResponse struct {
	*StopApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopApplication request.
func (r *StopApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
