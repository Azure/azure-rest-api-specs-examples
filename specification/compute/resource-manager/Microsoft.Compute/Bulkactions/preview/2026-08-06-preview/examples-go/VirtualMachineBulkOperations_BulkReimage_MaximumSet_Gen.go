package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/VirtualMachineBulkOperations_BulkReimage_MaximumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_BulkReimageOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkReimageOperation(ctx, "rgBulkactions", "useast2euap", armbulkactions.ExecuteReimageRequest{
		ExecutionParameters: &armbulkactions.ExecutionParameters{
			RetryPolicy: &armbulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](19),
				OnFailureAction:      to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
			},
		},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
		},
		ReimageParameters: &armbulkactions.ReimagePayload{
			BaseProfile: &armbulkactions.VirtualMachineReimageParameters{
				TempDisk:     to.Ptr(true),
				ExactVersion: to.Ptr("zjmkrnqjmzs"),
				OSProfile: &armbulkactions.OSProfileProvisioningData{
					AdminPassword: to.Ptr("<a-password-goes-here>"),
					CustomData:    to.Ptr("teyngslcznlxihiitqbul"),
				},
			},
			ResourceOverrides: []*armbulkactions.ReimageResourceOverride{
				{
					ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
					Profile: &armbulkactions.VirtualMachineReimageParameters{
						TempDisk:     to.Ptr(true),
						ExactVersion: to.Ptr("zjmkrnqjmzs"),
						OSProfile: &armbulkactions.OSProfileProvisioningData{
							AdminPassword: to.Ptr("<a-password-goes-here>"),
							CustomData:    to.Ptr("teyngslcznlxihiitqbul"),
						},
					},
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
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkReimageOperationResponse{
	// 	ReimageResourceOperationResponse: armbulkactions.ReimageResourceOperationResponse{
	// 		Description: to.Ptr("Reimage Resource request"),
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("useast2euap"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				ErrorCode: to.Ptr("TestErrorCode"),
	// 				ErrorDetails: to.Ptr("Test error details"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("0437c0c9-328e-4585-9808-31ee9460b19f"),
	// 					ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
	// 					SubscriptionID: to.Ptr("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2"),
	// 					Deadline: to.Ptr(time.Date(2026, time.June, 11, 19, 35, 45, 98000000, time.UTC)),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("TestErrorCode"),
	// 						ErrorDetails: to.Ptr("Test error details"),
	// 					},
	// 					FallbackOperationInfo: &armbulkactions.FallbackOperationInfo{
	// 						LastOpType: to.Ptr(armbulkactions.ResourceOperationType("Reimage")),
	// 						Status: to.Ptr("succeeded"),
	// 						Error: &armbulkactions.ResourceOperationError{
	// 							ErrorCode: to.Ptr("TestErrorCode"),
	// 							ErrorDetails: to.Ptr("Test error details"),
	// 						},
	// 					},
	// 					CompletedAt: to.Ptr(time.Date(2026, time.June, 11, 19, 35, 45, 98000000, time.UTC)),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](19),
	// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
	// 					},
	// 					ResourceNotificationDetails: &armbulkactions.ResourceNotificationDetails{
	// 						ResourceContext: to.Ptr(""),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
