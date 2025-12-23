package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateLinkServiceListPrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_NewListPrivateEndpointConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListPrivateEndpointConnectionsPager("rg1", "testPls", nil)
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
		// page.PrivateEndpointConnectionListResult = armnetwork.PrivateEndpointConnectionListResult{
		// 	Value: []*armnetwork.PrivateEndpointConnection{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/privateEndpointConnections/testPlePeConnection1"),
		// 			Name: to.Ptr("testPlePeConnection1"),
		// 			Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("approved it for some reason."),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/privateEndpointConnections/testPlePeConnection2"),
		// 			Name: to.Ptr("testPlePeConnection2"),
		// 			Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("rejected by some reason."),
		// 					Status: to.Ptr("Rejected"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
