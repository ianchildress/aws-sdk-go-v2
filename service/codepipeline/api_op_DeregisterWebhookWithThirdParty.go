// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeregisterWebhookWithThirdPartyInput
type DeregisterWebhookWithThirdPartyInput struct {
	_ struct{} `type:"structure"`

	// The name of the webhook you want to deregister.
	WebhookName *string `locationName:"webhookName" min:"1" type:"string"`
}

// String returns the string representation
func (s DeregisterWebhookWithThirdPartyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterWebhookWithThirdPartyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterWebhookWithThirdPartyInput"}
	if s.WebhookName != nil && len(*s.WebhookName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebhookName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeregisterWebhookWithThirdPartyOutput
type DeregisterWebhookWithThirdPartyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterWebhookWithThirdPartyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterWebhookWithThirdParty = "DeregisterWebhookWithThirdParty"

// DeregisterWebhookWithThirdPartyRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Removes the connection between the webhook that was created by CodePipeline
// and the external tool with events to be detected. Currently only supported
// for webhooks that target an action type of GitHub.
//
//    // Example sending a request using DeregisterWebhookWithThirdPartyRequest.
//    req := client.DeregisterWebhookWithThirdPartyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeregisterWebhookWithThirdParty
func (c *Client) DeregisterWebhookWithThirdPartyRequest(input *DeregisterWebhookWithThirdPartyInput) DeregisterWebhookWithThirdPartyRequest {
	op := &aws.Operation{
		Name:       opDeregisterWebhookWithThirdParty,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterWebhookWithThirdPartyInput{}
	}

	req := c.newRequest(op, input, &DeregisterWebhookWithThirdPartyOutput{})
	return DeregisterWebhookWithThirdPartyRequest{Request: req, Input: input, Copy: c.DeregisterWebhookWithThirdPartyRequest}
}

// DeregisterWebhookWithThirdPartyRequest is the request type for the
// DeregisterWebhookWithThirdParty API operation.
type DeregisterWebhookWithThirdPartyRequest struct {
	*aws.Request
	Input *DeregisterWebhookWithThirdPartyInput
	Copy  func(*DeregisterWebhookWithThirdPartyInput) DeregisterWebhookWithThirdPartyRequest
}

// Send marshals and sends the DeregisterWebhookWithThirdParty API request.
func (r DeregisterWebhookWithThirdPartyRequest) Send(ctx context.Context) (*DeregisterWebhookWithThirdPartyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterWebhookWithThirdPartyResponse{
		DeregisterWebhookWithThirdPartyOutput: r.Request.Data.(*DeregisterWebhookWithThirdPartyOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterWebhookWithThirdPartyResponse is the response type for the
// DeregisterWebhookWithThirdParty API operation.
type DeregisterWebhookWithThirdPartyResponse struct {
	*DeregisterWebhookWithThirdPartyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterWebhookWithThirdParty request.
func (r *DeregisterWebhookWithThirdPartyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
