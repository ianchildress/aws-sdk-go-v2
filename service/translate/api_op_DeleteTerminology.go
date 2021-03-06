// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/DeleteTerminologyRequest
type DeleteTerminologyInput struct {
	_ struct{} `type:"structure"`

	// The name of the custom terminology being deleted.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTerminologyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTerminologyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTerminologyInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/DeleteTerminologyOutput
type DeleteTerminologyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTerminologyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTerminology = "DeleteTerminology"

// DeleteTerminologyRequest returns a request value for making API operation for
// Amazon Translate.
//
// A synchronous action that deletes a custom terminology.
//
//    // Example sending a request using DeleteTerminologyRequest.
//    req := client.DeleteTerminologyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/DeleteTerminology
func (c *Client) DeleteTerminologyRequest(input *DeleteTerminologyInput) DeleteTerminologyRequest {
	op := &aws.Operation{
		Name:       opDeleteTerminology,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTerminologyInput{}
	}

	req := c.newRequest(op, input, &DeleteTerminologyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteTerminologyRequest{Request: req, Input: input, Copy: c.DeleteTerminologyRequest}
}

// DeleteTerminologyRequest is the request type for the
// DeleteTerminology API operation.
type DeleteTerminologyRequest struct {
	*aws.Request
	Input *DeleteTerminologyInput
	Copy  func(*DeleteTerminologyInput) DeleteTerminologyRequest
}

// Send marshals and sends the DeleteTerminology API request.
func (r DeleteTerminologyRequest) Send(ctx context.Context) (*DeleteTerminologyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTerminologyResponse{
		DeleteTerminologyOutput: r.Request.Data.(*DeleteTerminologyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTerminologyResponse is the response type for the
// DeleteTerminology API operation.
type DeleteTerminologyResponse struct {
	*DeleteTerminologyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTerminology request.
func (r *DeleteTerminologyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
