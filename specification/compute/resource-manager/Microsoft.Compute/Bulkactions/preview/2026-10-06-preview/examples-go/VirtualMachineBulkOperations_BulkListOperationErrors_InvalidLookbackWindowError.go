package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkListOperationErrors_InvalidLookbackWindowError.json
func ExampleVirtualMachineBulkOperationsClient_NewBulkListOperationErrorsPager_threeResponseWithAnInvalidLookbackWindowError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineBulkOperationsClient().NewBulkListOperationErrorsPager("example-rg", "eastus", &armbulkactions.VirtualMachineBulkOperationsClientBulkListOperationErrorsOptions{
		LookbackInMinutes: to.Ptr[int32](0)})
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
		// 		},
		// 	},
		// }
	}
}
