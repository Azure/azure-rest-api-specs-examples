package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/Occurrences_ListResources_MaximumSet_Gen.json
func ExampleOccurrencesClient_NewListResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListResourcesPager("rgcomputeschedule", "scheduled-action-01", "11111111-1111-1111-1111-111111111111", nil)
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
		// page = armcomputeschedule.OccurrencesClientListResourcesResponse{
		// 	OccurrenceResourceListResponse: armcomputeschedule.OccurrenceResourceListResponse{
		// 		Value: []*armcomputeschedule.OccurrenceResource{
		// 			{
		// 				Name: to.Ptr("myVm"),
		// 				ID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				Type: to.Ptr("Microsoft.ComputeSchedule/scheduledActions/occurrences/resources"),
		// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				NotificationSettings: []*armcomputeschedule.NotificationProperties{
		// 					{
		// 						Destination: to.Ptr("admin@contoso.com"),
		// 						Type: to.Ptr(armcomputeschedule.NotificationTypeEmail),
		// 						Language: to.Ptr(armcomputeschedule.LanguageEnUs),
		// 						Disabled: to.Ptr(true),
		// 					},
		// 				},
		// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:48.436Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armcomputeschedule.ResourceProvisioningStateSucceeded),
		// 				ErrorDetails: &armcomputeschedule.Error{
		// 					Code: to.Ptr("ResourceNotFound"),
		// 					Message: to.Ptr("The specified resource was not found."),
		// 					Target: to.Ptr("virtualMachines/myVm"),
		// 					Details: []*armcomputeschedule.Error{
		// 					},
		// 					Innererror: &armcomputeschedule.InnerError{
		// 						Code: to.Ptr("ResourceNotFoundError"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/providers/Microsoft.ComputeSchedule/scheduledActions?api-version=2026-04-15-preview&$skiptoken=abc123"),
		// 	},
		// }
	}
}
