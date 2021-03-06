// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package stsiface provides an interface to enable mocking the AWS Security Token Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package stsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// ClientAPI provides an interface to enable mocking the
// sts.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS STS.
//    func myFunc(svc stsiface.ClientAPI) bool {
//        // Make svc.AssumeRole request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sts.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        stsiface.ClientPI
//    }
//    func (m *mockClientClient) AssumeRole(input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AssumeRoleRequest(*sts.AssumeRoleInput) sts.AssumeRoleRequest

	AssumeRoleWithSAMLRequest(*sts.AssumeRoleWithSAMLInput) sts.AssumeRoleWithSAMLRequest

	AssumeRoleWithWebIdentityRequest(*sts.AssumeRoleWithWebIdentityInput) sts.AssumeRoleWithWebIdentityRequest

	DecodeAuthorizationMessageRequest(*sts.DecodeAuthorizationMessageInput) sts.DecodeAuthorizationMessageRequest

	GetCallerIdentityRequest(*sts.GetCallerIdentityInput) sts.GetCallerIdentityRequest

	GetFederationTokenRequest(*sts.GetFederationTokenInput) sts.GetFederationTokenRequest

	GetSessionTokenRequest(*sts.GetSessionTokenInput) sts.GetSessionTokenRequest
}

var _ ClientAPI = (*sts.Client)(nil)
