// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a DeletePipeline action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeletePipelineInput
type DeletePipelineInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline to be deleted.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePipelineInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeletePipelineOutput
type DeletePipelineOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePipeline = "DeletePipeline"

// DeletePipelineRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Deletes the specified pipeline.
//
//    // Example sending a request using DeletePipelineRequest.
//    req := client.DeletePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeletePipeline
func (c *Client) DeletePipelineRequest(input *DeletePipelineInput) DeletePipelineRequest {
	op := &aws.Operation{
		Name:       opDeletePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePipelineInput{}
	}

	req := c.newRequest(op, input, &DeletePipelineOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePipelineRequest{Request: req, Input: input, Copy: c.DeletePipelineRequest}
}

// DeletePipelineRequest is the request type for the
// DeletePipeline API operation.
type DeletePipelineRequest struct {
	*aws.Request
	Input *DeletePipelineInput
	Copy  func(*DeletePipelineInput) DeletePipelineRequest
}

// Send marshals and sends the DeletePipeline API request.
func (r DeletePipelineRequest) Send(ctx context.Context) (*DeletePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePipelineResponse{
		DeletePipelineOutput: r.Request.Data.(*DeletePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePipelineResponse is the response type for the
// DeletePipeline API operation.
type DeletePipelineResponse struct {
	*DeletePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePipeline request.
func (r *DeletePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
