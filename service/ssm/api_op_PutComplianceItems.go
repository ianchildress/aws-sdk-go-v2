// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutComplianceItemsRequest
type PutComplianceItemsInput struct {
	_ struct{} `type:"structure"`

	// Specify the compliance type. For example, specify Association (for a State
	// Manager association), Patch, or Custom:string.
	//
	// ComplianceType is a required field
	ComplianceType *string `min:"1" type:"string" required:"true"`

	// A summary of the call execution that includes an execution ID, the type of
	// execution (for example, Command), and the date/time of the execution using
	// a datetime object that is saved in the following format: yyyy-MM-dd'T'HH:mm:ss'Z'.
	//
	// ExecutionSummary is a required field
	ExecutionSummary *ComplianceExecutionSummary `type:"structure" required:"true"`

	// MD5 or SHA-256 content hash. The content hash is used to determine if existing
	// information should be overwritten or ignored. If the content hashes match,
	// the request to put compliance information is ignored.
	ItemContentHash *string `type:"string"`

	// Information about the compliance as defined by the resource type. For example,
	// for a patch compliance type, Items includes information about the PatchSeverity,
	// Classification, etc.
	//
	// Items is a required field
	Items []ComplianceItemEntry `type:"list" required:"true"`

	// Specify an ID for this resource. For a managed instance, this is the instance
	// ID.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// Specify the type of resource. ManagedInstance is currently the only supported
	// resource type.
	//
	// ResourceType is a required field
	ResourceType *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutComplianceItemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutComplianceItemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutComplianceItemsInput"}

	if s.ComplianceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComplianceType"))
	}
	if s.ComplianceType != nil && len(*s.ComplianceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ComplianceType", 1))
	}

	if s.ExecutionSummary == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionSummary"))
	}

	if s.Items == nil {
		invalidParams.Add(aws.NewErrParamRequired("Items"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if s.ResourceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if s.ResourceType != nil && len(*s.ResourceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceType", 1))
	}
	if s.ExecutionSummary != nil {
		if err := s.ExecutionSummary.Validate(); err != nil {
			invalidParams.AddNested("ExecutionSummary", err.(aws.ErrInvalidParams))
		}
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutComplianceItemsResult
type PutComplianceItemsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutComplianceItemsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutComplianceItems = "PutComplianceItems"

// PutComplianceItemsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Registers a compliance type and other compliance details on a designated
// resource. This action lets you register custom compliance details with a
// resource. This call overwrites existing compliance information on the resource,
// so you must provide a full list of compliance items each time that you send
// the request.
//
// ComplianceType can be one of the following:
//
//    * ExecutionId: The execution ID when the patch, association, or custom
//    compliance item was applied.
//
//    * ExecutionType: Specify patch, association, or Custom:string.
//
//    * ExecutionTime. The time the patch, association, or custom compliance
//    item was applied to the instance.
//
//    * Id: The patch, association, or custom compliance ID.
//
//    * Title: A title.
//
//    * Status: The status of the compliance item. For example, approved for
//    patches, or Failed for associations.
//
//    * Severity: A patch severity. For example, critical.
//
//    * DocumentName: A SSM document name. For example, AWS-RunPatchBaseline.
//
//    * DocumentVersion: An SSM document version number. For example, 4.
//
//    * Classification: A patch classification. For example, security updates.
//
//    * PatchBaselineId: A patch baseline ID.
//
//    * PatchSeverity: A patch severity. For example, Critical.
//
//    * PatchState: A patch state. For example, InstancesWithFailedPatches.
//
//    * PatchGroup: The name of a patch group.
//
//    * InstalledTime: The time the association, patch, or custom compliance
//    item was applied to the resource. Specify the time by using the following
//    format: yyyy-MM-dd'T'HH:mm:ss'Z'
//
//    // Example sending a request using PutComplianceItemsRequest.
//    req := client.PutComplianceItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutComplianceItems
func (c *Client) PutComplianceItemsRequest(input *PutComplianceItemsInput) PutComplianceItemsRequest {
	op := &aws.Operation{
		Name:       opPutComplianceItems,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutComplianceItemsInput{}
	}

	req := c.newRequest(op, input, &PutComplianceItemsOutput{})
	return PutComplianceItemsRequest{Request: req, Input: input, Copy: c.PutComplianceItemsRequest}
}

// PutComplianceItemsRequest is the request type for the
// PutComplianceItems API operation.
type PutComplianceItemsRequest struct {
	*aws.Request
	Input *PutComplianceItemsInput
	Copy  func(*PutComplianceItemsInput) PutComplianceItemsRequest
}

// Send marshals and sends the PutComplianceItems API request.
func (r PutComplianceItemsRequest) Send(ctx context.Context) (*PutComplianceItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutComplianceItemsResponse{
		PutComplianceItemsOutput: r.Request.Data.(*PutComplianceItemsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutComplianceItemsResponse is the response type for the
// PutComplianceItems API operation.
type PutComplianceItemsResponse struct {
	*PutComplianceItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutComplianceItems request.
func (r *PutComplianceItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
