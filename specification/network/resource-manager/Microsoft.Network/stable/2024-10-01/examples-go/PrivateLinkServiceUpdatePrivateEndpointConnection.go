package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_UpdatePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().UpdatePrivateEndpointConnection(ctx, "rg1", "testPls", "testPlePeConnection", armnetwork.PrivateEndpointConnection{
		Name: to.Ptr("testPlePeConnection"),
		Properties: &armnetwork.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armnetwork.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
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
	// res.PrivateEndpointConnection = armnetwork.PrivateEndpointConnection{
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/privateEndpointConnections/testPlePeConnection"),
	// 	Name: to.Ptr("testPlePeConnection"),
	// 	Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("approved it for some reason."),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 	},
	// }
}
