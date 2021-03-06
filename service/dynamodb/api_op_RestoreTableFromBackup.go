// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/RestoreTableFromBackupInput
type RestoreTableFromBackupInput struct {
	_ struct{} `type:"structure"`

	// The ARN associated with the backup.
	//
	// BackupArn is a required field
	BackupArn *string `min:"37" type:"string" required:"true"`

	// The name of the new table to which the backup must be restored.
	//
	// TargetTableName is a required field
	TargetTableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreTableFromBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreTableFromBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreTableFromBackupInput"}

	if s.BackupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupArn"))
	}
	if s.BackupArn != nil && len(*s.BackupArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("BackupArn", 37))
	}

	if s.TargetTableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetTableName"))
	}
	if s.TargetTableName != nil && len(*s.TargetTableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetTableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/RestoreTableFromBackupOutput
type RestoreTableFromBackupOutput struct {
	_ struct{} `type:"structure"`

	// The description of the table created from an existing backup.
	TableDescription *TableDescription `type:"structure"`
}

// String returns the string representation
func (s RestoreTableFromBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreTableFromBackup = "RestoreTableFromBackup"

// RestoreTableFromBackupRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Creates a new table from an existing backup. Any number of users can execute
// up to 4 concurrent restores (any type of restore) in a given account.
//
// You can call RestoreTableFromBackup at a maximum rate of 10 times per second.
//
// You must manually set up the following on the restored table:
//
//    * Auto scaling policies
//
//    * IAM policies
//
//    * Cloudwatch metrics and alarms
//
//    * Tags
//
//    * Stream settings
//
//    * Time to Live (TTL) settings
//
//    // Example sending a request using RestoreTableFromBackupRequest.
//    req := client.RestoreTableFromBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/RestoreTableFromBackup
func (c *Client) RestoreTableFromBackupRequest(input *RestoreTableFromBackupInput) RestoreTableFromBackupRequest {
	op := &aws.Operation{
		Name:       opRestoreTableFromBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreTableFromBackupInput{}
	}

	req := c.newRequest(op, input, &RestoreTableFromBackupOutput{})
	return RestoreTableFromBackupRequest{Request: req, Input: input, Copy: c.RestoreTableFromBackupRequest}
}

// RestoreTableFromBackupRequest is the request type for the
// RestoreTableFromBackup API operation.
type RestoreTableFromBackupRequest struct {
	*aws.Request
	Input *RestoreTableFromBackupInput
	Copy  func(*RestoreTableFromBackupInput) RestoreTableFromBackupRequest
}

// Send marshals and sends the RestoreTableFromBackup API request.
func (r RestoreTableFromBackupRequest) Send(ctx context.Context) (*RestoreTableFromBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreTableFromBackupResponse{
		RestoreTableFromBackupOutput: r.Request.Data.(*RestoreTableFromBackupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreTableFromBackupResponse is the response type for the
// RestoreTableFromBackup API operation.
type RestoreTableFromBackupResponse struct {
	*RestoreTableFromBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreTableFromBackup request.
func (r *RestoreTableFromBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
