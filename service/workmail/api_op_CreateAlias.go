// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/CreateAliasRequest
type CreateAliasInput struct {
	_ struct{} `type:"structure"`

	// The alias to add to the member set.
	//
	// Alias is a required field
	Alias *string `min:"1" type:"string" required:"true"`

	// The member (user or group) to which this alias is added.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The organization under which the member (user or group) exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAliasInput"}

	if s.Alias == nil {
		invalidParams.Add(aws.NewErrParamRequired("Alias"))
	}
	if s.Alias != nil && len(*s.Alias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Alias", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/CreateAliasResponse
type CreateAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAlias = "CreateAlias"

// CreateAliasRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Adds an alias to the set of a given member (user or group) of Amazon WorkMail.
//
//    // Example sending a request using CreateAliasRequest.
//    req := client.CreateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/CreateAlias
func (c *Client) CreateAliasRequest(input *CreateAliasInput) CreateAliasRequest {
	op := &aws.Operation{
		Name:       opCreateAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAliasInput{}
	}

	req := c.newRequest(op, input, &CreateAliasOutput{})
	return CreateAliasRequest{Request: req, Input: input, Copy: c.CreateAliasRequest}
}

// CreateAliasRequest is the request type for the
// CreateAlias API operation.
type CreateAliasRequest struct {
	*aws.Request
	Input *CreateAliasInput
	Copy  func(*CreateAliasInput) CreateAliasRequest
}

// Send marshals and sends the CreateAlias API request.
func (r CreateAliasRequest) Send(ctx context.Context) (*CreateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAliasResponse{
		CreateAliasOutput: r.Request.Data.(*CreateAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAliasResponse is the response type for the
// CreateAlias API operation.
type CreateAliasResponse struct {
	*CreateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAlias request.
func (r *CreateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
