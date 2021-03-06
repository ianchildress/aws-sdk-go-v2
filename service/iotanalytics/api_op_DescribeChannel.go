// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeChannelRequest
type DescribeChannelInput struct {
	_ struct{} `type:"structure"`

	// The name of the channel whose information is retrieved.
	//
	// ChannelName is a required field
	ChannelName *string `location:"uri" locationName:"channelName" min:"1" type:"string" required:"true"`

	// If true, additional statistical information about the channel is included
	// in the response.
	IncludeStatistics *bool `location:"querystring" locationName:"includeStatistics" type:"boolean"`
}

// String returns the string representation
func (s DescribeChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeChannelInput"}

	if s.ChannelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelName"))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeChannelInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IncludeStatistics != nil {
		v := *s.IncludeStatistics

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeStatistics", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeChannelResponse
type DescribeChannelOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the channel.
	Channel *Channel `locationName:"channel" type:"structure"`

	// Statistics about the channel. Included if the 'includeStatistics' parameter
	// is set to true in the request.
	Statistics *ChannelStatistics `locationName:"statistics" type:"structure"`
}

// String returns the string representation
func (s DescribeChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Channel != nil {
		v := s.Channel

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "channel", v, metadata)
	}
	if s.Statistics != nil {
		v := s.Statistics

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "statistics", v, metadata)
	}
	return nil
}

const opDescribeChannel = "DescribeChannel"

// DescribeChannelRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a channel.
//
//    // Example sending a request using DescribeChannelRequest.
//    req := client.DescribeChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeChannel
func (c *Client) DescribeChannelRequest(input *DescribeChannelInput) DescribeChannelRequest {
	op := &aws.Operation{
		Name:       opDescribeChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/channels/{channelName}",
	}

	if input == nil {
		input = &DescribeChannelInput{}
	}

	req := c.newRequest(op, input, &DescribeChannelOutput{})
	return DescribeChannelRequest{Request: req, Input: input, Copy: c.DescribeChannelRequest}
}

// DescribeChannelRequest is the request type for the
// DescribeChannel API operation.
type DescribeChannelRequest struct {
	*aws.Request
	Input *DescribeChannelInput
	Copy  func(*DescribeChannelInput) DescribeChannelRequest
}

// Send marshals and sends the DescribeChannel API request.
func (r DescribeChannelRequest) Send(ctx context.Context) (*DescribeChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChannelResponse{
		DescribeChannelOutput: r.Request.Data.(*DescribeChannelOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChannelResponse is the response type for the
// DescribeChannel API operation.
type DescribeChannelResponse struct {
	*DescribeChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChannel request.
func (r *DescribeChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
