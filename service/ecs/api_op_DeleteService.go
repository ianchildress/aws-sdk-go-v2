// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeleteServiceRequest
type DeleteServiceInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the service to delete. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// If true, allows you to delete a service even if it has not been scaled down
	// to zero tasks. It is only necessary to use this if the service is using the
	// REPLICA scheduling strategy.
	Force *bool `locationName:"force" type:"boolean"`

	// The name of the service to delete.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteServiceInput"}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeleteServiceResponse
type DeleteServiceOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deleted service.
	Service *Service `locationName:"service" type:"structure"`
}

// String returns the string representation
func (s DeleteServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteService = "DeleteService"

// DeleteServiceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deletes a specified service within a cluster. You can delete a service if
// you have no running tasks in it and the desired task count is zero. If the
// service is actively maintaining tasks, you cannot delete it, and you must
// update the service to a desired task count of zero. For more information,
// see UpdateService.
//
// When you delete a service, if there are still running tasks that require
// cleanup, the service status moves from ACTIVE to DRAINING, and the service
// is no longer visible in the console or in the ListServices API operation.
// After the tasks have stopped, then the service status moves from DRAINING
// to INACTIVE. Services in the DRAINING or INACTIVE status can still be viewed
// with the DescribeServices API operation. However, in the future, INACTIVE
// services may be cleaned up and purged from Amazon ECS record keeping, and
// DescribeServices calls on those services return a ServiceNotFoundException
// error.
//
// If you attempt to create a new service with the same name as an existing
// service in either ACTIVE or DRAINING status, you receive an error.
//
//    // Example sending a request using DeleteServiceRequest.
//    req := client.DeleteServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeleteService
func (c *Client) DeleteServiceRequest(input *DeleteServiceInput) DeleteServiceRequest {
	op := &aws.Operation{
		Name:       opDeleteService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceInput{}
	}

	req := c.newRequest(op, input, &DeleteServiceOutput{})
	return DeleteServiceRequest{Request: req, Input: input, Copy: c.DeleteServiceRequest}
}

// DeleteServiceRequest is the request type for the
// DeleteService API operation.
type DeleteServiceRequest struct {
	*aws.Request
	Input *DeleteServiceInput
	Copy  func(*DeleteServiceInput) DeleteServiceRequest
}

// Send marshals and sends the DeleteService API request.
func (r DeleteServiceRequest) Send(ctx context.Context) (*DeleteServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServiceResponse{
		DeleteServiceOutput: r.Request.Data.(*DeleteServiceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServiceResponse is the response type for the
// DeleteService API operation.
type DeleteServiceResponse struct {
	*DeleteServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteService request.
func (r *DeleteServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
