// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RemoveTagsFromResourceMessage
type RemoveTagsFromResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Neptune resource that the tags are removed from. This value is
	// an Amazon Resource Name (ARN). For information about creating an ARN, see
	// Constructing an Amazon Resource Name (ARN) (http://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing).
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`

	// The tag key (name) of the tag to be removed.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RemoveTagsFromResourceOutput
type RemoveTagsFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveTagsFromResource = "RemoveTagsFromResource"

// RemoveTagsFromResourceRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Removes metadata tags from an Amazon Neptune resource.
//
//    // Example sending a request using RemoveTagsFromResourceRequest.
//    req := client.RemoveTagsFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RemoveTagsFromResource
func (c *Client) RemoveTagsFromResourceRequest(input *RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromResourceInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsFromResourceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveTagsFromResourceRequest{Request: req, Input: input, Copy: c.RemoveTagsFromResourceRequest}
}

// RemoveTagsFromResourceRequest is the request type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceRequest struct {
	*aws.Request
	Input *RemoveTagsFromResourceInput
	Copy  func(*RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest
}

// Send marshals and sends the RemoveTagsFromResource API request.
func (r RemoveTagsFromResourceRequest) Send(ctx context.Context) (*RemoveTagsFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromResourceResponse{
		RemoveTagsFromResourceOutput: r.Request.Data.(*RemoveTagsFromResourceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromResourceResponse is the response type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceResponse struct {
	*RemoveTagsFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromResource request.
func (r *RemoveTagsFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}