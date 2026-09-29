package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkDelete_ForceDeleteSuccess.json
func ExampleVirtualMachineBulkOperationsClient_BulkDeleteOperation_twoForceDeleteMultipleVirtualMachines() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkDeleteOperation(ctx, "example-rg", "eastus", armbulkactions.ExecuteDeleteContent{
		ExecutionParameters: &armbulkactions.ExecutionParameters{},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
			},
		},
		ForceDeletion: to.Ptr(true),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkDeleteOperationResponse{
	// 	DeleteResourceOperationResponse: armbulkactions.DeleteResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Description: to.Ptr("Delete Resource request"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("2a9a732e-5572-4f62-9168-457941b38f36"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeDelete),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("b6a7f971-2cd6-43af-a2c5-31f928d7e460"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-02"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeDelete),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
