// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateLocationEfsRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateLocationEfsRequest
type CreateLocationEfsInput struct {
	_ struct{} `type:"structure"`

	// The subnet and security group that the Amazon EFS file system uses.
	//
	// Ec2Config is a required field
	Ec2Config *Ec2Config `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	//
	// EfsFilesystemArn is a required field
	EfsFilesystemArn *string `type:"string" required:"true"`

	// A subdirectory in the location’s path. This subdirectory in the EFS file
	// system is used to read data from the EFS source location or write data to
	// the EFS destination. By default, AWS DataSync uses the root directory.
	//
	// Subdirectory is a required field
	Subdirectory *string `type:"string" required:"true"`

	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []TagListEntry `type:"list"`
}

// String returns the string representation
func (s CreateLocationEfsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLocationEfsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLocationEfsInput"}

	if s.Ec2Config == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ec2Config"))
	}

	if s.EfsFilesystemArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EfsFilesystemArn"))
	}

	if s.Subdirectory == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subdirectory"))
	}
	if s.Ec2Config != nil {
		if err := s.Ec2Config.Validate(); err != nil {
			invalidParams.AddNested("Ec2Config", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CreateLocationEfs
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateLocationEfsResponse
type CreateLocationEfsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that
	// is created.
	LocationArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLocationEfsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLocationEfs = "CreateLocationEfs"

// CreateLocationEfsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Creates an endpoint for an Amazon EFS file system.
//
//    // Example sending a request using CreateLocationEfsRequest.
//    req := client.CreateLocationEfsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateLocationEfs
func (c *Client) CreateLocationEfsRequest(input *CreateLocationEfsInput) CreateLocationEfsRequest {
	op := &aws.Operation{
		Name:       opCreateLocationEfs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLocationEfsInput{}
	}

	req := c.newRequest(op, input, &CreateLocationEfsOutput{})
	return CreateLocationEfsRequest{Request: req, Input: input, Copy: c.CreateLocationEfsRequest}
}

// CreateLocationEfsRequest is the request type for the
// CreateLocationEfs API operation.
type CreateLocationEfsRequest struct {
	*aws.Request
	Input *CreateLocationEfsInput
	Copy  func(*CreateLocationEfsInput) CreateLocationEfsRequest
}

// Send marshals and sends the CreateLocationEfs API request.
func (r CreateLocationEfsRequest) Send(ctx context.Context) (*CreateLocationEfsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLocationEfsResponse{
		CreateLocationEfsOutput: r.Request.Data.(*CreateLocationEfsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLocationEfsResponse is the response type for the
// CreateLocationEfs API operation.
type CreateLocationEfsResponse struct {
	*CreateLocationEfsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLocationEfs request.
func (r *CreateLocationEfsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}