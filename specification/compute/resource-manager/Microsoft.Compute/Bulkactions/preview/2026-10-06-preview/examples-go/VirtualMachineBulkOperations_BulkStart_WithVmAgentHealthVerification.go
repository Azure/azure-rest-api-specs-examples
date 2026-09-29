package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkStart_WithVmAgentHealthVerification.json
func ExampleVirtualMachineBulkOperationsClient_BulkStartOperation_twoStartVirtualMachinesWithVMAgentHealthVerification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkStartOperation(ctx, "example-rg", "eastus", armbulkactions.ExecuteStartContent{
		ExecutionParameters: &armbulkactions.ExecutionParameters{
			RetryPolicy: &armbulkactions.RetryPolicy{
				RetryWindowInMinutes: to.Ptr[int32](30),
			},
			VerifyVMAgentHealth: to.Ptr(true),
		},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkStartOperationResponse{
	// 	StartResourceOperationResponse: armbulkactions.StartResourceOperationResponse{
	// 		Description: to.Ptr("Start Resource request"),
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					Deadline: to.Ptr(time.Date(2026, time.August, 31, 18, 0, 0, 0, time.UTC)),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 					Timezone: to.Ptr("UTC"),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					Deadline: to.Ptr(time.Date(2026, time.August, 31, 18, 0, 0, 0, time.UTC)),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 					Timezone: to.Ptr("UTC"),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
