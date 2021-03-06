// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a DisableStageTransition action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DisableStageTransitionInput
type DisableStageTransitionInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline in which you want to disable the flow of artifacts
	// from one stage to another.
	//
	// PipelineName is a required field
	PipelineName *string `locationName:"pipelineName" min:"1" type:"string" required:"true"`

	// The reason given to the user why a stage is disabled, such as waiting for
	// manual approval or manual tests. This message is displayed in the pipeline
	// console UI.
	//
	// Reason is a required field
	Reason *string `locationName:"reason" min:"1" type:"string" required:"true"`

	// The name of the stage where you want to disable the inbound or outbound transition
	// of artifacts.
	//
	// StageName is a required field
	StageName *string `locationName:"stageName" min:"1" type:"string" required:"true"`

	// Specifies whether artifacts will be prevented from transitioning into the
	// stage and being processed by the actions in that stage (inbound), or prevented
	// from transitioning from the stage after they have been processed by the actions
	// in that stage (outbound).
	//
	// TransitionType is a required field
	TransitionType StageTransitionType `locationName:"transitionType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DisableStageTransitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableStageTransitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableStageTransitionInput"}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}

	if s.Reason == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reason"))
	}
	if s.Reason != nil && len(*s.Reason) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Reason", 1))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}
	if s.StageName != nil && len(*s.StageName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StageName", 1))
	}
	if len(s.TransitionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TransitionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DisableStageTransitionOutput
type DisableStageTransitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableStageTransitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableStageTransition = "DisableStageTransition"

// DisableStageTransitionRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Prevents artifacts in a pipeline from transitioning to the next stage in
// the pipeline.
//
//    // Example sending a request using DisableStageTransitionRequest.
//    req := client.DisableStageTransitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DisableStageTransition
func (c *Client) DisableStageTransitionRequest(input *DisableStageTransitionInput) DisableStageTransitionRequest {
	op := &aws.Operation{
		Name:       opDisableStageTransition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableStageTransitionInput{}
	}

	req := c.newRequest(op, input, &DisableStageTransitionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DisableStageTransitionRequest{Request: req, Input: input, Copy: c.DisableStageTransitionRequest}
}

// DisableStageTransitionRequest is the request type for the
// DisableStageTransition API operation.
type DisableStageTransitionRequest struct {
	*aws.Request
	Input *DisableStageTransitionInput
	Copy  func(*DisableStageTransitionInput) DisableStageTransitionRequest
}

// Send marshals and sends the DisableStageTransition API request.
func (r DisableStageTransitionRequest) Send(ctx context.Context) (*DisableStageTransitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableStageTransitionResponse{
		DisableStageTransitionOutput: r.Request.Data.(*DisableStageTransitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableStageTransitionResponse is the response type for the
// DisableStageTransition API operation.
type DisableStageTransitionResponse struct {
	*DisableStageTransitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableStageTransition request.
func (r *DisableStageTransitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
