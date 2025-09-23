package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/DaprSubscriptions_CreateOrUpdate_RouteRulesAndMetadata.json
func ExampleDaprSubscriptionsClient_CreateOrUpdate_createOrUpdateDaprSubscriptionWithRouteRulesAndMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDaprSubscriptionsClient().CreateOrUpdate(ctx, "examplerg", "myenvironment", "mysubscription", armappcontainers.DaprSubscription{
		Properties: &armappcontainers.DaprSubscriptionProperties{
			Metadata: map[string]*string{
				"foo":   to.Ptr("bar"),
				"hello": to.Ptr("world"),
			},
			PubsubName: to.Ptr("mypubsubcomponent"),
			Routes: &armappcontainers.DaprSubscriptionRoutes{
				Default: to.Ptr("/products"),
				Rules: []*armappcontainers.DaprSubscriptionRouteRule{
					{
						Path:  to.Ptr("/widgets"),
						Match: to.Ptr("event.type == 'widget'"),
					},
					{
						Path:  to.Ptr("/gadgets"),
						Match: to.Ptr("event.type == 'gadget'"),
					}},
			},
			Topic: to.Ptr("inventory"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DaprSubscription = armappcontainers.DaprSubscription{
	// 	Name: to.Ptr("mysubscription"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/daprSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprSubscriptions/mysubscription"),
	// 	Properties: &armappcontainers.DaprSubscriptionProperties{
	// 		Metadata: map[string]*string{
	// 			"foo": to.Ptr("bar"),
	// 			"hello": to.Ptr("world"),
	// 		},
	// 		PubsubName: to.Ptr("mypubsubcomponent"),
	// 		Routes: &armappcontainers.DaprSubscriptionRoutes{
	// 			Default: to.Ptr("/products"),
	// 			Rules: []*armappcontainers.DaprSubscriptionRouteRule{
	// 				{
	// 					Path: to.Ptr("/widgets"),
	// 					Match: to.Ptr("event.type == 'widget'"),
	// 				},
	// 				{
	// 					Path: to.Ptr("/gadgets"),
	// 					Match: to.Ptr("event.type == 'gadget'"),
	// 			}},
	// 		},
	// 		Topic: to.Ptr("inventory"),
	// 	},
	// }
}
