package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().CreateOrUpdate(ctx, "resourcegroup", "example-RelayNamespace-5849", "{privateEndpointConnection name}", armrelay.PrivateEndpointConnection{
		Properties: &armrelay.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armrelay.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"),
			},
			PrivateLinkServiceConnectionState: &armrelay.ConnectionState{
				Description: to.Ptr("You may pass"),
				Status:      to.Ptr(armrelay.PrivateLinkConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.PrivateEndpointConnectionsClientCreateOrUpdateResponse{
	// 	PrivateEndpointConnection: &armrelay.PrivateEndpointConnection{
	// 		Name: to.Ptr("{privateEndpointConnection name}"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces/PrivateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/alitest/providers/Microsoft.Relay/namespaces/relay-private-endpoint-test/privateEndpointConnections/{privateEndpointConnection name}"),
	// 		Location: to.Ptr("South Central US"),
	// 		Properties: &armrelay.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armrelay.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armrelay.ConnectionState{
	// 				Description: to.Ptr("You may pass"),
	// 				Status: to.Ptr(armrelay.PrivateLinkConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armrelay.EndPointProvisioningStateUpdating),
	// 		},
	// 	},
	// }
}
