package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesGetOperationErrors_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationErrors_scheduledActionsVirtualMachinesGetOperationErrorsMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationErrors(ctx, "eastus2", armcomputeschedule.GetOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("01234567-89ab-cdef-0123-456789abcdef"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesGetOperationErrorsResponse{
	// 	GetOperationErrorsResponse: armcomputeschedule.GetOperationErrorsResponse{
	// 		Results: []*armcomputeschedule.OperationErrorsResult{
	// 			{
	// 				OperationID: to.Ptr("8f88ead2-fba8-4df2-8eaf-c7cf68a15574"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:43.809Z"); return t}()),
	// 				ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:43.809Z"); return t}()),
	// 				CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:43.809Z"); return t}()),
	// 				OperationErrors: []*armcomputeschedule.OperationErrorDetails{
	// 					{
	// 						ErrorCode: to.Ptr("VMAllocationFailed"),
	// 						ErrorDetails: to.Ptr("Failed to allocate virtual machine due to insufficient capacity."),
	// 						Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:43.809Z"); return t}()),
	// 						TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:43.809Z"); return t}()),
	// 						AzureOperationName: to.Ptr("Microsoft.Compute/virtualMachines/deallocate"),
	// 						CrpOperationID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 					},
	// 				},
	// 				RequestErrorCode: to.Ptr("OperationFailed"),
	// 				RequestErrorDetails: to.Ptr("The operation failed due to an internal error."),
	// 			},
	// 		},
	// 	},
	// }
}
