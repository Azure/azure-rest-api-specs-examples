package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/Occurrences_ListResources_PagedSuccess.json
func ExampleOccurrencesClient_NewListResourcesPager_twoListAPageOfResourcesInARecurringScheduledActionOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListResourcesPager("example-rg", "weekday-start", "88888888-8888-8888-8888-888888888888", nil)
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
		// 				NotificationSettings: []*armbulkactions.NotificationProperties{
		// 					{
		// 						Destination: to.Ptr("admin@contoso.com"),
		// 						Type: to.Ptr(armbulkactions.NotificationTypeEmail),
		// 						Language: to.Ptr(armbulkactions.LanguageEnUs),
		// 						Disabled: to.Ptr(false),
		// 					},
		// 				},
		// 				ScheduledTime: to.Ptr(time.Date(2026, time.September, 16, 14, 0, 0, 0, time.UTC)),
		// 				ProvisioningState: to.Ptr(armbulkactions.OccurrenceResourceProvisioningStateSucceeded),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-01"),
		// 				Name: to.Ptr("web-vm-01"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			},
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
		// 				ScheduledTime: to.Ptr(time.Date(2026, time.September, 16, 14, 0, 0, 0, time.UTC)),
		// 				ProvisioningState: to.Ptr(armbulkactions.OccurrenceResourceProvisioningStateFailed),
		// 				ErrorDetails: &armbulkactions.Error{
		// 					Code: to.Ptr("OperationNotAllowed"),
		// 					Message: to.Ptr("The virtual machine cannot be started while it is being deallocated."),
		// 					Target: to.Ptr("virtualMachines"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
		// 				Name: to.Ptr("web-vm-02"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start/occurrences/88888888-8888-8888-8888-888888888888/resources?api-version=2026-10-06-preview&$skiptoken=page2"),
		// 	},
		// }
	}
}
