// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomainEntryRequest
type DeleteDomainEntryInput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about your domain entries.
	//
	// DomainEntry is a required field
	DomainEntry *DomainEntry `locationName:"domainEntry" type:"structure" required:"true"`

	// The name of the domain entry to delete.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainEntryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainEntryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainEntryInput"}

	if s.DomainEntry == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainEntry"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomainEntryResult
type DeleteDomainEntryOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// delete domain entry request.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s DeleteDomainEntryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDomainEntry = "DeleteDomainEntry"

// DeleteDomainEntryRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a specific domain entry.
//
// The delete domain entry operation supports tag-based access control via resource
// tags applied to the resource identified by domainName. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteDomainEntryRequest.
//    req := client.DeleteDomainEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDomainEntry
func (c *Client) DeleteDomainEntryRequest(input *DeleteDomainEntryInput) DeleteDomainEntryRequest {
	op := &aws.Operation{
		Name:       opDeleteDomainEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDomainEntryInput{}
	}

	req := c.newRequest(op, input, &DeleteDomainEntryOutput{})
	return DeleteDomainEntryRequest{Request: req, Input: input, Copy: c.DeleteDomainEntryRequest}
}

// DeleteDomainEntryRequest is the request type for the
// DeleteDomainEntry API operation.
type DeleteDomainEntryRequest struct {
	*aws.Request
	Input *DeleteDomainEntryInput
	Copy  func(*DeleteDomainEntryInput) DeleteDomainEntryRequest
}

// Send marshals and sends the DeleteDomainEntry API request.
func (r DeleteDomainEntryRequest) Send(ctx context.Context) (*DeleteDomainEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainEntryResponse{
		DeleteDomainEntryOutput: r.Request.Data.(*DeleteDomainEntryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainEntryResponse is the response type for the
// DeleteDomainEntry API operation.
type DeleteDomainEntryResponse struct {
	*DeleteDomainEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomainEntry request.
func (r *DeleteDomainEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
