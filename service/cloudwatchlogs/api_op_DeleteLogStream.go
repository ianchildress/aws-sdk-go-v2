// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteLogStreamRequest
type DeleteLogStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The name of the log stream.
	//
	// LogStreamName is a required field
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLogStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLogStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLogStreamInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.LogStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogStreamName"))
	}
	if s.LogStreamName != nil && len(*s.LogStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteLogStreamOutput
type DeleteLogStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLogStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLogStream = "DeleteLogStream"

// DeleteLogStreamRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes the specified log stream and permanently deletes all the archived
// log events associated with the log stream.
//
//    // Example sending a request using DeleteLogStreamRequest.
//    req := client.DeleteLogStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteLogStream
func (c *Client) DeleteLogStreamRequest(input *DeleteLogStreamInput) DeleteLogStreamRequest {
	op := &aws.Operation{
		Name:       opDeleteLogStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLogStreamInput{}
	}

	req := c.newRequest(op, input, &DeleteLogStreamOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteLogStreamRequest{Request: req, Input: input, Copy: c.DeleteLogStreamRequest}
}

// DeleteLogStreamRequest is the request type for the
// DeleteLogStream API operation.
type DeleteLogStreamRequest struct {
	*aws.Request
	Input *DeleteLogStreamInput
	Copy  func(*DeleteLogStreamInput) DeleteLogStreamRequest
}

// Send marshals and sends the DeleteLogStream API request.
func (r DeleteLogStreamRequest) Send(ctx context.Context) (*DeleteLogStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLogStreamResponse{
		DeleteLogStreamOutput: r.Request.Data.(*DeleteLogStreamOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLogStreamResponse is the response type for the
// DeleteLogStream API operation.
type DeleteLogStreamResponse struct {
	*DeleteLogStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLogStream request.
func (r *DeleteLogStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
