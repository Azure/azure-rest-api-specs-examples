package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/RoutingIntentGet.json
func ExampleRoutingIntentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutingIntentClient().Get(ctx, "rg1", "virtualHub1", "Intent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoutingIntent = armnetwork.RoutingIntent{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routingIntent/Intent1"),
	// 	Name: to.Ptr("Intent1"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs/routingIntent"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RoutingIntentProperties{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingPolicies: []*armnetwork.RoutingPolicy{
	// 			{
	// 				Name: to.Ptr("InternetTraffic"),
	// 				Destinations: []*string{
	// 					to.Ptr("Internet")},
	// 					NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("PrivateTrafficPolicy"),
	// 					Destinations: []*string{
	// 						to.Ptr("PrivateTraffic")},
	// 						NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				}},
	// 			},
	// 		}
}
