// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomainRequest
type DeleteDomainInput struct {
	_ struct{} `type:"structure"`

	// The specific domain name to delete.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomainResult
type DeleteDomainOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// delete domain request.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s DeleteDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDomain = "DeleteDomain"

// DeleteDomainRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes the specified domain recordset and all of its domain records.
//
// The delete domain operation supports tag-based access control via resource
// tags applied to the resource identified by domainName. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteDomainRequest.
//    req := client.DeleteDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomain
func (c *Client) DeleteDomainRequest(input *DeleteDomainInput) DeleteDomainRequest {
	op := &aws.Operation{
		Name:       opDeleteDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDomainInput{}
	}

	req := c.newRequest(op, input, &DeleteDomainOutput{})
	return DeleteDomainRequest{Request: req, Input: input, Copy: c.DeleteDomainRequest}
}

// DeleteDomainRequest is the request type for the
// DeleteDomain API operation.
type DeleteDomainRequest struct {
	*aws.Request
	Input *DeleteDomainInput
	Copy  func(*DeleteDomainInput) DeleteDomainRequest
}

// Send marshals and sends the DeleteDomain API request.
func (r DeleteDomainRequest) Send(ctx context.Context) (*DeleteDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainResponse{
		DeleteDomainOutput: r.Request.Data.(*DeleteDomainOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainResponse is the response type for the
// DeleteDomain API operation.
type DeleteDomainResponse struct {
	*DeleteDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomain request.
func (r *DeleteDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
