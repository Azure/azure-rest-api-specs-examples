package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkHibernate_VmNotFoundError.json
func ExampleVirtualMachineBulkOperationsClient_BulkHibernateOperation_threeResponseWhenAVirtualMachineDoesNotExist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkHibernateOperation(ctx, "example-rg", "eastus", armbulkactions.ExecuteHibernateContent{
		ExecutionParameters: &armbulkactions.ExecutionParameters{},
		Resources: &armbulkactions.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/missing-vm"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkHibernateOperationResponse{
	// 	HibernateResourceOperationResponse: armbulkactions.HibernateResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Description: to.Ptr("Hibernate Resource request"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeHibernate),
	// 					State: to.Ptr(armbulkactions.OperationState("PendingScheduling")),
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/missing-vm"),
	// 				ErrorCode: to.Ptr("VmNotFound"),
	// 			},
	// 		},
	// 	},
	// }
}
