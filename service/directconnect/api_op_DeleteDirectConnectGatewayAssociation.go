// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteDirectConnectGatewayAssociationRequest
type DeleteDirectConnectGatewayAssociationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Direct Connect gateway association.
	AssociationId *string `locationName:"associationId" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string"`

	// The ID of the virtual private gateway.
	VirtualGatewayId *string `locationName:"virtualGatewayId" deprecated:"true" type:"string"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteDirectConnectGatewayAssociationResult
type DeleteDirectConnectGatewayAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted association.
	DirectConnectGatewayAssociation *DirectConnectGatewayAssociation `locationName:"directConnectGatewayAssociation" type:"structure"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDirectConnectGatewayAssociation = "DeleteDirectConnectGatewayAssociation"

// DeleteDirectConnectGatewayAssociationRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes the association between the specified Direct Connect gateway and
// virtual private gateway.
//
//    // Example sending a request using DeleteDirectConnectGatewayAssociationRequest.
//    req := client.DeleteDirectConnectGatewayAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteDirectConnectGatewayAssociation
func (c *Client) DeleteDirectConnectGatewayAssociationRequest(input *DeleteDirectConnectGatewayAssociationInput) DeleteDirectConnectGatewayAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteDirectConnectGatewayAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDirectConnectGatewayAssociationInput{}
	}

	req := c.newRequest(op, input, &DeleteDirectConnectGatewayAssociationOutput{})
	return DeleteDirectConnectGatewayAssociationRequest{Request: req, Input: input, Copy: c.DeleteDirectConnectGatewayAssociationRequest}
}

// DeleteDirectConnectGatewayAssociationRequest is the request type for the
// DeleteDirectConnectGatewayAssociation API operation.
type DeleteDirectConnectGatewayAssociationRequest struct {
	*aws.Request
	Input *DeleteDirectConnectGatewayAssociationInput
	Copy  func(*DeleteDirectConnectGatewayAssociationInput) DeleteDirectConnectGatewayAssociationRequest
}

// Send marshals and sends the DeleteDirectConnectGatewayAssociation API request.
func (r DeleteDirectConnectGatewayAssociationRequest) Send(ctx context.Context) (*DeleteDirectConnectGatewayAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDirectConnectGatewayAssociationResponse{
		DeleteDirectConnectGatewayAssociationOutput: r.Request.Data.(*DeleteDirectConnectGatewayAssociationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDirectConnectGatewayAssociationResponse is the response type for the
// DeleteDirectConnectGatewayAssociation API operation.
type DeleteDirectConnectGatewayAssociationResponse struct {
	*DeleteDirectConnectGatewayAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDirectConnectGatewayAssociation request.
func (r *DeleteDirectConnectGatewayAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
