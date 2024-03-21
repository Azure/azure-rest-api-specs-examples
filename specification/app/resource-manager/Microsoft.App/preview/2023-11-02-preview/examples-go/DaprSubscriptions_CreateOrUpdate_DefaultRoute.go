package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DaprSubscriptions_CreateOrUpdate_DefaultRoute.json
func ExampleDaprSubscriptionsClient_CreateOrUpdate_createOrUpdateDaprSubscriptionWithDefaultRouteOnly() {
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
			PubsubName: to.Ptr("mypubsubcomponent"),
			Routes: &armappcontainers.DaprSubscriptionRoutes{
				Default: to.Ptr("/products"),
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
	// 		PubsubName: to.Ptr("mypubsubcomponent"),
	// 		Routes: &armappcontainers.DaprSubscriptionRoutes{
	// 			Default: to.Ptr("/products"),
	// 			Rules: []*armappcontainers.DaprSubscriptionRouteRule{
	// 			},
	// 		},
	// 		Topic: to.Ptr("inventory"),
	// 	},
	// }
}
