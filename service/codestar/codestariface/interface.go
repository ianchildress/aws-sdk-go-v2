// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codestariface provides an interface to enable mocking the AWS CodeStar service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codestariface

import (
	"github.com/aws/aws-sdk-go-v2/service/codestar"
)

// ClientAPI provides an interface to enable mocking the
// codestar.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CodeStar.
//    func myFunc(svc codestariface.ClientAPI) bool {
//        // Make svc.AssociateTeamMember request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codestar.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        codestariface.ClientPI
//    }
//    func (m *mockClientClient) AssociateTeamMember(input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error) {
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
	AssociateTeamMemberRequest(*codestar.AssociateTeamMemberInput) codestar.AssociateTeamMemberRequest

	CreateProjectRequest(*codestar.CreateProjectInput) codestar.CreateProjectRequest

	CreateUserProfileRequest(*codestar.CreateUserProfileInput) codestar.CreateUserProfileRequest

	DeleteProjectRequest(*codestar.DeleteProjectInput) codestar.DeleteProjectRequest

	DeleteUserProfileRequest(*codestar.DeleteUserProfileInput) codestar.DeleteUserProfileRequest

	DescribeProjectRequest(*codestar.DescribeProjectInput) codestar.DescribeProjectRequest

	DescribeUserProfileRequest(*codestar.DescribeUserProfileInput) codestar.DescribeUserProfileRequest

	DisassociateTeamMemberRequest(*codestar.DisassociateTeamMemberInput) codestar.DisassociateTeamMemberRequest

	ListProjectsRequest(*codestar.ListProjectsInput) codestar.ListProjectsRequest

	ListResourcesRequest(*codestar.ListResourcesInput) codestar.ListResourcesRequest

	ListTagsForProjectRequest(*codestar.ListTagsForProjectInput) codestar.ListTagsForProjectRequest

	ListTeamMembersRequest(*codestar.ListTeamMembersInput) codestar.ListTeamMembersRequest

	ListUserProfilesRequest(*codestar.ListUserProfilesInput) codestar.ListUserProfilesRequest

	TagProjectRequest(*codestar.TagProjectInput) codestar.TagProjectRequest

	UntagProjectRequest(*codestar.UntagProjectInput) codestar.UntagProjectRequest

	UpdateProjectRequest(*codestar.UpdateProjectInput) codestar.UpdateProjectRequest

	UpdateTeamMemberRequest(*codestar.UpdateTeamMemberInput) codestar.UpdateTeamMemberRequest

	UpdateUserProfileRequest(*codestar.UpdateUserProfileInput) codestar.UpdateUserProfileRequest
}

var _ ClientAPI = (*codestar.Client)(nil)
