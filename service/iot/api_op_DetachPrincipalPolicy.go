// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the DetachPrincipalPolicy operation.
type DetachPrincipalPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the policy to detach.
	//
	// PolicyName is a required field
	PolicyName *string `location:"uri" locationName:"policyName" min:"1" type:"string" required:"true"`

	// The principal.
	//
	// If the principal is a certificate, specify the certificate ARN. If the principal
	// is an Amazon Cognito identity, specify the identity ID.
	//
	// Principal is a required field
	Principal *string `location:"header" locationName:"x-amzn-iot-principal" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachPrincipalPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachPrincipalPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachPrincipalPolicyInput"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachPrincipalPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amzn-iot-principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DetachPrincipalPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachPrincipalPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachPrincipalPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDetachPrincipalPolicy = "DetachPrincipalPolicy"

// DetachPrincipalPolicyRequest returns a request value for making API operation for
// AWS IoT.
//
// Removes the specified policy from the specified certificate.
//
// Note: This API is deprecated. Please use DetachPolicy instead.
//
//    // Example sending a request using DetachPrincipalPolicyRequest.
//    req := client.DetachPrincipalPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DetachPrincipalPolicyRequest(input *DetachPrincipalPolicyInput) DetachPrincipalPolicyRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, DetachPrincipalPolicy, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opDetachPrincipalPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/principal-policies/{policyName}",
	}

	if input == nil {
		input = &DetachPrincipalPolicyInput{}
	}

	req := c.newRequest(op, input, &DetachPrincipalPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachPrincipalPolicyRequest{Request: req, Input: input, Copy: c.DetachPrincipalPolicyRequest}
}

// DetachPrincipalPolicyRequest is the request type for the
// DetachPrincipalPolicy API operation.
type DetachPrincipalPolicyRequest struct {
	*aws.Request
	Input *DetachPrincipalPolicyInput
	Copy  func(*DetachPrincipalPolicyInput) DetachPrincipalPolicyRequest
}

// Send marshals and sends the DetachPrincipalPolicy API request.
func (r DetachPrincipalPolicyRequest) Send(ctx context.Context) (*DetachPrincipalPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachPrincipalPolicyResponse{
		DetachPrincipalPolicyOutput: r.Request.Data.(*DetachPrincipalPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachPrincipalPolicyResponse is the response type for the
// DetachPrincipalPolicy API operation.
type DetachPrincipalPolicyResponse struct {
	*DetachPrincipalPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachPrincipalPolicy request.
func (r *DetachPrincipalPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
