package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkReimage_WithReimagePayload.json
func ExampleVirtualMachineBulkOperationsClient_BulkReimageOperation_threeReimageVirtualMachinesWithPerVMTemporaryDiskSettings() {
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
		ExecutionParameters: &armbulkactions.ExecutionParameters{},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-01"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-02"),
			},
		},
		ReimageParameters: &armbulkactions.ReimagePayload{
			BaseProfile: &armbulkactions.VirtualMachineReimageParameters{
				TempDisk: to.Ptr(true),
			},
			ResourceOverrides: []*armbulkactions.ReimageResourceOverride{
				{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-02"),
					Profile: &armbulkactions.VirtualMachineReimageParameters{
						TempDisk: to.Ptr(false),
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
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("589be017-3996-450c-bdc7-6041c726d703"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationType("Reimage")),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-02"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("88b90cde-28db-42c3-b356-9c601c1799e4"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/ephemeral-vm-02"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationType("Reimage")),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
