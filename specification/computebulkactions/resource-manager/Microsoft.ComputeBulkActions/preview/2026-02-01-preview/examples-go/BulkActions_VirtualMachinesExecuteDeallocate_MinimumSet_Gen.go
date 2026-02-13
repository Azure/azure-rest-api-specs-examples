package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_VirtualMachinesExecuteDeallocate_MinimumSet_Gen.json
func ExampleBulkActionsClient_VirtualMachinesExecuteDeallocate_bulkActionsVirtualMachinesExecuteDeallocateMinimumSetGenGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkActionsClient().VirtualMachinesExecuteDeallocate(ctx, "eastus2euap", armcomputebulkactions.ExecuteDeallocateRequest{
		ExecutionParameters: &armcomputebulkactions.ExecutionParameters{},
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
	// res = armcomputebulkactions.BulkActionsClientVirtualMachinesExecuteDeallocateResponse{
	// 	DeallocateResourceOperationResponse: &armcomputebulkactions.DeallocateResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Description: to.Ptr("Deallocate Resource Request"),
	// 	},
	// }
}
