package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/Occurrences_ListResources_BasicSuccess.json
func ExampleOccurrencesClient_NewListResourcesPager_oneListResourcesInARecurringScheduledActionOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListResourcesPager("example-rg", "weekday-start", "77777777-7777-7777-7777-777777777777", nil)
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
		// page = armbulkactions.OccurrencesClientListResourcesResponse{
		// 	OccurrenceResourceListResponse: armbulkactions.OccurrenceResourceListResponse{
		// 		Value: []*armbulkactions.OccurrenceResource{
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-01"),
		// 				ScheduledTime: to.Ptr(time.Date(2026, time.September, 15, 14, 0, 0, 0, time.UTC)),
		// 				ProvisioningState: to.Ptr(armbulkactions.OccurrenceResourceProvisioningStateSucceeded),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-01"),
		// 				Name: to.Ptr("web-vm-01"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			},
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
		// 				ScheduledTime: to.Ptr(time.Date(2026, time.September, 15, 14, 0, 0, 0, time.UTC)),
		// 				ProvisioningState: to.Ptr(armbulkactions.OccurrenceResourceProvisioningStateSucceeded),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
		// 				Name: to.Ptr("web-vm-02"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
