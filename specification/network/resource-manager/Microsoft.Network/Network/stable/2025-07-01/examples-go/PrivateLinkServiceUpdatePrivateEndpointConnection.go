package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/PrivateLinkServiceUpdatePrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_UpdatePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().UpdatePrivateEndpointConnection(ctx, "rg1", "testPls", "testPlePeConnection", armnetwork.PrivateEndpointConnection{
		Name: to.Ptr("testPlePeConnection"),
		Properties: &armnetwork.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armnetwork.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
			},
			PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
				Description: to.Ptr("approved it for some reason."),
				Status:      to.Ptr("Approved"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.PrivateLinkServicesClientUpdatePrivateEndpointConnectionResponse{
	// 	PrivateEndpointConnection: armnetwork.PrivateEndpointConnection{
	// 		Name: to.Ptr("testPlePeConnection"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/privateEndpointConnections/testPlePeConnection"),
	// 		Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("approved it for some reason."),
	// 				Status: to.Ptr("Approved"),
	// 			},
	// 		},
	// 	},
	// }
}
