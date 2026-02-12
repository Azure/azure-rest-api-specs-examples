package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_VirtualMachinesExecuteStart_MaximumSet_Gen.json
func ExampleBulkActionsClient_VirtualMachinesExecuteStart_bulkActionsVirtualMachinesExecuteStartMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("D8E30CC0-2763-4FCC-84A8-3C5659281032", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkActionsClient().VirtualMachinesExecuteStart(ctx, "eastus2euap", armcomputebulkactions.ExecuteStartRequest{
		ExecutionParameters: &armcomputebulkactions.ExecutionParameters{
			RetryPolicy: &armcomputebulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](27),
			},
		},
		Resources: &armcomputebulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
			},
		},
		Correlationid: to.Ptr("4431320c-7a90-4300-b82b-73f0696ae50e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputebulkactions.BulkActionsClientVirtualMachinesExecuteStartResponse{
	// 	StartResourceOperationResponse: &armcomputebulkactions.StartResourceOperationResponse{
	// 		Type: to.Ptr("virtualMachine"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Results: []*armcomputebulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ErrorCode: to.Ptr("null"),
	// 				ErrorDetails: to.Ptr("null"),
	// 				Operation: &armcomputebulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("2a3fce8e-874c-45f4-9d27-1a804f3b7a0f"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 					OpType: to.Ptr(armcomputebulkactions.ResourceOperationTypeStart),
	// 					SubscriptionID: to.Ptr("D8E30CC0-2763-4FCC-84A8-3C5659281032"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.667Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputebulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputebulkactions.OperationStateSucceeded),
	// 					Timezone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armcomputebulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("null"),
	// 						ErrorDetails: to.Ptr("null"),
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.668Z"); return t}()),
	// 					RetryPolicy: &armcomputebulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](27),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Description: to.Ptr("Start Resource Request"),
	// 	},
	// }
}
