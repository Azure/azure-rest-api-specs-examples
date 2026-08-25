package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/BulkCreateCustom_VirtualMachinesGetOperationStatus_MaximumSet_Gen.json
func ExampleBulkCreateCustomClient_NewVirtualMachinesGetOperationStatusPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("700935bc-adf2-4176-b9ad-c571731c09fc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBulkCreateCustomClient().NewVirtualMachinesGetOperationStatusPager("local-test-rg", "eastus", "00000000-0000-0000-0000-000000000102", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Results {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbulkactions.BulkCreateCustomClientVirtualMachinesGetOperationStatusResponse{
		// 	BulkCreateCustomOperationStatusListResult: armbulkactions.BulkCreateCustomOperationStatusListResult{
		// 		Results: []*armbulkactions.ResourceOperation{
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/700935bc-adf2-4176-b9ad-c571731c09fc/resourceGroups/local-test-rg/providers/Microsoft.Compute/virtualMachines/vm-001"),
		// 				Operation: &armbulkactions.ResourceOperationDetails{
		// 					OperationID: to.Ptr("3303b164-d533-4f5d-8cb2-54f9585fc55c"),
		// 					ResourceID: to.Ptr("/subscriptions/700935bc-adf2-4176-b9ad-c571731c09fc/resourceGroups/local-test-rg/providers/Microsoft.Compute/virtualMachines/vm-001"),
		// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeCreate),
		// 					SubscriptionID: to.Ptr("700935bc-adf2-4176-b9ad-c571731c09fc"),
		// 					Deadline: to.Ptr(time.Date(2026, time.August, 4, 15, 40, 8, 280208200, time.UTC)),
		// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
		// 					State: to.Ptr(armbulkactions.OperationStatePendingScheduling),
		// 					Timezone: to.Ptr("UTC"),
		// 					RetryPolicy: &armbulkactions.RetryPolicy{
		// 						RetryCount: to.Ptr[int32](1),
		// 						RetryWindowInMinutes: to.Ptr[int32](60),
		// 					},
		// 				},
		// 				VirtualMachineInfo: &armbulkactions.VirtualMachineInfo{
		// 					VMSize: to.Ptr("Standard_D4as_v5"),
		// 					Zone: to.Ptr("1"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/700935bc-adf2-4176-b9ad-c571731c09fc/resourceGroups/local-test-rg/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom/00000000-0000-0000-0000-000000000102/virtualMachinesGetOperationStatus?api-version=2026-08-06-preview&$skiptoken=page2"),
		// 	},
		// }
	}
}
