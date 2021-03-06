// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// To be written.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/AddCommunicationToCaseRequest
type AddCommunicationToCaseInput struct {
	_ struct{} `type:"structure"`

	// The ID of a set of one or more attachments for the communication to add to
	// the case. Create the set by calling AddAttachmentsToSet
	AttachmentSetId *string `locationName:"attachmentSetId" type:"string"`

	// The AWS Support case ID requested or returned in the call. The case ID is
	// an alphanumeric string formatted as shown in this example: case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string `locationName:"caseId" type:"string"`

	// The email addresses in the CC line of an email to be added to the support
	// case.
	CcEmailAddresses []string `locationName:"ccEmailAddresses" type:"list"`

	// The body of an email communication to add to the support case.
	//
	// CommunicationBody is a required field
	CommunicationBody *string `locationName:"communicationBody" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AddCommunicationToCaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddCommunicationToCaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddCommunicationToCaseInput"}

	if s.CommunicationBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommunicationBody"))
	}
	if s.CommunicationBody != nil && len(*s.CommunicationBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CommunicationBody", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of the AddCommunicationToCase operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/AddCommunicationToCaseResponse
type AddCommunicationToCaseOutput struct {
	_ struct{} `type:"structure"`

	// True if AddCommunicationToCase succeeds. Otherwise, returns an error.
	Result *bool `locationName:"result" type:"boolean"`
}

// String returns the string representation
func (s AddCommunicationToCaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddCommunicationToCase = "AddCommunicationToCase"

// AddCommunicationToCaseRequest returns a request value for making API operation for
// AWS Support.
//
// Adds additional customer communication to an AWS Support case. You use the
// caseId value to identify the case to add communication to. You can list a
// set of email addresses to copy on the communication using the ccEmailAddresses
// value. The communicationBody value contains the text of the communication.
//
// The response indicates the success or failure of the request.
//
// This operation implements a subset of the features of the AWS Support Center.
//
//    // Example sending a request using AddCommunicationToCaseRequest.
//    req := client.AddCommunicationToCaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/AddCommunicationToCase
func (c *Client) AddCommunicationToCaseRequest(input *AddCommunicationToCaseInput) AddCommunicationToCaseRequest {
	op := &aws.Operation{
		Name:       opAddCommunicationToCase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddCommunicationToCaseInput{}
	}

	req := c.newRequest(op, input, &AddCommunicationToCaseOutput{})
	return AddCommunicationToCaseRequest{Request: req, Input: input, Copy: c.AddCommunicationToCaseRequest}
}

// AddCommunicationToCaseRequest is the request type for the
// AddCommunicationToCase API operation.
type AddCommunicationToCaseRequest struct {
	*aws.Request
	Input *AddCommunicationToCaseInput
	Copy  func(*AddCommunicationToCaseInput) AddCommunicationToCaseRequest
}

// Send marshals and sends the AddCommunicationToCase API request.
func (r AddCommunicationToCaseRequest) Send(ctx context.Context) (*AddCommunicationToCaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddCommunicationToCaseResponse{
		AddCommunicationToCaseOutput: r.Request.Data.(*AddCommunicationToCaseOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddCommunicationToCaseResponse is the response type for the
// AddCommunicationToCase API operation.
type AddCommunicationToCaseResponse struct {
	*AddCommunicationToCaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddCommunicationToCase request.
func (r *AddCommunicationToCaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
