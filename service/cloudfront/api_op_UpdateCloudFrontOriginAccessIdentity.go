// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to update an origin access identity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/UpdateCloudFrontOriginAccessIdentityRequest
type UpdateCloudFrontOriginAccessIdentityInput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentityConfig"`

	// The identity's configuration information.
	//
	// CloudFrontOriginAccessIdentityConfig is a required field
	CloudFrontOriginAccessIdentityConfig *CloudFrontOriginAccessIdentityConfig `locationName:"CloudFrontOriginAccessIdentityConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2018-11-05/"`

	// The identity's id.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the identity's
	// configuration. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s UpdateCloudFrontOriginAccessIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCloudFrontOriginAccessIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCloudFrontOriginAccessIdentityInput"}

	if s.CloudFrontOriginAccessIdentityConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("CloudFrontOriginAccessIdentityConfig"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.CloudFrontOriginAccessIdentityConfig != nil {
		if err := s.CloudFrontOriginAccessIdentityConfig.Validate(); err != nil {
			invalidParams.AddNested("CloudFrontOriginAccessIdentityConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCloudFrontOriginAccessIdentityInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.CloudFrontOriginAccessIdentityConfig != nil {
		v := s.CloudFrontOriginAccessIdentityConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2018-11-05/"}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentityConfig", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/UpdateCloudFrontOriginAccessIdentityResult
type UpdateCloudFrontOriginAccessIdentityOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentity"`

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *CloudFrontOriginAccessIdentity `type:"structure"`

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s UpdateCloudFrontOriginAccessIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCloudFrontOriginAccessIdentityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.CloudFrontOriginAccessIdentity != nil {
		v := s.CloudFrontOriginAccessIdentity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentity", v, metadata)
	}
	return nil
}

const opUpdateCloudFrontOriginAccessIdentity = "UpdateCloudFrontOriginAccessIdentity2018_11_05"

// UpdateCloudFrontOriginAccessIdentityRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Update an origin access identity.
//
//    // Example sending a request using UpdateCloudFrontOriginAccessIdentityRequest.
//    req := client.UpdateCloudFrontOriginAccessIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/UpdateCloudFrontOriginAccessIdentity
func (c *Client) UpdateCloudFrontOriginAccessIdentityRequest(input *UpdateCloudFrontOriginAccessIdentityInput) UpdateCloudFrontOriginAccessIdentityRequest {
	op := &aws.Operation{
		Name:       opUpdateCloudFrontOriginAccessIdentity,
		HTTPMethod: "PUT",
		HTTPPath:   "/2018-11-05/origin-access-identity/cloudfront/{Id}/config",
	}

	if input == nil {
		input = &UpdateCloudFrontOriginAccessIdentityInput{}
	}

	req := c.newRequest(op, input, &UpdateCloudFrontOriginAccessIdentityOutput{})
	return UpdateCloudFrontOriginAccessIdentityRequest{Request: req, Input: input, Copy: c.UpdateCloudFrontOriginAccessIdentityRequest}
}

// UpdateCloudFrontOriginAccessIdentityRequest is the request type for the
// UpdateCloudFrontOriginAccessIdentity API operation.
type UpdateCloudFrontOriginAccessIdentityRequest struct {
	*aws.Request
	Input *UpdateCloudFrontOriginAccessIdentityInput
	Copy  func(*UpdateCloudFrontOriginAccessIdentityInput) UpdateCloudFrontOriginAccessIdentityRequest
}

// Send marshals and sends the UpdateCloudFrontOriginAccessIdentity API request.
func (r UpdateCloudFrontOriginAccessIdentityRequest) Send(ctx context.Context) (*UpdateCloudFrontOriginAccessIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCloudFrontOriginAccessIdentityResponse{
		UpdateCloudFrontOriginAccessIdentityOutput: r.Request.Data.(*UpdateCloudFrontOriginAccessIdentityOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCloudFrontOriginAccessIdentityResponse is the response type for the
// UpdateCloudFrontOriginAccessIdentity API operation.
type UpdateCloudFrontOriginAccessIdentityResponse struct {
	*UpdateCloudFrontOriginAccessIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCloudFrontOriginAccessIdentity request.
func (r *UpdateCloudFrontOriginAccessIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
