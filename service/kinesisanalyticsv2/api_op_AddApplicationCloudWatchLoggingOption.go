// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationCloudWatchLoggingOptionRequest
type AddApplicationCloudWatchLoggingOptionInput struct {
	_ struct{} `type:"structure"`

	// The Kinesis Data Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Provides the Amazon CloudWatch log stream Amazon Resource Name (ARN).
	//
	// CloudWatchLoggingOption is a required field
	CloudWatchLoggingOption *CloudWatchLoggingOption `type:"structure" required:"true"`

	// The version ID of the Kinesis Data Analytics application. You can retrieve
	// the application version ID using DescribeApplication.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s AddApplicationCloudWatchLoggingOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationCloudWatchLoggingOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationCloudWatchLoggingOptionInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CloudWatchLoggingOption == nil {
		invalidParams.Add(aws.NewErrParamRequired("CloudWatchLoggingOption"))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}
	if s.CloudWatchLoggingOption != nil {
		if err := s.CloudWatchLoggingOption.Validate(); err != nil {
			invalidParams.AddNested("CloudWatchLoggingOption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationCloudWatchLoggingOptionResponse
type AddApplicationCloudWatchLoggingOptionOutput struct {
	_ struct{} `type:"structure"`

	// The application's ARN.
	ApplicationARN *string `min:"1" type:"string"`

	// The new version ID of the Kinesis Data Analytics application. Kinesis Data
	// Analytics updates the ApplicationVersionId each time you change the CloudWatch
	// logging options.
	ApplicationVersionId *int64 `min:"1" type:"long"`

	// The descriptions of the current CloudWatch logging options for the Kinesis
	// Data Analytics application.
	CloudWatchLoggingOptionDescriptions []CloudWatchLoggingOptionDescription `type:"list"`
}

// String returns the string representation
func (s AddApplicationCloudWatchLoggingOptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationCloudWatchLoggingOption = "AddApplicationCloudWatchLoggingOption"

// AddApplicationCloudWatchLoggingOptionRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds an Amazon CloudWatch log stream to monitor application configuration
// errors.
//
//    // Example sending a request using AddApplicationCloudWatchLoggingOptionRequest.
//    req := client.AddApplicationCloudWatchLoggingOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationCloudWatchLoggingOption
func (c *Client) AddApplicationCloudWatchLoggingOptionRequest(input *AddApplicationCloudWatchLoggingOptionInput) AddApplicationCloudWatchLoggingOptionRequest {
	op := &aws.Operation{
		Name:       opAddApplicationCloudWatchLoggingOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationCloudWatchLoggingOptionInput{}
	}

	req := c.newRequest(op, input, &AddApplicationCloudWatchLoggingOptionOutput{})
	return AddApplicationCloudWatchLoggingOptionRequest{Request: req, Input: input, Copy: c.AddApplicationCloudWatchLoggingOptionRequest}
}

// AddApplicationCloudWatchLoggingOptionRequest is the request type for the
// AddApplicationCloudWatchLoggingOption API operation.
type AddApplicationCloudWatchLoggingOptionRequest struct {
	*aws.Request
	Input *AddApplicationCloudWatchLoggingOptionInput
	Copy  func(*AddApplicationCloudWatchLoggingOptionInput) AddApplicationCloudWatchLoggingOptionRequest
}

// Send marshals and sends the AddApplicationCloudWatchLoggingOption API request.
func (r AddApplicationCloudWatchLoggingOptionRequest) Send(ctx context.Context) (*AddApplicationCloudWatchLoggingOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationCloudWatchLoggingOptionResponse{
		AddApplicationCloudWatchLoggingOptionOutput: r.Request.Data.(*AddApplicationCloudWatchLoggingOptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationCloudWatchLoggingOptionResponse is the response type for the
// AddApplicationCloudWatchLoggingOption API operation.
type AddApplicationCloudWatchLoggingOptionResponse struct {
	*AddApplicationCloudWatchLoggingOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationCloudWatchLoggingOption request.
func (r *AddApplicationCloudWatchLoggingOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
