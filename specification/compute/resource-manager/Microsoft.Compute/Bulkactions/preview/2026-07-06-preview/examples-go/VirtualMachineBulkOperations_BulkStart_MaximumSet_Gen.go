package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/VirtualMachineBulkOperations_BulkStart_MaximumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_BulkStartOperation_virtualMachineBulkOperationsBulkStartExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkStartOperation(ctx, "rgBulkactions", "useast2euap", armbulkactions.ExecuteStartContent{
		ExecutionParameters: &armbulkactions.ExecutionParameters{
			RetryPolicy: &armbulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](19),
				OnFailureAction:      to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
			},
		},
		ResourcesWithContext: &armbulkactions.ResourcesWithContext{
			Resources: []*armbulkactions.ResourceWithContext{
				{
					ResourceID:      to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
					ResourceContext: to.Ptr("startContext"),
				},
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
	// 		Location: to.Ptr("useast2euap"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				ErrorCode: to.Ptr("TestErrorCode"),
	// 				ErrorDetails: to.Ptr("Test error details"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("dbfcca08-7423-421b-a45f-a5cf5d00a85c"),
	// 					ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
	// 					SubscriptionID: to.Ptr("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-06-11T19:35:45.098Z"); return t}()),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("TestErrorCode"),
	// 						ErrorDetails: to.Ptr("Test error details"),
	// 					},
	// 					FallbackOperationInfo: &armbulkactions.FallbackOperationInfo{
	// 						LastOpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 						Status: to.Ptr("succeeded"),
	// 						Error: &armbulkactions.ResourceOperationError{
	// 							ErrorCode: to.Ptr("TestErrorCode"),
	// 							ErrorDetails: to.Ptr("Test error details"),
	// 						},
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-06-11T19:35:45.098Z"); return t}()),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](19),
	// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
	// 					},
	// 					ResourceNotificationDetails: &armbulkactions.ResourceNotificationDetails{
	// 						ResourceContext: to.Ptr("startContext"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
