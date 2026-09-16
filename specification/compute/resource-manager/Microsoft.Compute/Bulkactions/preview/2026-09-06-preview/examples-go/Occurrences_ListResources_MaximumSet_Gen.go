package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-09-06-preview/Occurrences_ListResources_MaximumSet_Gen.json
func ExampleOccurrencesClient_NewListResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListResourcesPager("rgcompute", "myScheduledAction", "67b5bada-4772-43fc-8dbb-402476d98a45", nil)
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
		// 				ResourceID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				NotificationSettings: []*armbulkactions.NotificationProperties{
		// 					{
		// 						Destination: to.Ptr("admin@contoso.com"),
		// 						Type: to.Ptr(armbulkactions.NotificationTypeEmail),
		// 						Language: to.Ptr(armbulkactions.LanguageEnUs),
		// 						Disabled: to.Ptr(true),
		// 					},
		// 				},
		// 				ScheduledTime: to.Ptr(time.Date(2026, time.August, 5, 0, 30, 0, 0, time.UTC)),
		// 				ProvisioningState: to.Ptr(armbulkactions.OccurrenceResourceProvisioningStateFailed),
		// 				ErrorDetails: &armbulkactions.Error{
		// 					Code: to.Ptr("InternalServerError"),
		// 					Message: to.Ptr("An internal error occurred."),
		// 					Target: to.Ptr("virtualMachines"),
		// 					Details: []*armbulkactions.Error{
		// 					},
		// 					Innererror: &armbulkactions.InnerError{
		// 						Code: to.Ptr("InnerErrorCode"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				Name: to.Ptr("myVm"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/scheduledActions/myScheduledAction/occurrences/67b5bada-4772-43fc-8dbb-402476d98a45/resources?api-version=2026-09-06-preview&$skiptoken=page2"),
		// 	},
		// }
	}
}
