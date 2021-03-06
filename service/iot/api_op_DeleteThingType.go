// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DeleteThingType operation.
type DeleteThingTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing type.
	//
	// ThingTypeName is a required field
	ThingTypeName *string `location:"uri" locationName:"thingTypeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteThingTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteThingTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteThingTypeInput"}

	if s.ThingTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingTypeName"))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingTypeInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the DeleteThingType operation.
type DeleteThingTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteThingTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteThingType = "DeleteThingType"

// DeleteThingTypeRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes the specified thing type. You cannot delete a thing type if it has
// things associated with it. To delete a thing type, first mark it as deprecated
// by calling DeprecateThingType, then remove any associated things by calling
// UpdateThing to change the thing type on any associated thing, and finally
// use DeleteThingType to delete the thing type.
//
//    // Example sending a request using DeleteThingTypeRequest.
//    req := client.DeleteThingTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteThingTypeRequest(input *DeleteThingTypeInput) DeleteThingTypeRequest {
	op := &aws.Operation{
		Name:       opDeleteThingType,
		HTTPMethod: "DELETE",
		HTTPPath:   "/thing-types/{thingTypeName}",
	}

	if input == nil {
		input = &DeleteThingTypeInput{}
	}

	req := c.newRequest(op, input, &DeleteThingTypeOutput{})
	return DeleteThingTypeRequest{Request: req, Input: input, Copy: c.DeleteThingTypeRequest}
}

// DeleteThingTypeRequest is the request type for the
// DeleteThingType API operation.
type DeleteThingTypeRequest struct {
	*aws.Request
	Input *DeleteThingTypeInput
	Copy  func(*DeleteThingTypeInput) DeleteThingTypeRequest
}

// Send marshals and sends the DeleteThingType API request.
func (r DeleteThingTypeRequest) Send(ctx context.Context) (*DeleteThingTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteThingTypeResponse{
		DeleteThingTypeOutput: r.Request.Data.(*DeleteThingTypeOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteThingTypeResponse is the response type for the
// DeleteThingType API operation.
type DeleteThingTypeResponse struct {
	*DeleteThingTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteThingType request.
func (r *DeleteThingTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
