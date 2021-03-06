// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iot1clickdevicesserviceiface provides an interface to enable mocking the AWS IoT 1-Click Devices Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iot1clickdevicesserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

// ClientAPI provides an interface to enable mocking the
// iot1clickdevicesservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT 1-Click Devices Service.
//    func myFunc(svc iot1clickdevicesserviceiface.ClientAPI) bool {
//        // Make svc.ClaimDevicesByClaimCode request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iot1clickdevicesservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        iot1clickdevicesserviceiface.ClientPI
//    }
//    func (m *mockClientClient) ClaimDevicesByClaimCode(input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
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
	ClaimDevicesByClaimCodeRequest(*iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) iot1clickdevicesservice.ClaimDevicesByClaimCodeRequest

	DescribeDeviceRequest(*iot1clickdevicesservice.DescribeDeviceInput) iot1clickdevicesservice.DescribeDeviceRequest

	FinalizeDeviceClaimRequest(*iot1clickdevicesservice.FinalizeDeviceClaimInput) iot1clickdevicesservice.FinalizeDeviceClaimRequest

	GetDeviceMethodsRequest(*iot1clickdevicesservice.GetDeviceMethodsInput) iot1clickdevicesservice.GetDeviceMethodsRequest

	InitiateDeviceClaimRequest(*iot1clickdevicesservice.InitiateDeviceClaimInput) iot1clickdevicesservice.InitiateDeviceClaimRequest

	InvokeDeviceMethodRequest(*iot1clickdevicesservice.InvokeDeviceMethodInput) iot1clickdevicesservice.InvokeDeviceMethodRequest

	ListDeviceEventsRequest(*iot1clickdevicesservice.ListDeviceEventsInput) iot1clickdevicesservice.ListDeviceEventsRequest

	ListDevicesRequest(*iot1clickdevicesservice.ListDevicesInput) iot1clickdevicesservice.ListDevicesRequest

	ListTagsForResourceRequest(*iot1clickdevicesservice.ListTagsForResourceInput) iot1clickdevicesservice.ListTagsForResourceRequest

	TagResourceRequest(*iot1clickdevicesservice.TagResourceInput) iot1clickdevicesservice.TagResourceRequest

	UnclaimDeviceRequest(*iot1clickdevicesservice.UnclaimDeviceInput) iot1clickdevicesservice.UnclaimDeviceRequest

	UntagResourceRequest(*iot1clickdevicesservice.UntagResourceInput) iot1clickdevicesservice.UntagResourceRequest

	UpdateDeviceStateRequest(*iot1clickdevicesservice.UpdateDeviceStateInput) iot1clickdevicesservice.UpdateDeviceStateRequest
}

var _ ClientAPI = (*iot1clickdevicesservice.Client)(nil)
