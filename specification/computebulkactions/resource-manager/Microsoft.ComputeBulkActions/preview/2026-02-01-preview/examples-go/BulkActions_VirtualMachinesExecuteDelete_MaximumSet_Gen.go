package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_VirtualMachinesExecuteDelete_MaximumSet_Gen.json
func ExampleBulkActionsClient_VirtualMachinesExecuteDelete_bulkActionsVirtualMachinesExecuteDeleteMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkActionsClient().VirtualMachinesExecuteDelete(ctx, "east2us2euap", armcomputebulkactions.ExecuteDeleteRequest{
		ExecutionParameters: &armcomputebulkactions.ExecutionParameters{
			RetryPolicy: &armcomputebulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](45),
			},
		},
		Resources: &armcomputebulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
			},
		},
		Correlationid: to.Ptr("dfe927c5-16a6-40b7-a0f7-8524975ed642"),
		ForceDeletion: to.Ptr(true),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputebulkactions.BulkActionsClientVirtualMachinesExecuteDeleteResponse{
	// 	DeleteResourceOperationResponse: &armcomputebulkactions.DeleteResourceOperationResponse{
	// 		Type: to.Ptr("virtualmachines"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Description: to.Ptr("Delete Resource response"),
	// 		Results: []*armcomputebulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ErrorCode: to.Ptr(""),
	// 				ErrorDetails: to.Ptr(""),
	// 				Operation: &armcomputebulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("2a3fce8e-874c-45f4-9d27-1a804f3b7a0f"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 					OpType: to.Ptr(armcomputebulkactions.ResourceOperationTypeDelete),
	// 					SubscriptionID: to.Ptr("YourSubscriptionId"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputebulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputebulkactions.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("UTC"),
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					RetryPolicy: &armcomputebulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](45),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 				ErrorCode: to.Ptr(""),
	// 				ErrorDetails: to.Ptr(""),
	// 				Operation: &armcomputebulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("264f0c8a-4d5f-496c-80df-b438624ce55f"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 					OpType: to.Ptr(armcomputebulkactions.ResourceOperationTypeDelete),
	// 					SubscriptionID: to.Ptr("YourSubscriptionId"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputebulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputebulkactions.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("UTC"),
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					RetryPolicy: &armcomputebulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](45),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
