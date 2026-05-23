package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rgtest01", "cachetest01", nil)
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
		// page = armredis.PrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionListResult: armredis.PrivateEndpointConnectionListResult{
		// 		Value: []*armredis.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("pectest01"),
		// 				Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
		// 				Properties: &armredis.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armredis.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("pectest01"),
		// 				Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
		// 				Properties: &armredis.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armredis.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
