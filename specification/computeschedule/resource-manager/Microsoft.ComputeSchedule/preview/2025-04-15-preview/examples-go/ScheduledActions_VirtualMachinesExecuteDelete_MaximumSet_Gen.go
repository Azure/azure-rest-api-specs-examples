package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesExecuteDelete_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteDelete_scheduledActionsVirtualMachinesExecuteDeleteMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteDelete(ctx, "east", armcomputeschedule.ExecuteDeleteContent{
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{
			RetryPolicy: &armcomputeschedule.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](4),
			},
		},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
			},
		},
		Correlationid: to.Ptr("dfe927c5-16a6-40b7-a0f7-8524975ed642"),
		ForceDeletion: to.Ptr(false),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteDeleteResponse{
	// 	DeleteResourceOperationResponse: &armcomputeschedule.DeleteResourceOperationResponse{
	// 		Type: to.Ptr("virtualmachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Results: []*armcomputeschedule.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ErrorCode: to.Ptr("ynukyltendgmn"),
	// 				ErrorDetails: to.Ptr("tifeuh"),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("0505H8E4-D41A-48FB-9CA5-4AF8D93BE75F"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationType("Delete")),
	// 					SubscriptionID: to.Ptr("YourSubscriptionId"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputeschedule.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("nwugsooykqggcokphgdj"),
	// 					TimeZone: to.Ptr("qkxnxnumvfqmsmpyccv"),
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](25),
	// 						RetryWindowInMinutes: to.Ptr[int32](4),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 				ErrorCode: to.Ptr(""),
	// 				ErrorDetails: to.Ptr(""),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("Create Resource request"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationType("Delete")),
	// 					SubscriptionID: to.Ptr("YourSubscriptionId"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputeschedule.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("nwugsooykqggcokphgdj"),
	// 					TimeZone: to.Ptr("qkxnxnumvfqmsmpyccv"),
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](4),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Description: to.Ptr("Delete Resource response"),
	// 	},
	// }
}
