package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkListOperationErrors_PaginatedSuccess.json
func ExampleVirtualMachineBulkOperationsClient_NewBulkListOperationErrorsPager_twoListFailedOperationErrorsWithPaginatedResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineBulkOperationsClient().NewBulkListOperationErrorsPager("example-rg", "eastus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbulkactions.VirtualMachineBulkOperationsClientBulkListOperationErrorsResponse{
		// 	ListBulkOperationErrorsResponse: armbulkactions.ListBulkOperationErrorsResponse{
		// 		Value: []*armbulkactions.ResourceOperation{
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-1"),
		// 				ErrorCode: to.Ptr("AllocationFailed"),
		// 				ErrorDetails: to.Ptr("Allocation failed because sufficient capacity was not available for the requested virtual machine size in eastus."),
		// 				Operation: &armbulkactions.ResourceOperationDetails{
		// 					OperationID: to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-1"),
		// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					State: to.Ptr(armbulkactions.OperationStateFailed),
		// 					CompletedAt: to.Ptr(time.Date(2026, time.August, 31, 18, 18, 0, 0, time.UTC)),
		// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
		// 						ErrorCode: to.Ptr("AllocationFailed"),
		// 						ErrorDetails: to.Ptr("Allocation failed because sufficient capacity was not available for the requested virtual machine size in eastus."),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-0"),
		// 				ErrorCode: to.Ptr("OperationNotAllowed"),
		// 				ErrorDetails: to.Ptr("The virtual machine is not configured to support hibernation."),
		// 				Operation: &armbulkactions.ResourceOperationDetails{
		// 					OperationID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff"),
		// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-0"),
		// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeHibernate),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					State: to.Ptr(armbulkactions.OperationStateFailed),
		// 					CompletedAt: to.Ptr(time.Date(2026, time.August, 31, 18, 12, 0, 0, time.UTC)),
		// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
		// 						ErrorCode: to.Ptr("OperationNotAllowed"),
		// 						ErrorDetails: to.Ptr("The virtual machine is not configured to support hibernation."),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/locations/eastus/listBulkOperationErrors?api-version=2026-10-06-preview&$skiptoken=page-2"),
		// 	},
		// }
	}
}
