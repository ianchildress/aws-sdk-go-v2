// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents the input of a DeleteCacheSecurityGroup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheSecurityGroupMessage
type DeleteCacheSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache security group to delete.
	//
	// You cannot delete the default security group.
	//
	// CacheSecurityGroupName is a required field
	CacheSecurityGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCacheSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCacheSecurityGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCacheSecurityGroupInput"}

	if s.CacheSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSecurityGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheSecurityGroupOutput
type DeleteCacheSecurityGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCacheSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCacheSecurityGroup = "DeleteCacheSecurityGroup"

// DeleteCacheSecurityGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes a cache security group.
//
// You cannot delete a cache security group if it is associated with any clusters.
//
//    // Example sending a request using DeleteCacheSecurityGroupRequest.
//    req := client.DeleteCacheSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheSecurityGroup
func (c *Client) DeleteCacheSecurityGroupRequest(input *DeleteCacheSecurityGroupInput) DeleteCacheSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteCacheSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCacheSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteCacheSecurityGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCacheSecurityGroupRequest{Request: req, Input: input, Copy: c.DeleteCacheSecurityGroupRequest}
}

// DeleteCacheSecurityGroupRequest is the request type for the
// DeleteCacheSecurityGroup API operation.
type DeleteCacheSecurityGroupRequest struct {
	*aws.Request
	Input *DeleteCacheSecurityGroupInput
	Copy  func(*DeleteCacheSecurityGroupInput) DeleteCacheSecurityGroupRequest
}

// Send marshals and sends the DeleteCacheSecurityGroup API request.
func (r DeleteCacheSecurityGroupRequest) Send(ctx context.Context) (*DeleteCacheSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCacheSecurityGroupResponse{
		DeleteCacheSecurityGroupOutput: r.Request.Data.(*DeleteCacheSecurityGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCacheSecurityGroupResponse is the response type for the
// DeleteCacheSecurityGroup API operation.
type DeleteCacheSecurityGroupResponse struct {
	*DeleteCacheSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCacheSecurityGroup request.
func (r *DeleteCacheSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
