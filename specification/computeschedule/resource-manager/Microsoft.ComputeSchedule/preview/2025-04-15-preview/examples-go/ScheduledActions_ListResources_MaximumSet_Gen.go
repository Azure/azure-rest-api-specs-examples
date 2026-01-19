package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_ListResources_MaximumSet_Gen.json
func ExampleScheduledActionsClient_NewListResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScheduledActionsClient().NewListResourcesPager("rgcomputeschedule", "myScheduledAction", nil)
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
		// page = armcomputeschedule.ScheduledActionsClientListResourcesResponse{
		// 	ResourceListResponse: armcomputeschedule.ResourceListResponse{
		// 		Value: []*armcomputeschedule.ScheduledActionResource{
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
		// 				ID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		// 				Name: to.Ptr("dkmlhpipnlqh"),
		// 				Type: to.Ptr("xgq"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/awac"),
		// 	},
		// }
	}
}
