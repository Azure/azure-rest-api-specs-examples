package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/DaprSubscriptions_List.json
func ExampleDaprSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDaprSubscriptionsClient().NewListPager("examplerg", "myenvironment", nil)
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
		// page.DaprSubscriptionsCollection = armappcontainers.DaprSubscriptionsCollection{
		// 	Value: []*armappcontainers.DaprSubscription{
		// 		{
		// 			Name: to.Ptr("mybulksubscription"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/daprSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprSubscriptions/mybulksubscription"),
		// 			Properties: &armappcontainers.DaprSubscriptionProperties{
		// 				BulkSubscribe: &armappcontainers.DaprSubscriptionBulkSubscribeOptions{
		// 					Enabled: to.Ptr(true),
		// 					MaxAwaitDurationMs: to.Ptr[int32](500),
		// 					MaxMessagesCount: to.Ptr[int32](123),
		// 				},
		// 				PubsubName: to.Ptr("mypubsubcomponent"),
		// 				Routes: &armappcontainers.DaprSubscriptionRoutes{
		// 					Default: to.Ptr("/products"),
		// 					Rules: []*armappcontainers.DaprSubscriptionRouteRule{
		// 					},
		// 				},
		// 				Topic: to.Ptr("inventory"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mydefaultsubscription"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/daprSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprSubscriptions/mydefaultsubscription"),
		// 			Properties: &armappcontainers.DaprSubscriptionProperties{
		// 				PubsubName: to.Ptr("mypubsubcomponent"),
		// 				Routes: &armappcontainers.DaprSubscriptionRoutes{
		// 					Default: to.Ptr("/products"),
		// 					Rules: []*armappcontainers.DaprSubscriptionRouteRule{
		// 					},
		// 				},
		// 				Topic: to.Ptr("inventory"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myroutingsubscription"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/daprSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprSubscriptions/myroutingsubscription"),
		// 			Properties: &armappcontainers.DaprSubscriptionProperties{
		// 				Metadata: map[string]*string{
		// 					"foo": to.Ptr("bar"),
		// 					"hello": to.Ptr("world"),
		// 				},
		// 				PubsubName: to.Ptr("mypubsubcomponent"),
		// 				Routes: &armappcontainers.DaprSubscriptionRoutes{
		// 					Default: to.Ptr("/products"),
		// 					Rules: []*armappcontainers.DaprSubscriptionRouteRule{
		// 						{
		// 							Path: to.Ptr("/widgets"),
		// 							Match: to.Ptr("event.type == 'widget'"),
		// 						},
		// 						{
		// 							Path: to.Ptr("/gadgets"),
		// 							Match: to.Ptr("event.type == 'gadget'"),
		// 					}},
		// 				},
		// 				Topic: to.Ptr("inventory"),
		// 			},
		// 	}},
		// }
	}
}
