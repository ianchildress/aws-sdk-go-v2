// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCatalogImportStatusRequest
type GetCatalogImportStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the catalog to migrate. Currently, this should be the AWS account
	// ID.
	CatalogId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetCatalogImportStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCatalogImportStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCatalogImportStatusInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCatalogImportStatusResponse
type GetCatalogImportStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the specified catalog migration.
	ImportStatus *CatalogImportStatus `type:"structure"`
}

// String returns the string representation
func (s GetCatalogImportStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCatalogImportStatus = "GetCatalogImportStatus"

// GetCatalogImportStatusRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the status of a migration operation.
//
//    // Example sending a request using GetCatalogImportStatusRequest.
//    req := client.GetCatalogImportStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCatalogImportStatus
func (c *Client) GetCatalogImportStatusRequest(input *GetCatalogImportStatusInput) GetCatalogImportStatusRequest {
	op := &aws.Operation{
		Name:       opGetCatalogImportStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCatalogImportStatusInput{}
	}

	req := c.newRequest(op, input, &GetCatalogImportStatusOutput{})
	return GetCatalogImportStatusRequest{Request: req, Input: input, Copy: c.GetCatalogImportStatusRequest}
}

// GetCatalogImportStatusRequest is the request type for the
// GetCatalogImportStatus API operation.
type GetCatalogImportStatusRequest struct {
	*aws.Request
	Input *GetCatalogImportStatusInput
	Copy  func(*GetCatalogImportStatusInput) GetCatalogImportStatusRequest
}

// Send marshals and sends the GetCatalogImportStatus API request.
func (r GetCatalogImportStatusRequest) Send(ctx context.Context) (*GetCatalogImportStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCatalogImportStatusResponse{
		GetCatalogImportStatusOutput: r.Request.Data.(*GetCatalogImportStatusOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCatalogImportStatusResponse is the response type for the
// GetCatalogImportStatus API operation.
type GetCatalogImportStatusResponse struct {
	*GetCatalogImportStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCatalogImportStatus request.
func (r *GetCatalogImportStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
