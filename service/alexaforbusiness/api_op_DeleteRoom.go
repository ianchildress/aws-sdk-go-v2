// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteRoomRequest
type DeleteRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room to delete. Required.
	RoomArn *string `type:"string"`
}

// String returns the string representation
func (s DeleteRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteRoomResponse
type DeleteRoomOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRoomOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRoom = "DeleteRoom"

// DeleteRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a room by the room ARN.
//
//    // Example sending a request using DeleteRoomRequest.
//    req := client.DeleteRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteRoom
func (c *Client) DeleteRoomRequest(input *DeleteRoomInput) DeleteRoomRequest {
	op := &aws.Operation{
		Name:       opDeleteRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRoomInput{}
	}

	req := c.newRequest(op, input, &DeleteRoomOutput{})
	return DeleteRoomRequest{Request: req, Input: input, Copy: c.DeleteRoomRequest}
}

// DeleteRoomRequest is the request type for the
// DeleteRoom API operation.
type DeleteRoomRequest struct {
	*aws.Request
	Input *DeleteRoomInput
	Copy  func(*DeleteRoomInput) DeleteRoomRequest
}

// Send marshals and sends the DeleteRoom API request.
func (r DeleteRoomRequest) Send(ctx context.Context) (*DeleteRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRoomResponse{
		DeleteRoomOutput: r.Request.Data.(*DeleteRoomOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRoomResponse is the response type for the
// DeleteRoom API operation.
type DeleteRoomResponse struct {
	*DeleteRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoom request.
func (r *DeleteRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
