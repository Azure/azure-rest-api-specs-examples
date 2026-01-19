package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/Occurrences_ListResources_MaximumSet_Gen.json
func ExampleOccurrencesClient_NewListResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListResourcesPager("rgcomputeschedule", "myScheduledAction", "CB26D7CB-3E27-465F-99C8-EAF7A4118245", nil)
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
		// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				NotificationSettings: []*armcomputeschedule.NotificationProperties{
		// 					{
		// 						Destination: to.Ptr("wbhryycyolvnypjxzlawwvb"),
		// 						Type: to.Ptr(armcomputeschedule.NotificationTypeEmail),
		// 						Language: to.Ptr(armcomputeschedule.LanguageEnUs),
		// 						Disabled: to.Ptr(true),
		// 					},
		// 				},
		// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:59.751Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armcomputeschedule.ResourceProvisioningStateSucceeded),
		// 				ErrorDetails: &armcomputeschedule.Error{
		// 					Code: to.Ptr("baxjmkbhoatqcj"),
		// 					Message: to.Ptr("chapcwfkqymeof"),
		// 					Target: to.Ptr("mkirmorowetsigohjamvk"),
		// 					Details: []*armcomputeschedule.Error{
		// 					},
		// 					Innererror: &armcomputeschedule.InnerError{
		// 						Code: to.Ptr("cgalioufsabcwatbxa"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				Name: to.Ptr("dkmlhpipnlqh"),
		// 				Type: to.Ptr("xgq"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/as"),
		// 	},
		// }
	}
}
