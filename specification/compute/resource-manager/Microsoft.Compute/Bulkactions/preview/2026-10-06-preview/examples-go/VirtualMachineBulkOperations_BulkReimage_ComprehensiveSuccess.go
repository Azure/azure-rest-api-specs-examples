package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkReimage_ComprehensiveSuccess.json
func ExampleVirtualMachineBulkOperationsClient_BulkReimageOperation_twoReimageVirtualMachinesWithSharedSettingsAndAPerVMOverride() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkReimageOperation(ctx, "example-rg", "eastus", armbulkactions.ExecuteReimageRequest{
		ExecutionParameters: &armbulkactions.ExecutionParameters{
			RetryPolicy: &armbulkactions.RetryPolicy{
				RetryWindowInMinutes: to.Ptr[int32](30),
			},
		},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
			},
		},
		ReimageParameters: &armbulkactions.ReimagePayload{
			BaseProfile: &armbulkactions.VirtualMachineReimageParameters{
				TempDisk:     to.Ptr(false),
				ExactVersion: to.Ptr("1.0.0"),
				OSProfile: &armbulkactions.OSProfileProvisioningData{
					CustomData: to.Ptr("I2Nsb3VkLWNvbmZpZwpwYWNrYWdlX3VwZ3JhZGU6IHRydWUK"),
				},
			},
			ResourceOverrides: []*armbulkactions.ReimageResourceOverride{
				{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
					Profile: &armbulkactions.VirtualMachineReimageParameters{
						TempDisk:     to.Ptr(false),
						ExactVersion: to.Ptr("1.1.0"),
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
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Description: to.Ptr("Reimage Resource request"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationType("Reimage")),
	// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					Deadline: to.Ptr(time.Date(2026, time.September, 15, 18, 0, 0, 0, time.UTC)),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 					Timezone: to.Ptr("UTC"),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationType("Reimage")),
	// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					Deadline: to.Ptr(time.Date(2026, time.September, 15, 18, 0, 0, 0, time.UTC)),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 					Timezone: to.Ptr("UTC"),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
